<template>
  <div class="flex flex-col mt-20">
    <div class="w-full max-w-3xl mx-auto mt-10">
      <v-carousel
        :items-to-show="3"
        :items-to-scroll="1"
        cycle
        class="h-auto"
      >
        <v-carousel-item
          v-for="(item, i) in SingleDummy.images"
          :key="i"
          class="flex items-center justify-center w-full hover:scale-105 transition-transform duration-300"
        >
          <img
            :src="item"
            :alt="item.alt"
            class="object-cover w-full h-64 rounded-md shadow-lg"
          />
        </v-carousel-item>
      </v-carousel>
     
          <div
        class="w-12 h-12 mt-4 cursor-pointer text-gray-400 hover:text-red-500 transition-all duration-300"
        @click="toggleBookmark"
      >
        <Heart />
      </div>
      
      
    </div>

    <div class="flex flex-col md:flex-row justify-between w-full p-4 mt-5">
      <div class="flex flex-col w-full md:w-2/3 space-y-4">
        <h1 class="text-2xl font-bold text-gray-800">Friday, September 6</h1>
        <p class="text-xl font-semibold text-gray-700">{{ SingleDummy.eventName }}</p>
        <h2 class="text-xl font-semibold text-gray-600">Location: {{ SingleDummy.location }}</h2>
        <h3 class="text-lg font-medium text-gray-600">About the Event</h3>
        <p class="text-gray-600">{{ SingleDummy.description }}</p>
      </div>

      <div class="border-2 border-gray-300 flex flex-col w-full md:w-1/3 p-4 space-y-4 rounded-md">
        <div class="flex flex-col space-y-2">
          <div class="flex items-center justify-between">
            <p class="text-lg font-semibold">General Admission</p>
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
          <p class="text-xl font-bold">{{ SingleDummy.price }}</p>
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
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { SingleDummy } from '../../components/dummy';
import Heart from "../../assets/icons/Heart.vue"
const route = useRoute();
const id = route.params.id;
const count = ref(1);
const toggleBookmark = () => {
  console.log(`Bookmark state: ${isBookmarked.value}`);
};
// Increase the ticket count by 1
const increaseCount = () => {
  count.value += 1;
};

const decreaseCount = () => {
  if (count.value > 1) {
    count.value -= 1;
  }
};

const reserveTicket = () => {
    // Logic to handle ticket reservation
    console.log(`Reserving ${count.value} tickets.`);
  
};

// Use this id to make a request from the backend by id to fetch specific data
</script>