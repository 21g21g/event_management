<template>
  <div class="flex flex-col w-full bg-gray-100">
    <div class="relative flex items-center p-2 rounded-md h-16 justify-center shadow-md mx-auto max-w-4xl w-full mt-8">
      <input 
        type="text" 
        v-model="search" 
        placeholder="Search by location" 
        class="w-full p-4 border border-gray-300 rounded-full focus:ring focus:ring-blue-200 outline-none pr-16  transition-all duration-300"
      />
      <button
        name="search" 
        class="absolute right-3 top-1/2 transform -translate-y-1/2 cursor-pointer p-3 rounded-full text-white bg-blue-600 hover:bg-blue-700 active:bg-blue-800 focus:outline-none focus:ring-2 focus:ring-blue-300 focus:ring-offset-2 shadow-md transition-all duration-300"
        @click="serchLocation" 
      >
        <Search/>
      </button>
    </div>

    <div class="flex flex-col p-4 bg-gray-100 mt-10">
      <h1 class="font-bold text-3xl mb-4 text-center">All Events</h1> 
      <div class="grid gap-6 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        <div 
          v-for="dummy in allEvent" 
          :key="dummy.id"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out"
        >
          <img :src="dummy.featuredImage" alt="there is no image by this name" class="w-full h-48 object-cover cursor-pointer"/>
          <div class="p-4">
            <p class="text-lg font-semibold text-gray-800 mb-2">Name: {{dummy.eventName}}</p>
            <p class="text-sm font-semibold text-gray-600">Location: {{dummy.location}}</p>
            <p class="text-sm font-semibold text-gray-600">Length: For {{dummy.preparationTime}}</p>
          </div>
        </div>
      </div>

      <button 
        class="mt-8 px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 transition-colors duration-300 mx-auto"
        @click="showMore"
      >
        See More
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue"
import { DummyUser } from "../dummy"
import Search from "../../assets/icons/Search.vue"

const eventshowAtonce = ref(8)
const allEvent = computed(() => {
  return DummyUser.slice(0, eventshowAtonce.value)
})

const showMore = () => {
  eventshowAtonce.value += 4
}

const search = ref("")
const serchLocation = () => {
  console.log("Search clicked:", search.value)
}
</script>