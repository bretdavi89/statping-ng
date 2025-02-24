package handlers

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/statping-ng/statping-ng/notifiers"
	"github.com/statping-ng/statping-ng/types/configs"
	"github.com/statping-ng/statping-ng/types/core"
	"github.com/statping-ng/statping-ng/types/null"
	"github.com/statping-ng/statping-ng/types/services"
	"github.com/statping-ng/statping-ng/utils"
)

func processSetupHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if core.App.Setup {
		sendErrorJson(errors.New("Statping has already been setup"), w, r)
		return
	}

	confgs, err := configs.LoadConfigForm(r)
	if err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	project := r.PostForm.Get("project")
	description := r.PostForm.Get("description")
	domain := r.PostForm.Get("domain")
	sendNews, _ := strconv.ParseBool(r.PostForm.Get("newsletter"))
	sendReports, _ := strconv.ParseBool(r.PostForm.Get("send_reports"))

	log.WithFields(utils.ToFields(core.App, confgs)).Debugln("new configs posted")

	if err = configs.ConnectConfigs(confgs, false); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	if err := confgs.Save(utils.Directory); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	exists := confgs.Db.HasTable("core")
	if !exists {
		if err := confgs.DropDatabase(); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := confgs.CreateDatabase(); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := configs.CreateAdminUser(); err != nil {
			sendErrorJson(err, w, r)
			return
		}

		if err := configs.TriggerSamples(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	if err = confgs.MigrateDatabase(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	log.Infoln("Migrating Notifiers...")
	notifiers.InitNotifiers()

	c := &core.Core{
		Name:                     project,
		Description:              description,
		ApiSecret:                utils.Params.GetString("API_SECRET"),
		Domain:                   domain,
		Version:                  core.App.Version,
		Started:                  utils.Now(),
		CreatedAt:                utils.Now(),
		UseCdn:                   null.NewNullBool(false),
		Footer:                   null.NewNullString(""),
		Language:                 confgs.Language,
		AllowReports:             null.NewNullBool(sendReports),
		NumberOfDaysForIncidents: null.NewNullInt64(30),
		NumberOfDaysForService:   null.NewNullInt64(90),
		ShowGraphs:               null.NewNullBool(true),
	}

	log.Infoln("Creating new Core")
	if err := c.Create(); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	core.App = c

	if sendNews {
		log.Infof("Sending email address %s to newsletter server", confgs.Email)
		if err := registerNews(confgs.Email, confgs.Domain); err != nil {
			log.Errorln(err)
		}
	}

	log.Infoln("Initializing new Statping instance")

	if _, err := services.SelectAllServices(true); err != nil {
		log.Errorln(err)
		sendErrorJson(err, w, r)
		return
	}

	services.CheckServices()

	core.App.Setup = true

	resetCookies()

	out := struct {
		Message string            `json:"message"`
		Config  *configs.DbConfig `json:"config"`
	}{
		"success",
		confgs,
	}
	returnJson(out, w, r)
}

func registerNews(email, domain string) error {
	if email == "" {
		return nil
	}
	v := url.Values{}
	v.Set("email", email)
	v.Set("domain", domain)
	v.Set("timezone", "UTC")
	resp, err := http.PostForm("https://news.statping.com/new", v)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
