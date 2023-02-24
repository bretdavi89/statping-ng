<template>
  <div>
    <div class="card contain-card">
      <div class="card-header">
        {{ incident.id ? `${$t('update')} ${incident.title}` : $t('incident_create') }}
        <transition name="slide-fade">
          <button
            v-if="incident.id"
            class="btn btn-sm float-right btn-danger btn-sm"
            @click="removeEdit"
          >
            {{ $t('close') }}
          </button>
        </transition>
      </div>
      <div class="card-body">
        <form @submit="saveIncident">
          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('title') }}</label>
            <div class="col-sm-8">
              <input
                id="title"
                v-model="incident.title"
                type="text"
                name="title"
                class="form-control"
                placeholder="Incident Title"
                required
              >
            </div>
          </div>

          <div class="form-group row">
            <label class="col-sm-4 col-form-label">{{ $t('description') }}</label>
            <div class="col-sm-8">
              <textarea
                id="description"
                v-model="incident.description"
                rows="5"
                name="description"
                class="form-control"
                required
              />
              <font-awesome-icon
                class="mdicon"
                icon="fab fa-markdown"
              />
            </div>
          </div>

          <div class="form-group row">
            <div class="col-sm-12">
              <button
                :disabled="!incident.title || !incident.description"
                type="submit"
                class="btn btn-block btn-primary"
                :class="{'btn-primary': !incident.id, 'btn-secondary': incident.id}"
                @click="saveIncident"
              >
                {{ incident.id ? $t('incident_edit') : $t('incident_create') }}
              </button>
            </div>
          </div>
          <div
            id="alerter"
            class="alert alert-danger d-none"
            role="alert"
          />
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Api from '../API';

export default {
    name: 'FormIncident',
    components: {
    },
    props: {
        in_incident: {
            type: Object
        },
        edit: {
            type: Function
        }
    },
    data () {
        return {
            incident: {
                title: '',
                description: '',
                service: 0
            }
        };
    },
    watch: {
        in_incident () {
            this.incident = this.in_incident;
        }
    },
    created () {
        this.serviceID = Number(this.$route.params.id);
        this.incident.service = Number(this.$route.params.id);
    },
    methods: {
        startChange (e) {
            window.console.log(e);
        },
        endChange (e) {
            window.console.log(e);
        },
        removeEdit () {
            this.incident = {};
            this.edit(false);
        },
        async saveIncident (e) {
            e.preventDefault();
            if (this.incident.id) {
                await this.updateIncident();
            } else {
                await this.createIncident();
            }
        },
        async createIncident () {
            await Api.incident_create(this.serviceID, this.incident);
            const incidents = await Api.incidents();
            this.$store.commit('setIncidents', incidents);
            this.incident = {};
        },
        async updateIncident () {
            await Api.incident_update(this.incident);
            const incidents = await Api.incidents();
            this.$store.commit('setIncidents', incidents);
            this.edit(false);
        },
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
