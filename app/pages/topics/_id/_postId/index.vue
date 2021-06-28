<template>
  <div class="main">
    <div class="page-header">
      <div class="container-xl">
        <div class="row align-items-center">
          <div class="col">
            <div class="page-pretitle">
              Ezyo Forum
            </div>
            <h2 class="page-title">{{post.name}}</h2>
            <div class="text-muted mt-1">
              #{{post.id}}
              Created {{creationTime}} | Last activity {{lastUpdate || creationTime}}
            </div>
          </div>
          <div class="col-auto">
            <div class="btn-list">
              <a href="#" class="btn btn-white d-none d-md-inline-flex">
                <template v-if="post.isUserSubscribed">
                  Unsubscribe
                </template>
                <template v-else>
                  Subscribe
                </template>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="page-body">
      <div class="container-lg">
        <Post :post="post" :footer=true :title=false :short=false :like="true"></Post>
        <div class="pt-3 pb-3">
          <div class="card container-lg mt-3 p-0">
            <Editor :postID="this.post.id" :isReply="true" v-model="editor"/>
                <div class="p-3"><button class="btn btn-outline-light text-dark float-end cl-auto" @click.prevent="create">Reply</button></div>
                
                          
          </div>

        </div>

        <template v-if="replies.length!=0">

          <div class="card mb-2" v-for="reply in replies.slice().reverse()" :key="reply.id">
            <div class="card-body" v-html="$parseDelta(reply.body)">
            </div>
            <!-- Card footer -->
            <div class="card-footer">
              <div class="row align-items-center">
                <div class="col-auto ms-auto">
                  <div class="avatar-list avatar-list-stacked">
                    <span class="avatar avatar-sm avatar-rounded">JL</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>


<script>
  import Editor from '../../../../components/Editor.vue';

  export default {
    components: {
      Editor
    },

    head() {
      return {
        title: `${this.post.name} - Ezyo Forum`,
      }
    },
    data() {
      return {
        body: "",
        editor: {
          body: "",
          text: ""
        },
        post: {},
        replies: [],
      }
    },
    async fetch() {
      this.post = await this.$axios.$get(`/api/post/${this.$route.params.postId}`);
      this.replies = await this.$axios.$get(`/api/replies/${this.$route.params.postId}`);
    },

    mounted() {
      console.log("-->",this.replies);
    },
    computed: {
      creationTime() {
        return this.$relativeTime(this.post.createdAt * 1000)
      },
      lastUpdate() {
        if (this.replies.length) {
          return this.$relativeTime(this.replies[this.replies.length - 1].createdAt * 1000)
        }
      },
    },
    methods: {
      async create() {
        const postID = this.post.id;

        const {body} = this.editor;


        let createdEntity,
            data = {body, postID},
            apiUrl = '/api/replies/create';

        try {
          createdEntity = await this.$axios.$post(apiUrl, data);
          this.$emit('success', createdEntity);
          this.$fetch()
          this.editor.html = ""
        } catch (e) {
          console.log(e.response?.data || e);
          this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred';

          this.$emit('error', e);
          return false;
        }
      }
    }
  }

</script>
