<template>
  <footer class="bg-gray-900 text-white py-8">
    <div class="container mx-auto px-4">
      <!-- Footer Top: Main Links and Contact -->
      <div class="flex flex-col md:flex-row justify-between items-center">
        <!-- Branding and Description -->
        <div class="mb-4 md:mb-0">
          <h2 class="text-2xl font-bold">MinabEvent</h2>
          <p class="text-gray-400 mt-2">
            Discover and book the best events happening near you. Join our community and never miss out on the latest events.
          </p>
        </div>

        <!-- Navigation Links -->
        <div class="flex space-x-8">
          <a href="/" class="text-gray-400 hover:text-white">Home</a>
          <a href="/events" class="text-gray-400 hover:text-white">Events</a>
          <a href="/about" class="text-gray-400 hover:text-white">About</a>
          <a href="/contact" class="text-gray-400 hover:text-white">Contact</a>
        </div>
      </div>

      <!-- Footer Middle: Social Media Links -->
      <div class="flex justify-center space-x-4 mt-8">
        <a href="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-facebook" class="w-6 h-6" />
        </a>
        <a href="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-twitter" class="w-6 h-6" />
        </a>
        <a href="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-instagram" class="w-6 h-6" />
        </a>
        <a href="#" class="text-gray-400 hover:text-white">
          <nuxt-icon name="mdi-linkedin" class="w-6 h-6" />
        </a>
      </div>

      <!-- Footer Bottom: Copyright and Legal Links -->
      <div class="flex flex-col md:flex-row justify-between items-center mt-8">
        <p class="text-gray-400">&copy; 2024 MinabEvent. All rights reserved.</p>
        <div class="flex space-x-4">
          <a href="/terms" class="text-gray-400 hover:text-white">Terms of Service</a>
          <a href="/privacy" class="text-gray-400 hover:text-white">Privacy Policy</a>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { useRouter } from "vue-router";

// Importing nuxt-icon to use for social media icons
// import { defineNuxtPlugin } from '#app';
// import 'nuxt-icon/css/index.css';
</script>