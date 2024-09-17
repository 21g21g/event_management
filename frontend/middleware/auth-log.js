import {useAuthStore} from "../stores/authstore.ts"
export default defineNuxtRouteMiddleware((to,from)=>{
    const authStore=useAuthStore()

    if(!authStore.isLoggedin){
      if(to.path==="/user" ||to.path==="/user/createEvent" ||to.path==="/user/bookMark" ||to.path==="/user/ticketView"){
        return navigateTo('/auth/login')
      }

    }
   
    

})