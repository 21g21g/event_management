<script setup>
import { useRouter } from 'vue-router';
import entertain from "../../assets/css/entairment.jpeg";
import sport from "../../assets/css/sport.jpeg";
import food from "../../assets/css/food.jpeg";
import health from "../../assets/css/health.jpeg";
import tech from "../../assets/css/tech.jpeg";
import education from "../../assets/css/education.jpeg";
import { useQuery } from '@vue/apollo-composable';
import { GET_ALL_EVENTS_WTHOUT_FILTER } from "../../utils/queries";
import { computed, onMounted } from 'vue';

const router = useRouter();

// Query
const { result, loading, error,refetch } = useQuery(GET_ALL_EVENTS_WTHOUT_FILTER);

// Cover images mapping
const coverImages = {
  Sport:sport,
  Entertainment: entertain,
  health,
  Food:food,
  Tech:tech,
  Education:education
};

// Get cover image for a category
const getCoverImage = (category) => coverImages[category] || '';

// Unique categories
const uniqueCategories = computed(() => {
  if(loading.value){
   return []
}else if(error.value){
  return []
}
else{
  const events = result.value?.events || [];
  if (!events.length) return [];
  return [
    ...new Map(
      events.map(event => [
        event.category, 
        { category: event.category, coverImage: getCoverImage(event.category) }
      ])
    ).values()
  ];
}
  
});
const handleClickShow = (category, image) => {
  router.push({
    path: '/catagoryDetail',
    query: {
      category,
      image
    }
  });
};

const fetchLatestEvents = async () => {
  if (loading.value) return; 
  try {
    await refetch();
  } catch (err) {
    console.error('Error refetching events:', err);
  }
};

onMounted(() => {
  fetchLatestEvents();
});

</script>
<template>
        <div >
        <h1 class="font-bold text-3xl p-4">Explore Event By Category</h1>
       <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-8 p-4">
         <h1 v-if="loading">Loading...</h1>
        <h1 v-else-if="error">{{ error.message }}</h1>     
        <div 
        v-else 
        v-for="event in uniqueCategories" 
        :key="event.category"
        class="cursor-pointer  rounded-lg border-2 hover:shadow-xl shadow-lg transition-shadow duration-300 ease-in-out overflow-hidden" 
        @click="handleClickShow(event.category, event.coverImage)"
      >
        <img 
          :src="event.coverImage" 
          alt="No image available" 
          class="w-full h-48 object-cover rounded-lg" 
        />
        <div class="p-4">
          <h1 class="text-center text-lg font-semibold text-gray-800">{{ event.category }}</h1>
        </div>
      </div>
    </div>
  </div>
</template>

