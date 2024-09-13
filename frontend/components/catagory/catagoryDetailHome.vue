<template>
  <div class="flex flex-col md:flex-row p-4 gap-6">
    <!-- Filters Section -->
    <div class="flex flex-col bg-white shadow-md p-4 rounded-lg w-full md:w-1/4">
      <h1 class="text-2xl font-bold mb-4">Filters</h1>
      <!-- Date Filters -->
      <div class="flex flex-col mb-4">
  <div class="flex flex-row items-center mb-2">
    <label for="title" class="cursor-pointer">Title</label>
    <input
      type="text"
      v-model="filters.title"
      id="title"
      class="mr-2 w-40 border border-gray-300 rounded p-2 focus:border-blue-500 focus:outline-none"
    />
  </div>
  <div class="flex flex-row items-center mb-2">
    <label for="venue" class="cursor-pointer">Venue</label>
    <input
      type="text"
      v-model="filters.venue"
      id="venue"
      class="mr-2 w-40 border border-gray-300 rounded p-2 focus:border-blue-500 focus:outline-none"
    />
  </div>
</div>
      <!-- Price Filters -->
      <div class="flex flex-col">
        <p class="text-lg font-semibold mb-2">Price</p>
        <!-- <div class="flex flex-row items-center mb-2">
          <input type="checkbox" id="free" v-model="filters.price.free" class="mr-2">
          <label for="free" class="cursor-pointer">Free</label>
        </div>
        <div class="flex flex-row items-center">
          <input type="checkbox" id="paid" v-model="filters.price.paid" class="mr-2">
          <label for="paid" class="cursor-pointer">Paid</label>
        </div> -->
         <div class="flex flex-row items-center mb-2">
        <input type="checkbox" id="free" v-model="filters.price.free" class="mr-2" />
        <label for="free" class="cursor-pointer">Free</label>
      </div>
      <div class="flex flex-row items-center">
        <input type="checkbox" id="paid" v-model="filters.price.paid" class="mr-2" />
        <label for="paid" class="cursor-pointer">Paid</label>
      </div>
      </div>
    </div>
     
    <!-- Events Section -->
    <div class="flex-1 min-h-screen p-4">
      <h1 class="text-3xl font-bold mb-6">Explore {{ category }} Events</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
        <h1 v-if="loading">Loading...</h1>
        <h1 v-else-if="error">there is ocuured error</h1>
        <div
        v-else
          v-for="event in filteredEvents"
          :key="event.id"
          class="transition-shadow duration-300 ease-in-out cursor-pointer bg-white shadow-md p-4 rounded-lg hover:shadow-lg"
          @click="handleClick(event.id)"
        >
          <div class="flex flex-col items-center">
            <img
              :src="event.featured_image"
              alt="No image available"
              class="rounded-full w-24 h-24 object-cover mb-4"
            />
            <h3 class="text-center text-sm font-semibold text-gray-700">Title:{{ event.title }}</h3>
             <h3 class="text-center text-sm font-semibold text-gray-700">Venue:{{ event.venue }}</h3>
            <p class="text-center text-base font-semibold text-gray-700">price:{{ event.price }}</p>


          </div>
        </div>
      </div>

      <!-- View More Button -->
      <div class="flex justify-center mt-8">
        <button
          @click="viewMoreButton"
          class="bg-blue-500 text-white font-semibold py-2 px-6 rounded-full hover:bg-blue-600 transition duration-300 shadow-md hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-opacity-50"
        >
          View More
        </button>
      </div>
    </div>
  </div>
</template>
<script setup>
const props=defineProps({
    category:String
})
const { category } = props;

import { computed, onMounted, ref } from "vue"
import {DummyUser} from "../dummy"
import { useRouter, useRoute } from 'vue-router';
import { useQuery } from "@vue/apollo-composable";
import {GET_EVENT_BY_CATEGORY} from "../graphql/queries"

const {result,loading,error,refetch}=useQuery(GET_EVENT_BY_CATEGORY,{category})
const router=useRouter()
const filters = ref({
  
    title: '',
    venue: '',
  price: {
    free: false,
    paid: false,
  },
});

// Computed property to filter events based on selected filters
const eventData = computed(() => {
  if (loading.value) return [];
  if (error.value) return [];
  return Array.isArray(result.value?.events) ? result.value.events : [];
});
const filteredEvents = computed(() => {
  return eventData.value.filter((event) => {
    const matchesTitle = filters.value.title
      ? event.title.toLowerCase().includes(filters.value.title.toLowerCase())
      : true;
    const matchesVenue = filters.value.venue
      ? event.venue.toLowerCase().includes(filters.value.venue.toLowerCase())
      : true;
    const matchesPrice = (filters.value.price.free && event.price === 'free') ||
      (filters.value.price.paid && event.price ==='paid') ||
      (!filters.value.price.free && !filters.value.price.paid); 

    return matchesTitle && matchesVenue && matchesPrice;
  });
});




const showOnce=ref(10)
const showAtOnce=computed(()=>{
    return DummyUser.slice(0,showOnce.value)
})
const viewMoreButton=()=>{
    showOnce.value +=5

}
const handleClick=(id)=>{
  console.log(id)
  router.push(`event/${id}`,{params:{id}})
}
onMounted(()=>{
  refetch()

})
</script>
