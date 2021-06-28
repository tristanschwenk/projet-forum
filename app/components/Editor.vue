<template>
  <div class="editor">
    <client-only>
      <div class="mt-2 white">
        <quill-editor ref="editor" :content="value.html" :options="editorOption" @change="editorChange"/>
      </div>
    </client-only>
    <!-- <div class="row justify-content-end" v-if="isReply">
      <div class="col-1 mt-2 mb-3"><a class="btn btn-outline-link w-100" @click.prevent="reply">
          Reply
        </a>
      </div>
    </div> -->
  </div>
</template>

<style lang="scss">
.white {
  background-color: white;
}
</style>

<script>
  export default {
    props: {
      postID: {
        type: Number,
      },
      // isReply: {
      //   type: Boolean,
      //   default: true
      // },
      value: {
        type: Object,
        default: {
          html: '',
          quill: {},
          delta: [],
          body: "",
          text: '',
        }
      },
    },
    data() {
      return {
        editorOption: {
          theme: 'snow',
          modules: {
            toolbar: [
              ['bold', 'italic', 'underline', 'strike'],
              ['blockquote', 'code-block'],
              [{
                'header': [3, 4, 5, 6, false]
              }], // custom button values
              [{
                'list': 'ordered'
              }, {
                'list': 'bullet'
              }],
              [{
                'script': 'sub'
              }, {
                'script': 'super'
              }], // superscript/subscript
              [{
                'indent': '-1'
              }, {
                'indent': '+1'
              }], // outdent/indent

              [{
                'align': []
              }],

              ['clean']
            ]
          }
        }
      }
    },
    computed: {
      quill() {
        return this.$refs.editor.quill
      }
    },
    mounted() {

    },
    methods: {
      // async reply() {

      //   this.$emit('submit', {
      //     body: JSON.stringify(this.quill.getContents()),
      //     postID: this.postID
      //   });

      //   try {
      //     let createdEntity = await this.$axios.$post('/api/replies/create', {
      //       body: JSON.stringify(this.quill.getContents()),
      //       postID: this.postID
      //     })

      //     this.$emit('success', createdEntity);
      //     this.content = "";
      //   } catch (e) {
      //     console.log(e.response?.data || e);
      //     this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred';
      //     this.$emit('failure', e.response?.data || e);
      //   }
      // },
      editorChange({html, quill, text}) {
        const delta = quill.getContents().ops;
        const body = JSON.stringify(delta);
        this.$emit('input', {html, quill, delta, text, body});
      }
    }
  }

</script>
