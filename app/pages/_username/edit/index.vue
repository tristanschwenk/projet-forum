<template>
  <div class="page-body">
    <div class="container-xl">
      <div class="row">
        <div class="col-md-6 col-lg-3">
          <div class="card">
            <div class="card-body p-4 text-center">
              <span class="avatar avatar-xl mb-3 avatar-rounded"
                style="background-image: url(/images/avatar.jpg);"></span>
              <h3 class="m-0 mb-1">{{ user.displayName }}</h3>
              <div class="text-muted">@{{ user.userName }}</div>
            </div>
          </div>
        </div>
        <div class="col-md-6 col-lg-9">
          <div class="card">
                <div class="card-header">
                  <h3 class="card-title">Basic form</h3>
                </div>
                <div class="card-body">
                  <form @submit.prevent="save">
                    <div class="form-group mb-3 ">
                      <label class="form-label">Username</label>
                      <div>
                        <input type="text" class="form-control" aria-describedby="emailHelp" placeholder="Enter email" :value="user.userName" readonly>
                      </div>
                    </div><div class="form-group mb-3 ">
                      <label class="form-label">Display name</label>
                      <div>
                        <input type="text" class="form-control" aria-describedby="emailHelp" placeholder="Enter email" v-model="user.displayName">
                      </div>
                    </div>
                    <div class="form-group mb-3 ">
                      <label class="form-label">Email address</label>
                      <div>
                        <input type="email" class="form-control" aria-describedby="emailHelp" placeholder="Enter email" v-model="user.email">
                        <small class="form-hint">We'll never share your email with anyone else.</small>
                      </div>
                    </div>
                    <div class="form-footer">
                      <button type="submit" class="btn btn-primary">Submit</button>
                    </div>
                  </form>
                </div>
              </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        user: {}
      }
    },
    validate({params, $auth}) {
      let username = params.username;
      if(username.startsWith('@')) username = username.substring(1);
      return username == $auth.$state.user.userName
    },
    beforeMount() {
      this.user = {
        ...this.$auth.$state.user
      }
    },
    methods: {
      async save() {
        const res = await this.$axios.$patch('api/user/me', this.user);
        await this.$auth.fetchUser();
        this.$router.push(`/@${this.user.userName}`);
      }
    },
  }

</script>