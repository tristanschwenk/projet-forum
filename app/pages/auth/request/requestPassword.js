export default {
  auth: "guest",
  layout: "empty",
  data() {
    return {
      email: "",
      error: null,
    }
  },
  mounted() {
    window.vueInstance = this;
  },
  methods: {
    async passwordRequest(e) {
      try {
        const {email} = this

        const newRequest = await this.$axios.$post('/api/auth/passwordrequest', {
          email,
        })

        
      } catch (e) {
        console.log(e.response?.data || e);
        this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred'
      }
    }
  },
}