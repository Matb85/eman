// import i18n from './lang/i18n.js'
import head from "./config/headconfig.js";

export default {
  // server: { port: 3000, host: '192.168.1.64' },
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head,
  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: ["~assets/scss/global.scss"],
  styleResources: {
    scss: ["~assets/scss/_variables.scss"],
  },

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [],

  // Use all the new features from ES2015+ (https://nuxtjs.org/docs/2.x/configuration-glossary/configuration-modern/)
  modern: process.env.NODE_ENV === "production",

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: false,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    "@nuxt/typescript-build",
    // https://go.nuxtjs.dev/stylelint
    "@nuxtjs/stylelint-module",
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    "@nuxtjs/style-resources",

    "@nuxtjs/axios", // https://go.nuxtjs.dev/axios

    "@nuxtjs/auth-next", // https://auth.nuxtjs.org/

    // "@nuxtjs/pwa", // https://go.nuxtjs.dev/pwa

    // ["nuxt-i18n", i18n], // https://i18n.nuxtjs.org/

    ["nuxt-buefy", { css: false, defaultButtonRounded: true }], // https://go.nuxtjs.dev/buefy buttons are rounded!
  ],

  // Axios module configuration (https://go.nuxtjs.dev/config-axios)
  axios: {
    baseURL: process.env.AXIOS_BASE_URL,
  },

  auth: {
    localStorage: false,
    strategies: {
      local: {
        property: "token",
        user: {
          property: "data.user",
        },
        endpoints: {
          login: { url: "/auth/login", method: "post" },
          logout: false,
          user: {
            method: "post",
            url: "/graphql",
            data: { query: "query { user {email, firstName, lastName, enterprises}}" },
          },
        },
      },
    },
    redirect: {
      login: "/login",
      logout: "/register",
      callback: "/login",
      home: "/core",
    },
  },
  router: {
    middleware: ["auth", "enterprises"],
    base: process.env.ROUTER_BASE_URL || "/",
  },

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
    standalone: true,
    hardSource: true,
    parallel: true,
    cache: true,
    html: {
      minify: {
        minifyCSS: false,
        minifyJS: false,
      },
    },
  },
};
