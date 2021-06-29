require('dotenv').config()

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'app',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ],
    script: []
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    "@/node_modules/@tabler/core/dist/css/tabler",
    "@/assets/scss/main",
    "@/node_modules/quill/dist/quill.core.css",
    "@/node_modules/quill/dist/quill.snow.css",
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: '@/node_modules/@tabler/core/dist/js/tabler.js', mode: 'client' },
    '~/plugins/relativeTime.js',
    '~/plugins/parseDelta.js',
    '~/plugins/getInitials.js',
    {src: '~/plugins/quill.js', mode: 'client'},
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    // '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/stylelint
    '@nuxtjs/stylelint-module',
    '@nuxtjs/dotenv'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    '@nuxtjs/auth-next',
    '@nuxtjs/dayjs',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    proxy: true
  },

  proxy: {
    '/api/': {
      target: 'http://localhost:8080',
      pathRewrite: { '^/api/': '' }
    }
  },

  router: {
    middleware: ['auth']
  },

  auth: {
    plugins: [
      "~/plugins/auth.js"
    ],
    redirect: {
      login: '/auth/login',
      logout: '/auth/login',
      callback: '/auth/login',
      home: '/topics'
    },
    localStorage: false,
    cookie: {
      prefix: 'auth.',
      options: {
        path: '/',
        maxAge: 10800
      }
    },
    strategies: {
      local: {
        scheme: 'refresh',
        token: {
          property: 'accessToken',
          maxAge: 300,
        },
        refreshToken: {
          property: 'refreshToken',
          data: 'refreshToken',
          maxAge: 900
        },
        user: {
          property: "user",
          // autoFetch: false
        },
        endpoints: {
          login: { url: 'api/auth/login', method: 'post'},
          refresh: { url: 'api/auth/refresh', method: 'post'},
          logout: false,
          user: { url: '/api/user/me', method: 'get'},
        }
      },
      github: {
        clientId: process.env.GITHUB_ID,
        clientSecret: process.env.GITHUB_SECRET,
      },
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },

  dayjs: {
    plugins: [
      'relativeTime',
    ]
  }
}
