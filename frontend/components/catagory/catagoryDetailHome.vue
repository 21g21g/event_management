<template>
  <div class="flex flex-col md:flex-row p-4 gap-6">
    <!-- Filters Section -->
    <div class="flex flex-col bg-white shadow-md p-4 rounded-lg w-full md:w-1/4">
      <h1 class="text-2xl font-bold mb-4">Filters</h1>
      <!-- Date Filters -->
      <div class="flex flex-col mb-4">
        <p class="text-lg font-semibold mb-2">Date</p>
        <div class="flex flex-row items-center mb-2">
          <input type="checkbox" id="today" v-model="filters.date.today" class="mr-2">
          <label for="today" class="cursor-pointer">Today</label>
        </div>
        <div class="flex flex-row items-center mb-2">
          <input type="checkbox" id="tomorrow" v-model="filters.date.tomorrow" class="mr-2">
          <label for="tomorrow" class="cursor-pointer">Tomorrow</label>
        </div>
        <div class="flex flex-row items-center">
          <input type="checkbox" id="thisWeek" v-model="filters.date.thisWeek" class="mr-2">
          <label for="thisWeek" class="cursor-pointer">This Week</label>
        </div>
      </div>
      <!-- Price Filters -->
      <div class="flex flex-col">
        <p class="text-lg font-semibold mb-2">Price</p>
        <div class="flex flex-row items-center mb-2">
          <input type="checkbox" id="free" v-model="filters.price.free" class="mr-2">
          <label for="free" class="cursor-pointer">Free</label>
        </div>
        <div class="flex flex-row items-center">
          <input type="checkbox" id="paid" v-model="filters.price.paid" class="mr-2">
          <label for="paid" class="cursor-pointer">Paid</label>
        </div>
      </div>
    </div>

    <!-- Events Section -->
    <div class="flex-1 min-h-screen p-4">
      <h1 class="text-3xl font-bold mb-6">Explore {{ category }} Events</h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
        <div
          v-for="dummy in filteredEvents"
          :key="dummy.eventName"
          class="transition-shadow duration-300 ease-in-out cursor-pointer bg-white shadow-md p-4 rounded-lg hover:shadow-lg"
          @click="handleClick(dummy.id)"
        >
          <div class="flex flex-col items-center">
            <img
              :src="dummy.featuredImage"
              alt="No image available"
              class="rounded-full w-24 h-24 object-cover mb-4"
            />
            <h3 class="text-center text-lg font-semibold text-gray-700">{{ dummy.eventName }}</h3>
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
defineProps({
    category:String
})
import { computed, ref } from "vue"
import {DummyUser} from "../dummy"
import { useRouter } from 'vue-router';
const router=useRouter()
const filters = ref({
  date: {
    today: false,
    tomorrow: false,
    thisWeek: false,
  },
  price: {
    free: false,
    paid: false,
  },
});

// Computed property to filter events based on selected filters
const filteredEvents = computed(() => {
  return DummyUser.filter((event) => {
    let dateMatch = true;
    let priceMatch = true;

    // Date filter logic (add actual logic to match event dates)
    if (filters.value.date.today) {
      dateMatch = true; // Replace with actual date comparison logic
    } else if (filters.value.date.tomorrow) {
      dateMatch = true; // Replace with actual date comparison logic
    } else if (filters.value.date.thisWeek) {
      dateMatch = true; // Replace with actual date comparison logic
    }

    // Price filter logic (replace `event.isFree` with actual property)
    if (filters.value.price.free && !event.isFree) {
      priceMatch = false;
    }
    if (filters.value.price.paid && event.isFree) {
      priceMatch = false;
    }

    return dateMatch && priceMatch;
  }).slice(0, showOnce.value);
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
</script>
// by accepting catagory send http request to the backend  filtering by category.