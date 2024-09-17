<template>
  <div class="relative p-4 mt-24">
    <img :src="currentImage" alt="Cover Image" class="w-full h-96 object-cover rounded-lg shadow-lg"/>

    <div class="absolute inset-0 flex flex-col justify-center items-center text-center text-white  bg-opacity-50 p-8 rounded-lg">
      <h1 class="text-4xl font-bold mb-4">Welcome to Minabevent</h1>
      <p class="text-lg mb-6">Discover and manage your events with ease. Find the best events happening near you and never miss out!</p>
   
    </div>
  </div>
</template>

<script setup>

import cover from "../../assets/css/for.avif"
import event from "../../assets/css/event.jpeg"
import ev1 from "../../assets/css/eve1.webp"
import ev2 from "../../assets/css/eve3.avif"
import ev3 from "../../assets/css/eve3.avif"
import { ref,onMounted } from "vue"
const events=[event,ev1,ev2,ev3,cover]
const currentIndex = ref(0);
const currentImage = ref(events[currentIndex.value]);
const changeImage = () => {
  currentIndex.value = (currentIndex.value + 1) % events.length;
  currentImage.value = events[currentIndex.value];
};

onMounted(() => {
  setInterval(changeImage, 1000); 
});
</script>