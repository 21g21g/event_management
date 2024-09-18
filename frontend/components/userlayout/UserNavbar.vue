
<script setup>

import { ref, watch, computed } from 'vue';
import Avater from '../../assets/icons/Avater.vue';
import Toogle from "../../assets/icons/Toogle.vue";
import { useAuthStore } from "../../stores/authstore";
import { useQuery } from "@vue/apollo-composable";
import gql from 'graphql-tag';
import {GET_USER_BY_ID} from "../../utils/queries"
import { useRouter } from 'vue-router';
const router=useRouter()
const authStore = useAuthStore();
const onhover = ref(false);
const isMobileMenuOpen = ref(false);

const hoverClick = () => {
  onhover.value = !onhover.value;
};

const userid = ref(authStore.userId);

const { result: data, loading, error } = useQuery(GET_USER_BY_ID, { id: userid.value });



const userData = computed(() => {
  return data.value?.users_by_pk || null;
});

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
};

const clickLogout=()=>{
  console.log("the user clicked the logout button")
  localStorage.removeItem("token")
  authStore.setUserId(null);
  authStore.setIsloggedin(false);
  router.replace("/auth/login")
  
}
</script>
<style scoped>
.router-link-exact-active{
  color: blueviolet !important;
}

</style>

<template>
  <div class="w-full bg-white shadow-md sticky top-0 z-50">
    <div class="flex justify-between items-center px-4 py-4 md:px-8">
      <NuxtLink to="/user" class="text-2xl font-bold text-blue-600 hover:text-blue-800">
        MinabEvent
      </NuxtLink>
      <button
        @click="toggleMobileMenu"
        class="md:hidden text-blue-600 focus:outline-none"
        aria-label="Toggle navigation menu"
      >
        <Toogle />
      </button>
      <div class="hidden md:flex items-center space-x-6">
        <NuxtLink to="/user/createEvent" exact class="text-blue-600 font-medium px-4 py-2 rounded hover:bg-gray-100 transition">
          Create Event
        </NuxtLink>
        <NuxtLink to="/user/bookMark" exact class="text-blue-600 font-medium px-4 py-2 rounded hover:bg-gray-100 transition">
          Book Marks
        </NuxtLink>

        <div class="relative flex items-center cursor-pointer" @click="hoverClick">
          <Avater />
          <h3 v-if="loading">Loading...</h3>
          <h3 v-else-if="userData" class="text-blue-600 font-medium ml-2">
            {{ userData.username }}
          </h3>
          <h3 v-else class="text-red-500">
            {{ error ? 'Error loading user' : 'No user data found' }}
          </h3>

          <div v-if="onhover" class="absolute top-10 right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border h-28 border-gray-200 transition-transform transform origin-top-right z-10">
            <!-- <NuxtLink to="/user/profile" class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200">
              Profile
            </NuxtLink> -->
            <button @click="clickLogout"
                    class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200">
              Logout
            </button>
            <!-- <NuxtLink to="/help" class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200">
              Help
            </NuxtLink> -->
          </div>
        </div>
      </div>

      <div
        v-if="isMobileMenuOpen"
        class="md:hidden fixed inset-0 bg-opacity-50"
        @click="toggleMobileMenu"
      >
        <div class="absolute top-16 right-2 w-48 flex flex-col bg-white shadow-lg border-t border-gray-200 p-6">
          <NuxtLink to="/user/createEvent" exact class="text-blue-600 font-medium px-4 py-2 rounded hover:bg-gray-100 transition">
            Create Event
          </NuxtLink>
          <NuxtLink to="/user/bookMark" exact class="text-blue-600 font-medium px-4 py-2 rounded hover:bg-gray-100 transition">
            Book Marks
          </NuxtLink>
          <div class="relative flex items-center cursor-pointer" @click="hoverClick">
            <Avater />
            <h3 v-if="loading">Loading...</h3>
            <h3 v-else-if="userData" class="text-blue-600 font-medium ml-2">
              {{ userData.username }}
            </h3>
            <h3 v-else class="text-red-500">
              {{ error ? 'Error loading user' : 'No user data found' }}
            </h3>

            <div
              v-if="onhover"
              class="absolute top-10 right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 transition-transform transform origin-top-right z-10"
            >
              <!-- <NuxtLink
                to="/user/profile"
                exact
                class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200"
              >
                Profile
              </NuxtLink> -->
              <button
                @click="clickLogout"
                exact
                class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200"
              >
                Logout
              </button>
              <!-- <NuxtLink
                to="/help"
                exact
                class="block px-4 py-2 text-gray-700 hover:bg-gray-100 transition-colors duration-200"
              >
                Help
              </NuxtLink> -->
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

