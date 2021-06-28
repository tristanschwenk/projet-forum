<template>
  <div v-if="hasData">
    <div class="row">
      <div class="col">
        <div class="text-truncate">
          <nuxt-link :to='"/topics/" + topic.id' v-if="topic"><strong>{{ entity.name }}</strong></nuxt-link>
          <nuxt-link :to='"/topics/" + post.topicId + "/" + post.id' v-if="post"><strong>{{ entity.name }}</strong></nuxt-link>
        </div>
        <div class="text-truncate" v-html="$parseDelta(entity.body)"></div>
        <div class="text-muted">{{ relativeTime }}</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    topic: {
      type: Object,
    },
    post: {
      type: Object,
    },
    
  },
  mounted(){
          window.vueInstance = this;
  },
  computed: {
    hasData() {
      return Boolean(this.topic || this.post);
    },
    entity() {
      return this.topic || this.post;
    },
    relativeTime() {
      const entity = (this.topic)?this.topic:this.post;

      let time = this.$dayjs(entity.createdAt*1000).fromNow();
      
      time = time[0].toUpperCase() + time.slice(1)
      return time
    }
  }
};
</script>
