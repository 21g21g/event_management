// import {useAuthStore} from "../stores/authstore.ts"
export default defineNuxtRouteMiddleware((to,from)=>{
  if (process.client) {
    const token=localStorage.getItem("token")
  if (!token) {
      if(to.path==="/user" ||to.path==="/user/createEvent" ||to.path==="/user/bookMark" ||to.path==="/user/ticketView"){
        return navigateTo('/auth/login')
      }

    }
   
    
  }
})



