<template>
  <div class="page page-center">
      <div class="container-tight py-4">
        <div class="text-center mb-4">
          <a href="."><Logo /></a>
        </div>
        <form class="card card-md" method="post" autocomplete="off" @submit.prevent="login">
          <div class="card-body">
            <h2 class="card-title text-center mb-4">Login to your account</h2>
            <div class="alert alert-danger" v-if="error">
              {{ error }}
            </div>
            <div class="alert alert-info" v-if="$auth.$state.redirect">
              You have to login before accessing to
              <strong>{{ $auth.$state.redirect }}</strong>
            </div>
            <div class="mb-3">
              <label class="form-label">Email address</label>
              <input type="email" class="form-control" placeholder="Enter email" v-model="email" required>
            </div>
            <div class="mb-2">
              <label class="form-label">
                Password
              </label>
              <div class="input-group input-group-flat">
                <input type="password" class="form-control"  placeholder="Password" v-model="password" required>
                <span class="input-group-text">
                  <a href="#" class="link-secondary" title="Show password" data-bs-toggle="tooltip"><!-- Download SVG icon from http://tabler-icons.io/i/eye -->
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><circle cx="12" cy="12" r="2" /><path d="M22 12c-2.667 4.667 -6 7 -10 7s-7.333 -2.333 -10 -7c2.667 -4.667 6 -7 10 -7s7.333 2.333 10 7" /></svg>
                  </a>
                </span>
              </div>
            </div>
            <div class="form-footer">
              <button type="submit" class="btn btn-primary w-100">Sign in</button>
            </div>
          </div>
          <div class="hr-text">or</div>
          <div class="card-body">
            <div class="row">
              <div class="col">
                <a @click.prevent="$auth.loginWith('github')" href="#" class="btn btn-white w-100">
                  <!-- Download SVG icon from http://tabler-icons.io/i/brand-github -->
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon text-github" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M9 19c-4.3 1.4 -4.3 -2.5 -6 -3m12 5v-3.5c0 -1 .1 -1.4 -.5 -2c2.8 -.3 5.5 -1.4 5.5 -6a4.6 4.6 0 0 0 -1.3 -3.2a4.2 4.2 0 0 0 -.1 -3.2s-1.1 -.3 -3.5 1.3a12.3 12.3 0 0 0 -6.2 0c-2.4 -1.6 -3.5 -1.3 -3.5 -1.3a4.2 4.2 0 0 0 -.1 3.2a4.6 4.6 0 0 0 -1.3 3.2c0 4.6 2.7 5.7 5.5 6c-.6 .6 -.6 1.2 -.5 2v3.5" /></svg>
                  Login with Github
                </a>
              </div>
            </div>
          </div>
        </form>
        <div class="text-center text-muted mt-3">
          Don't have account yet? <nuxt-link to="/auth/register">Sign up</nuxt-link>
        </div>
        <div class="text-center text-muted mt-3">
        Forgot your password? <nuxt-link to="/auth/request">Request a new one</nuxt-link>
        </div>
      </div>
  </div>
</template>

<script src="./login.js"></script>
