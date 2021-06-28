<template>
  <div class="main">
    <div class="page-header">
      <div class="container-xl">
        <div class="row align-items-center">
          <div class="col">
            <div class="page-pretitle">
              Ezyo Forum
            </div>
            <h2 class="page-title">Topics</h2>
          </div>
          <div class="col-auto">
            <div class="btn-list">
              <a href="#" class="btn btn-white" data-bs-toggle="modal"
                data-bs-target="#modal-new-topic">
                <icon-plus />
                New
              </a>
            </div>
          </div>
        </div>
      </div>

    </div>

    <div class="page-body">
      <div class="container-xl">
        <div class="card">
          <div class="card-body">
            <div class="divide-y">
              <div v-for="topic in topics.slice().reverse()" :key="topic.id">
                <Topic :topic="topic"></Topic>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <NewEntityModal type="topic" id="modal-new-topic" ref="modal" @success="$fetch()"></NewEntityModal>
  </div>

</template>


<script>
  import {
    mapGetters
  } from "vuex";

  export default {
    head() {
      return {
        title: "Ezyo Forum"
      }
    },
    data() {
      return {
        topics: [],
      }
    },
    async fetch() {
      this.topics = await this.$axios.$get(`/api/topic`);
    },
    computed: {
      ...mapGetters(["isAuthenticated", "loggedInUser"]),
    },
  };

</script>
