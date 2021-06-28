<template>
  <section class="container">
    <client-only>
      <quill-editor class="form-control" ref="editor" v-model="content" :options="editorOption" @blur="onEditorBlur($event)"
        @focus="onEditorFocus($event)" @ready="onEditorReady($event)" />

    </client-only>
  </section>
</template>

<script>
  export default {
    name: 'quill-example-nuxt',
    data() {
      return {
        content: '<p>I am Example</p>',
        editorOption: {
          // Some Quill options...
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
    mounted() {
     
    },
    methods: {
      onEditorBlur(editor) {
        console.log('editor blur!', editor)
      },
      onEditorFocus(editor) {
        console.log('editor focus!', editor)
      },
      onEditorReady(editor) {
        console.log('editor ready!', editor)
      }
    }
  }

</script>

<style lang="scss" scoped>
  .container {
    width: 60%;
    margin: 0 auto;
    padding: 50px 0;

    .quill-editor {
      min-height: 200px;
      max-height: 400px;
      overflow-y: auto;
    }
  }

</style>
