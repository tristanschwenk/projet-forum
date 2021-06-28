export default {
  auth: "guest",
  layout: "empty",
  data() {
    return {
      password: "",
      confim: "",
      error: null
    }
  },
  async asyncData({
    params
  }) {
    const id = params.id
    return {
      id
    }
  },
  methods: {
    async sendData() {
      if (this.verifPass()) {
        try {
          const {
            password,
            id
          } = this
          const data = await this.$axios.$post('/api/auth/passwordreset', {
            password,
            id
          })
          return
        } catch(e) {
          console.log(e.response?.data || e);
          this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred'
          return
        }
      }
      this.error = "Both passwords are not identical"
    },
    verifPass() {
      return (this.password == this.confim)
    }
  }
}
