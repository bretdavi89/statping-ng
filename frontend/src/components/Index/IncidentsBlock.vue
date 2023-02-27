<template>
  <div
    v-if="incidents.length > 0"
  >
    <div class="row">
      <h5
        class="h5 col-12 mb-3 mt-2 text-dim"
      >
        <font-awesome-icon
          :icon="expanded ? 'minus' : 'plus'"
          class="pointer mr-3"
          @click="toggle"
        /> {{ $t('incidents') }}
      </h5>
      <div
        v-for="incident in incidents"
        v-if="expanded"
        :key="incident.id"
        class="col-12 mt-2"
      >
        <span class="braker mt-1 mb-3" />
        <h6>
          {{ incident.title }}
          <span class="font-2 float-right">{{ niceDate(incident.created_at) }}</span>
        </h6>
        <div
          class="font-2 mb-3"
          v-html="markdown(incident.description)"
        />
        <IncidentUpdate
          v-for="(update, i) in incident.updates"
          :key="i"
          :update="update"
          :admin="false"
        />
      </div>
    </div>
  </div>
</template>

<script>
import Api from '../../API';
import IncidentUpdate from '@/components/Elements/IncidentUpdate';

export default {
    name: 'IncidentsBlock',
    components: {
        IncidentUpdate
    },
    props: {
        service: {
            type: Object,
            required: true
        }
    },
    data () {
        return {
            expanded: false
        };
    },
    computed: {
        core () {
            return this.$store.getters.core;
        },
        incidents () {
            const incidents = this.$store.getters.serviceIncidents(this.service.id);
            const incidentCutoff = this.daysAgo(this.core.number_of_days_for_incidents);
            return incidents.filter(i => this.isBefore(incidentCutoff, i.created_at));
        }
    },
    methods: {
        toggle () {
            this.expanded = !this.expanded;
        },
        badgeClass (val) {
            switch (val.toLowerCase()) {
                case 'resolved':
                    return 'badge-success';
                case 'update':
                    return 'badge-info';
                case 'investigating':
                    return 'badge-danger';
            }
        },
        async incident_updates (incident) {
            return await Api.incident_updates(incident);
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
