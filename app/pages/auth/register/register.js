export default {
  auth: "guest",
  layout: "empty",
  data() {
    return {
      displayName: "",
      userName: "",
      email: "",
      password: "",
      error: null
    }
  },
  mounted() {
    window.vueInstance = this;
  },
  methods: {
    async register() {
      const {displayName, userName, email, password} = this

      try {
        const newUser = await this.$axios.$post('/api/auth/register', {
          displayName,
          userName,
          email,
          password
        })

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
