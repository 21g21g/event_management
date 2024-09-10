// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    "vuetify-nuxt-module",
    '@nuxt/icon',
    '@pinia/nuxt'
  ],
  plugins: [
    './plugins/apollo.client.ts'  
  
  ],
  
  vuetify: {
    moduleOptions: {
      /* module specific options */
    },
    vuetifyOptions: {
      /* vuetify options */
    }
  },
  vite: {
    plugins: [
      require('vite-plugin-vue-inspector')({
        toggleButtonVisibility: 'always', // Example option
      }),
    ],
  },
  
  
  })