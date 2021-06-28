export default {
  layout: "empty",
  data() {
    return {
      email: "",
      password: "",
      error: null,
    }
  },
  mounted() {
    window.vueInstance = this;
  },
  methods: {
    async login(e) {
      try {
        await this.$auth.loginWith('local', {
          data: {
            email: this.email,
            password: this.password
          }
        });

        this.$router.push('/topics')
      } catch (e) {
        console.log(e.response?.data || e);
        this.error = e.response?.data?.validationError || e.response?.data?.authError || 'An error occurred'
      }
    }
  },
}