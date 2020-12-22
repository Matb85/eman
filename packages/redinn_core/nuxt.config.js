//import i18n from './lang/i18n.js'
import head from "./config/headconfig.js"
export default {
  //server: { port: 3000, host: '192.168.1.64' },
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: head,
  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: ["~assets/scss/global.scss"],
  styleResources: {
    scss: ["~assets/scss/variables.scss"],
  },
  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [],

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
    // https://go.nuxtjs.dev/buefy
    "nuxt-buefy",
    // https://go.nuxtjs.dev/axios
    "@nuxtjs/axios",
    // https://go.nuxtjs.dev/pwa
    //'@nuxtjs/pwa',
    // https://i18n.nuxtjs.org/
    //['nuxt-i18n', i18n],
  ],

  // Axios module configuration (https://go.nuxtjs.dev/config-axios)
  axios: {},
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
}
/*
<title>Redinn_Core</title>
<meta data-n-head="ssr" charset="utf-8">
<meta data-n-head="ssr" name="viewport" content="width=device-width, initial-scale=1">
<meta data-n-head="ssr" data-hid="description" name="description" content="">
<meta data-n-head="ssr" data-hid="charset" charset="utf-8">
<meta data-n-head="ssr" data-hid="mobile-web-app-capable" name="mobile-web-app-capable" content="yes">
<meta data-n-head="ssr" data-hid="apple-mobile-web-app-title" name="apple-mobile-web-app-title" content="redinn_core">
<meta data-n-head="ssr" data-hid="og:type" name="og:type" property="og:type" content="website">
<meta data-n-head="ssr" data-hid="og:title" name="og:title" property="og:title" content="redinn_core">
<meta data-n-head="ssr" data-hid="og:site_name" name="og:site_name" property="og:site_name" content="redinn_core">
<meta data-n-head="ssr" data-hid="og:description" name="og:description" property="og:description" content="## Build Setup">
<link data-n-head="ssr" rel="icon" type="image/x-icon" href="/favicon.ico">
<link data-n-head="ssr" rel="stylesheet" type="text/css" href="//cdn.materialdesignicons.com/5.0.45/css/materialdesignicons.min.css">
<link data-n-head="ssr" rel="shortcut icon" href="/_nuxt/icons/icon_64x64.5f6a36.png">
<link data-n-head="ssr" rel="apple-touch-icon" href="/_nuxt/icons/icon_512x512.5f6a36.png" sizes="512x512">
<link data-n-head="ssr" rel="manifest" href="/_nuxt/manifest.e352a634.json" data-hid="manifest">
<link rel="preload" href="/_nuxt/runtime.js" as="script">
<link rel="preload" href="/_nuxt/commons/app.js" as="script">
<link rel="preload" href="/_nuxt/vendors/app.js" as="script">
<link rel="preload" href="/_nuxt/app.js" as="script">
<link rel="preload" href="/_nuxt/pages/media/addstory/upload.js" as="script">
<link rel="preload" href="/_nuxt/vendors/pages/media/addstory/upload.js" as="script">
<style data-vue-ssr-id="797461a0:0">/*! Buefy v0.9.3 | MIT License | github.com/buefy/buefy */
