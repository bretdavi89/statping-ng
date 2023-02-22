<template>
  <div
    class="card shadow mb-4"
    role="alert"
  >
    <div class="card-body pb-2">
      <h3 class="mb-3 font-weight-bold">
        {{ message.title }}
      </h3>
      <div
        class="mb-2"
        v-html="markdown(message.description)"
      />
      <div class="col-12 mb-0">
        <div class="dates">
          <div class="start">
            <strong>STARTS</strong> {{ niceDate(message.start_on) }}
            <span />
          </div>
          <div class="ends">
            <strong>ENDS</strong> {{ niceDate(message.end_on) }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
    name: 'MessageBlock',
    props: {
        message: {
            type: Object,
            required: true,
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
