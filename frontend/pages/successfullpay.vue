<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();




const goHome = () => {
  router.push('/');
};
</script>

<style scoped>
body {
 margin: 0;
 padding: 0;
}
</style>
<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="bg-white shadow-md rounded-lg p-8 text-center">
      <div class="text-green-500 text-6xl mb-6">✔️</div>
      <h1 class="text-3xl font-bold text-green-500 mb-4">Payment Successful!</h1>
      <p class="text-lg text-gray-700 mb-2">Thank you for your payment.</p>
      
      <p class="text-lg text-gray-700 mb-6">
        You make a transaction Successful.
      </p>
      <button
        class="bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-4 rounded-md transition duration-300"
        @click="goHome"
      >
        Return to Home
      </button>
    </div>
  </div>
</template>

