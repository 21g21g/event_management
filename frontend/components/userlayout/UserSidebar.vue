<template>
<div class="w-28 min-h-screen mt-20  ">  
        <div class="flex hover:bg-amber-600 gap-2 cursor-pointer" @click="userhomeClick">
            <Home/>
        <span class="text-black">Home</span>
         </div>
      
      
</div>
</template>
<script setup>
import Home from "../../assets/icons/Home.vue"
import { useRouter } from 'vue-router';

const router=useRouter()
const userhomeClick=()=>{
      router.push("/user")
      
}
</script>