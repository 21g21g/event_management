<script setup>
const props=defineProps({
    category:String
})
const { category } = props;

import { computed, onMounted, ref, watch } from "vue"
import { useRouter, useRoute } from 'vue-router';
import { useQuery } from "@vue/apollo-composable";
const search=ref("")
const limit=ref(6)
// const totalEvents = ref(0);
const filters = ref({
  price: {
    free: false,
    paid: false,
  },
});
import {SEARCH_TERMS} from "../../utils/queries"
// const {result,loading,error,refetch}=useQuery(GET_EVENT_BY_CATEGORY,{category});
const {result,loading,error,refetch}=useQuery(SEARCH_TERMS,{search:"%",category:category,limit:limit.value})
const router=useRouter()

watch([search,limit],()=>{
  const searchTerm=search.value.trim()===""?"%":`%${search.value}%`
 refetch({search:searchTerm,limit:limit.value})


})
const total=computed(()=>{
  return result.value?.events_aggregate.aggregate.count

})

const eventData = computed(() => {
  return result.value?.events ? result.value.events : [];
});
const filteredEvents = computed(() => {
  return eventData.value.filter((event) => {

    const matchesPrice = (filters.value.price.free && event.price === 'free') ||
      (filters.value.price.paid && event.price ==='paid') ||
      (!filters.value.price.free && !filters.value.price.paid); 

    return matchesPrice;
  });
});


const viewMoreButton=()=>{
    limit.value +=3

}
const canViewMore = computed(() => {
  return limit.value < total.value;
});
const handleClick=(id)=>{
  console.log(id)
  router.push(`event/${id}`,{params:{id}})
}
onMounted(() => {
 

   refetch();
});
</script>

<template>
  <div class="flex flex-col md:flex-row p-4 gap-6">
    <div class="flex flex-col bg-white shadow-md p-4 rounded-lg w-full md:w-1/4">
      <h1 class="text-2xl font-bold mb-4">Filters</h1>
      <div class="flex flex-col mb-4">
  <div class="flex flex-row items-center mb-2">
    <label for="title" class="cursor-pointer">Search</label>
    <input
      type="text"
      v-model="search"
      id="title"
      class="mr-2 w-40 border border-gray-300 rounded p-2 focus:border-blue-500 focus:outline-none"
    />
  </div>
  
</div>
      <div class="flex flex-col">
        <p class="text-lg font-semibold mb-2">Price</p>
      
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
       
     
        <div class="flex-1 min-h-screen p-4">
      <h1 class="text-3xl font-bold mb-6">Explore {{ category }} Events</h1>
        <div v-if="filteredEvents.length === 0 && !loading" class="text-black ml-20">
        <p class="text-2xl text-orange-800 text-center">No events found. Please try a different search or check back later.</p>
      </div>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
        <h1 v-if="loading">Loading...</h1>
        <h1 v-else-if="error">there is ocuured error</h1>
        <div
        v-else
          v-for="event in filteredEvents"
          :key="event.id"
          class="transition-shadow duration-300 ease-in-out cursor-pointer  p-4 rounded-lg "
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

      <div class="flex justify-center mt-8">
        <button
          @click="viewMoreButton"
          :disabled=!canViewMore
          class="bg-blue-500 text-white font-semibold py-2 px-6 rounded-full disabled:bg-gray-400 disabled:cursor-not-allowed hover:bg-blue-600 transition duration-300 shadow-md hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-opacity-50"
        >
          View More
        </button>
      </div>
    </div>
  </div>
</template>
