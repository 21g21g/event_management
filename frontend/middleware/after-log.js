
import {useAuthStore} from "../stores/authstore.ts"
export default defineNuxtRouteMiddleware((to,from)=>{
  console.log(to.path)
    const authStore=useAuthStore()
    if (authStore.isLoggedin) {
      if (to.path === '/auth/login' || to.path === '/auth/signup' ||to.path==='/') {
        return navigateTo('/user')
      }
    }

})