<template>
  <div
    class="col-12 mb-3 pb-2 border-bottom"
    role="alert"
  >
    <span
      class="font-weight-bold text-capitalize"
      :class="{'text-success': update.type.toLowerCase()==='resolved', 'text-danger': update.type.toLowerCase()==='investigating', 'text-warning': update.type.toLowerCase()==='update'}"
    >{{ update.type }}</span>
    <span class="text-muted">- {{ update.message }}
      <button
        v-if="admin"
        type="button"
        class="close"
        @click="delete_update(update)"
      >
        <span aria-hidden="true">&times;</span>
      </button>
    </span>
    <span class="d-block small">{{ ago(update.created_at) }} ago</span>
  </div>
</template>

<script>
import Api from '@/API';

export default {
    name: 'IncidentUpdate',
    props: {
        update: {
            required: true
        },
        admin: {
            required: true
        },
        onUpdate: {
            required: false
        }
    },
    methods: {
        async delete_update (update) {
            this.res = await Api.incident_update_delete(update);
            if (this.res.status === 'success') {
                this.onUpdate();
            }
        },
    }
};
</script>

<style scoped>

</style>
