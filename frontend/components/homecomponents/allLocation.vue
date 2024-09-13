<template>
  <div class="flex flex-col w-full bg-gray-100 mt-10 ">
    <div class="relative flex items-center p-2 rounded-md h-16 justify-center shadow-md mx-auto max-w-4xl w-full">
      <input 
        type="text" 
        v-model="search" 
        placeholder="Search by location" 
        class="w-full p-4 border border-gray-300 rounded-full focus:ring focus:ring-blue-200 outline-none pr-16 transition-all duration-300"
      />
     
    </div>

    <div class="flex flex-col p-4 bg-gray-100">
      <h1 class="font-bold text-3xl mb-4 text-center">All Events</h1> 
       <div v-if="filteredData.length===0" class="text-black ml-20">
        <p class="text-2xl text-orange-800 text-center">No events found. Please try a different search or check back later.</p>
          </div>
      <div class="grid gap-6 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">There is an error</div>
        
        <div 
          v-else
          v-for="event in filteredData.slice(0, eventshowAtonce)" 
          :key="event.id"
          @click="detailPage(event.id)"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out"
        >
          <img :src="event.featured_image" alt="Event image" class="w-full h-48 object-cover cursor-pointer"/>
          <div class="p-4">
            <p class="text-lg font-semibold text-gray-800 mb-2">Event Name: {{event.title}}</p>
            <p class="text-sm font-semibold text-gray-600">Location: {{event.address}}</p>
            <p class="text-sm font-semibold text-gray-600">Preparation Date: For {{event.preparation_date}}</p>
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
import { ref, computed, onMounted } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import { GET_ALL_EVENTS } from '../graphql/queries';
import { useRouter } from 'vue-router';
const router=useRouter()
const search = ref('');
const eventshowAtonce = ref(8);
const { result, loading, error, refetch } = useQuery(GET_ALL_EVENTS);
// Compute eventData based on query result
const eventData = computed(() => {
  if (loading.value) return [];
  if (error.value) return [];
  return Array.isArray(result.value?.events) ? result.value.events : [];
});

// Filtered data based on search
const filteredData = computed(() => {
  if (!search.value.trim()) {
    return eventData.value;
  }
  return eventData.value.filter((event) =>
    event.address.toLowerCase().includes(search.value.toLowerCase())
  );
});

// Function to load more events
const showMore = () => {
  eventshowAtonce.value += 4;
};

// Navigate to event detail page
const detailPage = (id) => {
  router.push(`event/${id}`);
};

onMounted(()=>{
  refetch()
})
</script>
