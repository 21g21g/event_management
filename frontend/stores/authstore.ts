import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        userId: null,
        isLoggedin:false,
        
      }),

actions: {
    
    
    setUserId(id:any) {
      this.userId = id;
      console.log(this.userId);
    },
    setIsloggedin(value:boolean){
      this.isLoggedin=value;
      console.log(this.isLoggedin)
    },
  }
 

});