<template>
  <form @submit.prevent="saveSettings">
    <div class="card">
      <div class="card-header">
        Statping-ng Settings
      </div>
      <div class="card-body">
        <div class="form-group">
          <label>{{ $t('project_name') }}</label>
          <input
            id="project"
            v-model="core.name"
            type="text"
            class="form-control"
            placeholder="Great Uptime"
          >
        </div>

        <div class="form-group">
          <label>{{ $t('description') }}</label>
          <input
            id="description"
            v-model="core.description"
            type="text"
            class="form-control"
            placeholder="Great Uptime"
          >
        </div>

        <div class="form-group row">
          <div class="col-8 col-sm-9">
            <label>{{ $t('domain') }}</label>
            <input
              id="domain"
              v-model="core.domain"
              type="url"
              class="form-control"
            >
          </div>
          <div class="col-4 col-sm-3 mt-sm-1 mt-0">
            <label class="d-inline d-sm-none">{{ $t('enable_cdn') }}</label>
            <label class="d-none d-sm-block">{{ $t('enable_cdn') }}</label>
            <span
              id="using_cdn"
              class="switch"
              @click="core.using_cdn = !!core.using_cdn"
            >
              <input
                id="switch-normal"
                v-model="core.using_cdn"
                type="checkbox"
                name="using_cdn"
                class="switch"
                :checked="core.using_cdn"
              >
              <label for="switch-normal" />
            </span>
          </div>
        </div>

        <div class="form-group">
          <label>{{ $t('footer') }}</label>
          <textarea
            id="footer"
            v-model="core.footer"
            rows="4"
            class="form-control"
          >{{ core.footer }}</textarea>
          <small class="form-text text-muted">{{ $t('footer_notes') }}</small>
        </div>

        <div class="form-group">
          <label>{{ $t('language') }}</label>
          <select
            v-model="core.language"
            class="form-control"
          >
            <option value="en">
              English
            </option>
            <option value="es">
              Spanish
            </option>
            <option value="fr">
              French
            </option>
            <option value="ru">
              Russian
            </option>
            <option value="de">
              German
            </option>
            <option value="cs">
              Czech
            </option>
            <option value="ja">
              Japanese
            </option>
            <option value="it">
              Italian
            </option>
            <option value="ko">
              Korean
            </option>
            <option value="zh">
              Chinese
            </option>
          </select>
        </div>

        <div class="form-group">
          <label>{{ $t('number_of_days_for_incidents') }}</label>
          <input
            id="number_of_days_for_incidents"
            v-model="core.number_of_days_for_incidents"
            type="text"
            class="form-control"
            placeholder="30"
          >
          <small class="form-text text-muted">{{ $t('number_of_days_for_incidents_desc') }}</small>
        </div>

        <div class="form-group">
          <label>{{ $t('number_of_days_for_service') }}</label>
          <input
            id="number_of_days_for_service"
            v-model="core.number_of_days_for_service"
            type="text"
            class="form-control"
            placeholder="90"
          >
          <small class="form-text text-muted">{{ $t('number_of_days_for_service_desc') }}</small>
        </div>

        <div class="form-group row mt-3">
          <label class="col-sm-10 col-form-label">{{ $t('show_graphs') }}</label>
          <div class="col-sm-2 float-right">
            <span
              id="show_graphs"
              class="switch"
              @click="core.show_graphs = !!core.show_graphs"
            >
              <input
                id="switch_show_graphs"
                v-model="core.show_graphs"
                type="checkbox"
                name="show_graphs"
                class="switch"
                :checked="core.show_graphs"
              >
              <label for="switch_show_graphs" />
            </span>
          </div>
          <div class="col-12">
            <small>{{ $t('show_graphs_desc') }}</small>
          </div>
        </div>

        <div class="form-group row mt-3">
          <label class="col-sm-10 col-form-label">{{ $t('send_reports') }}</label>
          <div class="col-sm-2 float-right">
            <span
              id="allow_report"
              class="switch"
              @click="core.allow_reports = !!core.allow_reports"
            >
              <input
                id="switch_allow_report"
                v-model="core.allow_reports"
                type="checkbox"
                name="allow_report"
                class="switch"
                :checked="core.allow_reports"
              >
              <label for="switch_allow_report" />
            </span>
          </div>
          <div class="col-12">
            <small>{{ $t('send_reports_desc') }}</small>
          </div>
        </div>
      </div>
      <div class="card-footer">
        <button
          id="save_core"
          type="submit"
          class="btn btn-primary btn-block"
          :disabled="loading"
          @click.prevent="saveSettings"
        >
          <font-awesome-icon
            v-if="loading"
            icon="circle-notch"
            class="mr-2"
            spin
          />{{ $t('save_settings') }}
        </button>
      </div>
    </div>
  </form>
</template>

<script>
import Api from '../API';

export default {
    name: 'CoreSettings',
    data () {
        return {
            loading: false
        };
    },
    computed: {
        core () {
            return this.$store.getters.core;
        }
    },
    methods: {
        async saveSettings () {
            this.loading = true;
            const c = this.core;
            c.number_of_days_for_service = Number(c.number_of_days_for_service);
            c.number_of_days_for_incidents = Number(c.number_of_days_for_incidents);
            await Api.core_save(c);
            this.$store.commit('setCore', c);
            this.$i18n.locale = c.language || 'en';
            this.loading = false;
        },
        selectAll () {
            this.$refs.input.select();
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
