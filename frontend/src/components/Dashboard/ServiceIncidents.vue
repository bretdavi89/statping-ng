<template>
  <div class="col-12">
    <div
      v-for="incident in incidents"
      :key="incident.id"
      class="card contain-card mb-4"
    >
      <div class="card-header">
        Incident: {{ incident.title }}

        <div class="btn-group float-right">
          <button
            href="#"
            class="btn btn-sm btn-outline-secondary"
            @click.prevent="editIncident(incident, edit)"
          >
            <font-awesome-icon icon="edit" />
          </button>
          <button
            class="btn btn-sm btn-danger"
            @click="deleteIncident(incident)"
          >
            <font-awesome-icon icon="times" />
          </button>
        </div>
      </div>

      <FormIncidentUpdates :incident="incident" />

      <span class="font-2 p-2 pl-3">Created: {{ niceDate(incident.created_at) }} | Last Update: {{ niceDate(incident.updated_at) }}</span>
    </div>
    <FormIncident
      :edit="editChange"
      :in_incident="incident"
    />
  </div>
</template>

<script>
import Api from '../../API';

const FormIncidentUpdates = () => import(/* webpackChunkName: "dashboard" */ '@/forms/IncidentUpdates');
const FormIncident = () => import(/* webpackChunkName: "dashboard" */ '@/forms/Incident');

export default {
    name: 'ServiceIncidents',
    components: { FormIncidentUpdates, FormIncident },
    data () {
        return {
            edit: false,
            serviceID: 0,
            incident: {
                title: '',
                description: '',
                service: 0
            }
        };
    },

    computed: {
        incidents () {
            return this.$store.getters.serviceIncidents(this.serviceID);
        }
    },

    created () {
        this.serviceID = Number(this.$route.params.id);
        this.incident.service = Number(this.$route.params.id);
    },

    methods: {
        editChange (v) {
            this.incident = {};
            this.edit = v;
        },
        editIncident (incident, mode) {
            this.incident = incident;
            this.edit = !mode;
        },
        async delete (i) {
            await Api.incident_delete(i);
            const incidents = await Api.incidents();
            this.$store.commit('setIncidents', incidents);
        },
        async deleteIncident (incident) {
            const modal = {
                visible: true,
                title: 'Delete Incident',
                body: `Are you sure you want to delete Incident ${incident.title}?`,
                btnColor: 'btn-danger',
                btnText: 'Delete Incident',
                func: () => this.delete(incident)
            };
            this.$store.commit('setModal', modal);
        },


        async loadIncidents () {
            this.incidents = await Api.incidents_service(this.serviceID);
        }

    }
};
</script>
