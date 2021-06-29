<template>
  <div class="page-body">
    <div class="container-xl">
      <div class="row">
        <div class="col-md-6 col-lg-3">
          <div class="card">
            <div class="card-body p-4 text-center">
              <span class="avatar avatar-xl mb-3 avatar-rounded" :style="{backgroundImage: `url(${user.avatar})`}">
                {{(!user.avatar)?$getInitials(user.displayName):''}}
              </span>
              <h3 class="m-0 mb-1">{{ user.displayName }}</h3>
              <div class="text-muted">@{{ user.userName }}</div>
            </div>
            <div class="d-flex" v-if="isCurrentUser">
              <nuxt-link :to="'@'+user.userName+'/edit'" class="card-btn">
                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-pencil" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                  <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M4 20h4l10.5 -10.5a1.5 1.5 0 0 0 -4 -4l-10.5 10.5v4"></path>
                  <line x1="13.5" y1="6.5" x2="17.5" y2="10.5"></line>
                </svg>
                Edit
              </nuxt-link>
            </div>
          </div>
        </div>
        <div class="col-md-6 col-lg-9">
          <div class="card">
            <div class="empty" v-if="posts.length==0">
              <div class="empty-img"><img src="@/static/images/notfound.svg" height="128" alt="">
              </div>
              <p class="empty-title">No post found</p>
              <p class="empty-subtitle text-muted">
                This user seems very discrete...
              </p>
            </div>

          <div class="card-body" v-if="posts.length!=0">
            <div class="divide-y">
              <div v-for="post in posts.slice().reverse()" :key="post.id">
                <Topic :post="post"></Topic>
              </div>
            </div>
          </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import { mapGetters } from 'vuex'

  export default {
    data() {
      return {
        user: {},
        posts: []
      }
    },
    mounted() {
      window.vueInstance = this;
    },
    async asyncData({
      params,
      $axios,
      error
    }) {
      let username = params.username;
      if(username.startsWith('@')) username = username.substring(1);

      let user, posts;
      try {
        user = await $axios.$get(`/api/user/${username}`);
        posts = await $axios.$get(`/api/post/all/byUser/${user.id}`);
      } catch (err) {
        error({ statusCode: 404, message: err.message })
        return
      }
      
      return {
        user,
        posts
      }
    },
    computed: {
      ...mapGetters(['loggedInUser']),
      isCurrentUser() {
        return this.user.id == this.loggedInUser.id
      }
    },
    methods: {
      closed() {
        alert('modal closed');
      }
    },
  }

</script>
