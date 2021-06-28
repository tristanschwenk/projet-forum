<template>
  <div>
    <header class="navbar navbar-expand-md navbar-light d-print-none">
      <div class="container-xl">
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu">
          <span class="navbar-toggler-icon"></span>
        </button>
        <h1 class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
          <a class="navbar-brand" href="/">Ezyo Forum</a>
        </h1>
        <div class="navbar-nav flex-row order-md-last">  
          <div class="nav-item dropdown d-none d-md-flex me-3">
            <a href="#" class="nav-link px-0" data-bs-toggle="dropdown" tabindex="-1" aria-label="Show notifications">
              <!-- Download SVG icon from http://tabler-icons.io/i/bell -->
              <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                <path d="M10 5a2 2 0 0 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6"></path>
                <path d="M9 17v1a3 3 0 0 0 6 0v-1"></path>
              </svg>
              <span class="badge bg-red"></span>
            </a>
            <div class="dropdown-menu dropdown-menu-end dropdown-menu-card">
              <div class="card">
                <div class="card-body">
                  <div class="divide-y">
                    <div v-for="notif in notifications" :key="notif.id"></div>
                    <NuxtLink to="/notifications">Voir plus</NuxtLink>
                  </div>
                  
                </div>
              </div>
            </div>
          </div>
          <div class="nav-item dropdown">
            <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown"
              aria-label="Open user menu">
              <span class="avatar avatar-sm"
                >{{getInitials}}</span>
              <div class="d-none d-xl-block ps-2">
                <div>{{loggedInUser.displayName}}</div>
                <div class="mt-1 small text-muted">@{{loggedInUser.userName}}</div>
              </div>
            </a>
            <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
              <nuxt-link :to="'/@'+loggedInUser.userName" class="dropdown-item">Profile &amp; account</nuxt-link>
              <a href="#" class="dropdown-item">Feedback</a>
              <div class="dropdown-divider"></div>
              <a href="#" class="dropdown-item">Settings</a>
              <a href="#" class="dropdown-item" @click.prevent="logout">Logout</a>
            </div>
          </div>
        </div>
      </div>
    </header>
    <div class="navbar-expand-md">
      <div class="collapse navbar-collapse" id="navbar-menu">
        <div class="navbar navbar-light">
          <div class="container-xl">
            <ul class="navbar-nav">
              <li class="nav-item">
                <nuxt-link to="/topics" class="nav-link">
                  <span class="nav-link-icon d-md-none d-lg-inline-block">
                    <!-- Download SVG icon from http://tabler-icons.io/i/home -->
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24"
                      stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                      <polyline points="5 12 3 12 12 3 21 12 19 12"></polyline>
                      <path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7"></path>
                      <path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                    </svg>
                  </span>
                  <span class="nav-link-title">
                    Topics
                  </span>
                </nuxt-link>
              </li>
            </ul>
            <Searchbar />
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
  import {
    mapGetters
  } from 'vuex'

  export default {
    async asyncData({params, $axios}) {
      const notifications = await $axios.$get(`/api/notification/?`);
      return {
        notifications
      }
    },
    mounted() {
      console.log(this.notifications);
    },
    computed: {
      ...mapGetters(['isAuthenticated', 'loggedInUser']),
      profilePic() {
        const avatarUrl = encodeURI(`https://eu.ui-avatars.com/api/?name=${this.$auth.user.displayName}`);
        return {
          backgroundImage: `url(${avatarUrl});`
        }
      },
      getInitials() {
        let temp = this.loggedInUser.displayName.split(' ')

        if (temp.length == 0) return ''
        if (temp.length == 1) return temp[0].slice(0,2).toUpperCase()
        return (temp[0].slice(0,1)+temp[1].slice(0,1)).toUpperCase()
        
      }
    },
    methods: {
      async logout() {
        await this.$auth.logout();
        this.$router.push('/auth/login')
      },
    },
  }

</script>
