<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { useQuery } from '@vue/apollo-composable';
import { GET_ALL_EVENTS } from '../../utils/queries';
import { useRouter } from 'vue-router';

const router = useRouter();
const search = ref('');
const limit = ref(6);
const offset = ref(0); 
const totalEvents = ref(0); 
const errorMsg = ref('');  // Error message to handle any API error

const { result, loading, error, refetch } = useQuery(GET_ALL_EVENTS, {
  search: "%", 
  limit: limit.value,
  offset: offset.value,
});

const eventData = computed(() => {
  return result.value?.events || [];
});
console.log(eventData)

watch([search, limit], async() => {
  offset.value = 0; 
 await fetchEvents();
});

const fetchEvents =  () => {
  try {
    let searchQuery = search.value.trim() === "" ? "%" : `%${search.value}%`;
    const { data } =  refetch({ search: searchQuery, limit: limit.value, offset: offset.value });
    
    if (data && data.events) {
      totalEvents.value = data.events.length;
      errorMsg.value = ''; // Reset error message on successful fetch
    } else {
      throw new Error('Failed to fetch events data');
    }
  } catch (err) {
    console.error('Error fetching events:', err);
    errorMsg.value = 'Error fetching events. Please try again later.';
  }
};

const filteredData = computed(() => {
  return eventData.value;
});

const canShowMore = computed(() => {
  return totalEvents.value === limit.value && !loading.value;
});

const canShowLess = computed(() => {
  return limit.value > 6 && !loading.value;
});

const showMore = async () => {
  offset.value = limit.value; 
  limit.value += 4; 
  await fetchEvents();
};

const showLess = async () => {
  if (limit.value > 6) {
    limit.value -= 4; 
    offset.value = limit.value - 4;
    await fetchEvents();
  }
};

// Navigate to event detail page
const detailPage = (id) => {
  router.push(`/event/${id}`);
};

onMounted(async () => {
  await fetchEvents();
});
</script>

<style scoped>
.router-link-exact-active {
  color: blueviolet !important;
}
</style>
<template>
  <div class="flex flex-col w-full bg-gray-100 mt-10">
    <div class="relative flex items-center p-2 rounded-md h-16 justify-center shadow-md mx-auto max-w-4xl w-full">
      <input 
        type="text" 
        v-model="search" 
        placeholder="Search" 
        class="w-full p-4 border border-gray-300 rounded-full focus:ring focus:ring-blue-200 outline-none pr-16 transition-all duration-300"
      />
    </div>

    <div class="flex flex-col p-4 bg-gray-100">
      <h1 class="font-bold text-3xl mb-4 text-center">All Events</h1>
      
      <div v-if="filteredData.length === 0 && !loading" class="text-black ml-20">
        <p class="text-2xl text-orange-800 text-center">No events found. Please try a different search or check back later.</p>
      </div>
      
      <div class="grid gap-6 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        <div v-if="loading">Loading...</div>
        <div v-else-if="error">There is an error</div>
        
        <div 
          v-else
          v-for="event in filteredData" 
          :key="event.id"
          @click="detailPage(event.id)"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out cursor-pointer"
        >
          <img :src="event.featured_image" alt="Event image" class="w-full h-48 object-cover" />
          <div class="p-4">
            <p class="text-lg text-gray-800 mb-2"><span class="font-bold">Event Name</span>: {{ event.title }}</p>
            <p class="text-sm font-semibold text-gray-600"><span class="font-bold">Location</span>: {{ event.address }}</p>
            <p class="text-sm font-semibold text-gray-600">Preparation Date: For {{ event.preparation_date }}</p>
          </div>
        </div>
      </div>
      
      <div class="text-center flex flex-row gap-3 mt-6">
        <button 
          class="px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 transition-colors duration-300 mx-auto disabled:bg-gray-400 disabled:cursor-not-allowed"
          :disabled="!canShowMore"
          @click="showMore"
        >
          See More
        </button>
        
        <button 
          class="px-4 py-2 bg-blue-500 text-white font-semibold rounded-lg hover:bg-blue-600 transition-colors duration-300 mx-auto disabled:bg-gray-400 disabled:cursor-not-allowed"
          :disabled="!canShowLess"
          @click="showLess"
        >
          See Less
        </button>
      </div>
    </div>
  </div>
</template>
