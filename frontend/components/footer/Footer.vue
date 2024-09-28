<script setup>
import { useRouter } from "vue-router";



</script>
<style scoped>
.router-link-exact-active {
  color: blueviolet !important;
}
</style>
<template>
  <footer class="bg-slate-600 text-white py-4">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0">
          <h2 class="text-2xl font-bold">MinabEvent</h2>
          
        </div>

        <div class="flex space-x-8">
          <NuxtLink to="/" class="text-gray-400 hover:text-white">Home</NuxtLink>
          <NuxtLink to="/events" class="text-gray-400 hover:text-white">Events</NuxtLink>
          <NuxtLink to="/about" class="text-gray-400 hover:text-white">About</NuxtLink>
          <NuxtLink to="/contact" class="text-gray-400 hover:text-white">Contact</NuxtLink>
        </div>
      </div>

      <div class="flex justify-center space-x-4 mt-8">
        <NuxtLink to="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-facebook" class="w-6 h-6" />
        </NuxtLink>
        <NuxtLink to="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-twitter" class="w-6 h-6" />
        </NuxtLink>
        <NuxtLink to="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-instagram" class="w-6 h-6" />
        </NuxtLink>
        <NuxtLink to="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-linkedin" class="w-6 h-6" />
        </NuxtLink>
      </div>

      <!-- Footer Bottom: Copyright and Legal Links -->
      <!-- <div class="flex flex-col md:flex-row justify-between items-center mt-8">
        <div class="flex space-x-4">
          <NuxtLink to="/terms" class="text-gray-400 hover:text-white">Terms of Service</NuxtLink>
          <NuxtLink to="/privacy" class="text-gray-400 hover:text-white">Privacy Policy</NuxtLink>
        </div>
      </div> -->
    </div>
  </footer>
</template>

