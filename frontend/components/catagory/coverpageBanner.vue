<template>
  <div class="bg-blue-600 w-full h-96 flex flex-row justify-between items-center relative overflow-hidden">
    <div class="absolute inset-0 bg-gradient-to-r from-blue-600 via-blue-600 to-transparent"></div>

    <div class="relative  flex flex-col p-6 text-white max-w-sm">
      <h1 class="text-4xl font-bold mb-2">{{ category }} Events</h1>
      <p class="text-lg mb-4">Discover the best {{ category }} events in your area</p>
     
    </div>

    <img :src="image" alt="Category Image" class="relative  object-cover h-full w-1/2 rounded-l-lg shadow-lg transition-transform duration-500 hover:scale-105"/>
  </div>
</template>

<script setup>
defineProps({
  category: String,
  image: String,
})
</script>

