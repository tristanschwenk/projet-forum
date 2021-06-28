<template>
<div class="page page-center">
      <div class="container-tight py-4">
        <div class="text-center mb-4">
          <a href="."><Logo /></a>
        </div>
        <form class="card card-md" method="post" autocomplete="off" @submit.prevent="sendData">
          <div class="card-body">
            <h2 class="card-title text-center mb-4">Create new password</h2>
            <div class="alert alert-danger" v-if="error">
              {{ error }}
            </div>
            <div class="alert alert-info" v-if="$auth.$state.redirect">
              You have to login before accessing to
              <strong>{{ $auth.$state.redirect }}</strong>
            </div>
            <div class="mb-3">
              <label class="form-label">Password</label>
              <input type="password" class="form-control" placeholder="Enter password" v-model="password" required>
            </div>

            <div class="mb-3">
              <label class="form-label">Confim Password</label>
              <input type="password" class="form-control" placeholder="Confirm password" v-model="confim" required>
            </div>
            <div class="form-footer">
              <button type="submit" class="btn btn-primary w-100">Send Request</button>
            </div>
          </div>
        </form>
        <div class="text-center text-muted mt-3">
          Don't have account yet? <nuxt-link to="/auth/register">Sign up</nuxt-link>
        </div>
      </div>
  </div>
  </template>

<script src="./passwordReset.js"></script>