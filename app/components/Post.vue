<template>
  <div class="card">
    <div class="card-body">
      <div v-if="title">
        <h3 class="card-title">{{ post.name }}</h3>
        
      </div>
      <div :class="{short:short}" v-html="$parseDelta(post.body)"></div>
    </div>
    <div class="card-footer" v-if="footer">
      <nuxt-link v-if="button" :to='this.post.topicId + "/" + this.post.id'
        class="btn btn-outline-light text-dark btn-pill">
        Read More
      </nuxt-link>

      <div class="action-list" v-if="like">
        <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-thumb-up" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <path d="M7 11v8a1 1 0 0 1 -1 1h-2a1 1 0 0 1 -1 -1v-7a1 1 0 0 1 1 -1h3a4 4 0 0 0 4 -4v-1a2 2 0 0 1 4 0v5h3a2 2 0 0 1 2 2l-1 5a2 3 0 0 1 -2 2h-7a3 3 0 0 1 -3 -3"></path>
</svg>

<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-thumb-down" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <path d="M7 13v-8a1 1 0 0 0 -1 -1h-2a1 1 0 0 0 -1 1v7a1 1 0 0 0 1 1h3a4 4 0 0 1 4 4v1a2 2 0 0 0 4 0v-5h3a2 2 0 0 0 2 -2l-1 -5a2 3 0 0 0 -2 -2h-7a3 3 0 0 0 -3 3"></path>
</svg>
      </div>
    </div>
  </div>
</template>

<style lang="scss">
  .short {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

</style>

<script>
  export default {
    props: {
      post: {
        type: Object,
        default: {},
      },
      footer: {
        type: Boolean,
        default: false
      },
      button: {
        type: Boolean,
        default: false
      },
      title: {
        type: Boolean,
        default: true
      },
      short: {
        type: Boolean,
        default: true
      },
      like: {
        type: Boolean,
        default: false
      }
    },
    mounted() {
      window.vueInstance = this;
    },
    computed: {
      
    },
  };

</script>
