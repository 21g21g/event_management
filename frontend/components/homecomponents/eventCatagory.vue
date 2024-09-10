<template>
  <div class="bg-gray-100 min-h-screen">
    <h1 class="font-bold text-3xl p-4">Explore Event By Category</h1>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8 p-4">
      <h1 v-if="loading">Loading...</h1>
      <h1 v-else-if="error">There is an error</h1>
      <div 
        v-else 
        v-for="event in uniqueCategories" 
        :key="event.category"
        class="cursor-pointer bg-white rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300 ease-in-out overflow-hidden" 
        @click="handleClickShow(event.category, event.coverImage)"
      >
        <img 
          :src="event.coverImage" 
          alt="No image available" 
          class="w-full h-48 object-cover rounded-t-lg" 
        />
        <div class="p-4">
          <h1 class="text-center text-lg font-semibold text-gray-800">{{ event.category }}</h1>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import entertain from "../../assets/css/entairment.jpeg";
import sport from "../../assets/css/sport.jpeg";
import food from "../../assets/css/food.jpeg";
import health from "../../assets/css/health.jpeg";
import tech from "../../assets/css/tech.jpeg";
import education from "../../assets/css/education.jpeg";
import { useQuery } from '@vue/apollo-composable';
import { GET_ALL_EVENTS } from "../graphql/queries";
import { computed, onMounted } from 'vue';

// Router instance
const router = useRouter();

// Query
const { result, loading, error,refetch } = useQuery(GET_ALL_EVENTS);

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
});
const fetchLatestEvents = async () => {
  try {
    await refetch();
  } catch (err) {
    console.error('Error refetching events:', err);
  }
};

// Handle category click
const handleClickShow = (category, image) => {
  router.push({
    path: '/categoryDetail',
    query: {
      category,
      image
    }
  });
};
onMounted(() => {
  fetchLatestEvents();
});
</script>