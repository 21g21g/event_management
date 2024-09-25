// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    "vuetify-nuxt-module",
    '@nuxt/icon',

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
  
    server:{
        hmr:true,
      
        
        
    },
   
  },
  
  // ssr:true,
  // runtimeConfig:{
  //   public:{
  //     hasuraAdminSecret:process.env.HASURA_ADMIN_SECRET
  //   }
  // }
  
  })