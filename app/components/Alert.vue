<template>
  <div :class="'alert alert-'+type">
      <div class="d-flex">
      <div>
        <h4 class="alert-title">{{title}}</h4>
        <div class="text-muted">
          <slot></slot>
        </div>
      </div>
    </div>
    <a class="btn-close" @click="close" data-bs-dismiss="alert" aria-label="close"></a>
  </div>
</template>

<script>
  export default {
    props: {
      type: {
        type: String,
        default: 'info'
      },
      title: {
        type: String,
        default: ''
      },
    },
    methods: {
      close() {
        this.$emit('close')
      }
    },
  }

</script>
