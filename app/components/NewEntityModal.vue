<template>
  <div class="modal modal-blur" :id="id" tabindex="-1" aria-modal="true" role="dialog">
    <div class="modal-dialog modal-dialog-centered" :class="{ 'modal-lg': large }" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">Create a new {{this.type}}</h5>
          <button @click="emit('close')" type="button" class="btn-close" data-bs-dismiss="modal"
            aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <div v-if="error" class="alert alert-danger mb-3" role="alert">
            <h4 class="alert-title">I'm so sorry&hellip;</h4>
            <div class="text-muted">{{error}}</div>
          </div>

          <div class="row mb-3 align-items-end">
            <div class="col">
              <label class="form-label">Name</label>
              <input type="text" class="form-control" v-model="name" required>
            </div>
          </div>
          <div>
            <label class="form-label">Description</label>
            <Editor v-model="editor" required/> </div>
        </div>
        
        <div class="modal-footer">
          <button type="button" class="btn me-auto" data-bs-dismiss="modal">Close</button>
          <button type="button" class="btn btn-primary" data-bs-dismiss="modal" @click.prevent="create" :disabled="btnDisabled">Create</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Editor from './Editor.vue';
  export default {
  components: { Editor },
    data() {
      return {
        name: "",
        editor: {
          body: "",
          text: ""
        },
        error: false
      }
    },

    props: {
      id: {
        type: String,
        required: true,
      },
      topicId: {
        type: Number,
      },
      type: {
        type: String,
        default: "topic"
      },
      large: {
        type: Boolean,
        default: false
      }
    },
    computed: {
      btnDisabled() {
        return this.name=="" || this.editor.text?.length==0 || this.editor.text == '\n';
      }
    },
    methods: {
      emit(eventName) {
        this.$emit(eventName);
      },

      async create() {
        const {
          name,
          topicId
        } = this;

        const {body} = this.editor;

        let createdEntity,
            data = {
              name,
              body
            },
            apiUrl = '/api/topic/create';

        // If topicId is set, then it's a post
        if (topicId) { 

          data = {
            ...data,
            topicId,
          };

          apiUrl = '/api/post/create';
        }

        try {
          createdEntity = await this.$axios.$post(apiUrl, data);
          this.$emit('success', createdEntity);
        } catch (e) {
          console.log(e.response?.data || e);
          this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred';

          this.$emit('error', e);
          return false;
        }
      }
    },
  }

</script>
