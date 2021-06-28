<template>
  <div class="main">
    <div class="page-header">
      <div class="container-xl">
        <div class="row align-items-center">
          <div class="col">
            <div class="page-pretitle">
              Ezyo Forum
            </div>
            <h2 class="page-title">{{topic.name}}</h2>
            <div class="text-muted mt-1">
              #{{topic.id}}
              Created {{creationTime}} | Last activity {{lastUpdate || creationTime}}
            </div>

            <div class="mt-1">
                <a class="text-muted" data-bs-toggle="collapse" href="#collapseExample" role="button"
                  aria-expanded="false" aria-controls="collapseExample">
                  See description
                </a>
              <div class="collapse mt-1" id="collapseExample" v-html="$parseDelta(topic.body)"></div>
            </div>
          </div>
          <div class="col-auto"> 
            <div class="btn-list">
              <a href="#" class="btn btn-white" data-bs-toggle="modal"
                data-bs-target="#modal-new-post">
                <!-- Download SVG icon from http://tabler-icons.io/i/edit -->
                <!-- SVG icon code -->
                New
              </a>
              <a href="#" class="btn btn-white d-none d-md-inline-flex">
                <!-- Download SVG icon from http://tabler-icons.io/i/bell -->
                <!-- SVG icon code -->
                Subscribe
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="page-body">
      <div class="container-lg">
        <div class="grid-post">
          <div v-for="post in posts.slice().reverse()" :key="post.id">
            <Post :post="post" :footer=true :button=true></Post>
          </div>
        </div>
      </div>
    </div>
    <NewEntityModal type="post" id="modal-new-post" :topicId="topic.id" :large="true" @success="$fetch()"></NewEntityModal>
  </div>
</template>

<style lang="scss">
  .grid-post {
    display: grid;
    grid-template-columns: 1fr;
    grid-auto-rows: 1fr;
    gap: 1em;
  }

</style>

<script>
  export default {
    head() {
      return {
        title: `${this.topic.name} - Ezyo Forum`,
      }
    },
    async asyncData({
      params,
      $axios
    }) {
      const topic = await $axios.$get(`/api/topic/${params.id}`)
      const posts = await $axios.$get(`/api/post/all/${params.id}`);
      return {
        posts,
        topic, 
        params
      };
    },
    async fetch() {
      this.posts = await this.$axios.$get(`/api/post/all/${this.params.id}`);
    },
    beforeMounted() {
      window.vueInstance = this;
      console.log("->>>>", this.posts);
    },
    computed: {
      creationTime() {
        return this.$relativeTime(this.topic.createdAt * 1000)
      },
      lastUpdate() {
        if (this.posts.length) {
          return this.$relativeTime(this.posts[this.posts.length - 1].createdAt * 1000)
        }
      }
    }
  }

</script>
