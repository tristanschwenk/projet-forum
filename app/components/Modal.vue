<template>
<div class="modal modal-blur fade" :id="id" tabindex="-1" aria-modal="true" role="dialog">
  <div class="modal-dialog modal-dialog-centered" role="document">
    <div class="modal-content">
      <button v-if="closeable" @click="emit('close')" type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      <div :class="'modal-status bg-'+type"></div>
      <div class="modal-body text-center py-4">
        <!-- Download SVG icon from http://tabler-icons.io/i/alert-triangle -->
        <!-- SVG icon code with class="mb-2 text-danger icon-lg" -->
        <h3>{{title}}</h3>
        <div class="text-muted">{{message}}</div>
      </div>
      <div class="modal-footer">
        <div class="w-100">
          <div class="row">
            <div v-if="cancelable" class="col">
              <a @click="emit('cancel')" href="#" :class="'btn w-100 btn-'+cancelColor" data-bs-dismiss="modal">
                {{cancelText}}
              </a>
            </div>
            <div class="col">
              <a @click="emit('confirm')" href="#" :class="'btn w-100 btn-'+(confirmColor || type)" data-bs-dismiss="modal">
                {{confirmText}}
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  props:{
    id: {
      type: String,
      required: true,
    },
    type: {
      type: String,
      default: "info",
    },
    closeable: {
      type: Boolean,
      default: false,
    },
    cancelable: {
      type: Boolean,
      default: false,
    },
    cancelText: {
      type: String,
      default: "Cancel",
    },
    cancelColor: {
      type: String,
      default: "white",
    },
    confirmText: {
      type: String,
      default: "Confirm",
    },
    confirmColor: {
      type: String,
      default: null,
    },
    title: {
      type: String,
      default: 'Alert',
    },
    message: {
      type: String,
      default: ''
    }
  },
  methods: {
    emit(eventName) {
      this.$emit(eventName);
    },
  },
}
</script>