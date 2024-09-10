import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        userId: null,
        
      }),

actions: {
    
    
    setUserId(id:any) {
      this.userId = id;
      console.log(this.userId);
    },
  }
 

});