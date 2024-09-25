
// import {useAuthStore} from "../stores/authstore.ts"
export default defineNuxtRouteMiddleware((to,from)=>{
  console.log(to.path)
  
    if (process.client) {
      const token=localStorage.getItem("token")
    if (token) {
      if (to.path === '/auth/login' || to.path === '/auth/signup' ) {
        return navigateTo('/user')
      }
    }
  }

})


