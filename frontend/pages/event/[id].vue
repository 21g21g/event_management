<template>
  <div class="flex flex-col mt-20">
    <div class="w-full max-w-3xl mx-auto mt-10 relative "
     @mouseover="showHeart" 
     @mouseleave="hideHeart"
     >
      <v-carousel
        :items-to-show="3"
        :items-to-scroll="1"
        cycle
        class="h-auto"
      >
        <v-carousel-item
          v-for="(item, i) in eventData?.imagestores"
          :key="i"
          class="flex items-center justify-center w-full hover:scale-105 transition-transform duration-300"
        >
          <img
            :src="item.url"
            :alt="item.alt"
            class="object-cover w-full h-64 rounded-md shadow-lg"
          />
        </v-carousel-item>
      </v-carousel>
      
    
      
    </div>
    
 <Heart 
        v-if="heart"  
        @click="toggleBookmark" 
        class="absolute w-12 h-12 cursor-pointer transition-transform duration-300 transform hover:scale-125"
        style="top: 30%; left: 50%;" 
      />    <div class="flex flex-col md:flex-row justify-between w-full p-4 mt-5">
      <h1 v-if="loading">Loading...</h1>
      <h1 v-else-if="error">There is an Error</h1>
      <div v-else class="flex flex-col w-full md:w-2/3 space-y-4">
        <h1 class="text-2xl font-bold text-gray-800">Start Date: {{ eventData?.preparation_date }}</h1>
        <p class="text-xl font-semibold text-gray-700">Title: {{ eventData?.title }}</p>
        <h2 class="text-xl font-semibold text-gray-600">Location: {{ eventData?.address }}</h2>
        <h3 class="text-lg font-medium text-gray-600">About the Event</h3>
        <p class="text-gray-600">{{ eventData?.description }}</p>
      </div>

      <div class="border-2 border-gray-300 flex flex-col w-full md:w-1/3 p-4 space-y-4 rounded-md">
        <div class="flex flex-col space-y-2">
          <div class="flex items-center justify-between">
            <p class="text-lg font-semibold">Number of Ticket</p>
            <div class="flex items-center space-x-2">
              <button 
                @click="decreaseCount"
                :disabled="count <= 1"
                class="bg-gray-200 hover:bg-gray-300 rounded-md p-1"
              >
                <span class="text-xl">-</span>
              </button>
              <p class="text-lg">{{ count }}</p>
              <button
                @click="increaseCount"
                class="bg-gray-200 hover:bg-gray-300 rounded-md p-1"
              >
                <span class="text-xl">+</span>
              </button>
            </div>
          </div>
          <p v-if="eventData?.price==='paid'" class="text-xl font-bold">{{ eventData?.specific_price}}</p>
          <p v-else>{{eventData?.price}}</p>
        </div>
        <button
          @click="reserveTicket"
          class="bg-blue-500 text-white hover:bg-blue-600 rounded-md py-2 px-4 text-lg font-semibold transition-colors duration-300"
        >
          Reserve Ticket
        </button>
      </div>
    </div>
  </div>
</template>
<script setup>

import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { SingleDummy } from '../../components/dummy';
import Heart from "../../assets/icons/Heart.vue";
import { useQuery,useMutation } from '@vue/apollo-composable';
import { GET_EVENT_BY_ID } from "../../components/graphql/queries";
import {USER_MAKE_BOOK_MARK} from "../../components/graphql/queries"
import { useAuthStore } from '../../stores/authstore';
const authStore=useAuthStore()
const route = useRoute();
const router=useRouter()
const id = ref(route.params.id);
const user_id=ref(authStore.userId)
const { result, loading, error, refetch } = useQuery(GET_EVENT_BY_ID, { id: String(id.value) });
const {mutate:bookMark}=useMutation(USER_MAKE_BOOK_MARK)
const count = ref(1);
const heart=ref(false)
const toggleBookmark =  () => {
  if (!localStorage.getItem("token")) {
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    router.push('/auth/login');
    return;
  }
  else{
     try {
    bookMark({

        user_id: String(user_id.value),
        event_id: String(id.value),
        
      
    });

  }catch (error) {
    console.log('Unexpected error:', error);
  }
  }

 
};

// Increase the ticket count by 1
const increaseCount = () => {
  count.value += 1;
};

// Decrease the ticket count by 1
const decreaseCount = () => {
  if (count.value > 1) {
    count.value -= 1;
  }
};
const showHeart=()=>{
  heart.value=true

}
const hideHeart=()=>{
  setTimeout(() => {
     heart.value=false
  }, 500);
 
}
// Reserve Ticket
const reserveTicket = () => {
  console.log(`Reserving ${count.value} tickets.`);
};

// Refetch data on mount
onMounted(() => {
  refetch()

});

// Computed property for handling the data safely
const eventData = computed(() => {
  if (loading.value) return null;
  if (error.value) return { error: true };
  return result.value?.events_by_pk || {};
});
</script>