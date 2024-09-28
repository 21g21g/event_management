<script setup>
import Home from "../../assets/icons/Home.vue"
import Ticket from "../../assets/icons/Ticket.vue"
import { useRouter } from 'vue-router';
import { ref } from "vue";

const router = useRouter()

</script>
<style scoped>
.router-link-exact-active {
  color: blueviolet !important;
}

</style>
<template>
  <div class="relative md:w-64 min-h-screen bg-slate-700  mt-16 text-white shadow-lg rounded-lg p-4">
    <div class="flex flex-col space-y-6">
      <NuxtLink 
        exact
        class="flex items-center gap-4 p-3 rounded-lg cursor-pointer transition duration-300 ease-in-out hover:bg-amber-600 hover:text-white"
        to="/user"
      >
        <Home class="w-6 h-6 bg-white"/>
        <span class="text-lg font-semibold">Home</span>
      </NuxtLink>

      <NuxtLink
        exact
        class="flex items-center gap-4 p-3 rounded-lg cursor-pointer transition duration-300 ease-in-out hover:bg-amber-600 hover:text-white"
        to="/user/ticketView"
      >
        <Ticket class="w-6 text-white h-6 bg-white"/>
        <span class="text-lg font-semibold">Ticket</span>
      </NuxtLink>
    </div>

   
  </div>
</template>
