<script setup>
import { ref } from 'vue';
const isMenuOpen = ref(false);

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};
</script>

<style scoped>
.router-link-exact-active {
  color: blueviolet !important;
}
</style>
<template>
  <nav class="bg-white shadow-lg p-4">
    <div class="container mx-auto flex items-center justify-between">
      <NuxtLink to="/" exact class="text-2xl font-bold text-blue-600 hover:text-blue-800 transition duration-300">
        minabevent
      </NuxtLink>
      <div class="flex items-center space-x-4 ">
        <div class="hidden md:flex">
           <NuxtLink 
          to="/auth/login" 
          exact 
          class="text-blue-600 px-4 py-2 rounded-lg hover:bg-blue-100 transition duration-300"
        >
          Login
        </NuxtLink>
        <NuxtLink 
          to="/auth/signup" 
          exact 
          class="text-blue-600 hover:bg-blue-100  px-4 py-2 rounded-lg transition duration-300"
        >
          Signup
        </NuxtLink>
        </div>
       
        <button 
          @click="toggleMenu"
          class="md:hidden sm:block text-blue-600 hover:text-blue-800 transition duration-300 focus:outline-none"
        >
          <svg 
            class="w-6 h-6" 
            fill="none" 
            stroke="currentColor" 
            viewBox="0 0 24 24" 
            xmlns="http://www.w3.org/2000/svg"
          >
            <path 
              stroke-linecap="round" 
              stroke-linejoin="round" 
              stroke-width="2" 
              d="M4 6h16M4 12h16m-7 6h7" 
            />
          </svg>
        </button>
      </div>
    </div>
    <div 
      v-show="isMenuOpen"
      class="md:hidden absolute top-16 right-2 bg-white shadow-lg rounded-lg w-40"
    >
      <NuxtLink 
        to="/auth/login" 
        exact 
        class="block px-4 py-2 text-blue-600 hover:bg-blue-100"
      >
        Login
      </NuxtLink>
      <NuxtLink 
        to="/auth/signup" 
        exact 
        class="block px-4 py-2 text-blue-600 hover:bg-blue-100"
      >
        Signup
      </NuxtLink>
    </div>
  </nav>
</template>


