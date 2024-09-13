<template>
   
  <div class="bookmarks-list max-w-xl mx-auto p-5 bg-white shadow-md rounded-lg mt-20">
  <h2 class="text-2xl font-bold text-center mb-5">Your Bookmarked Events</h2> 
    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    <div v-else-if="error" class="text-center text-red-500">
      Error fetching bookmarks: {{ error.message }}
    </div>
         <div 
          v-else
          v-for="event in result?.bookmarks" 
          :key="event.id"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out"
        >
          <img :src="event.event.featured_image" alt="Event image" class="w-full h-48 object-cover cursor-pointer"/>
          <div class="p-4">
            <p class="text-lg font-semibold text-gray-800 mb-2">Event Name: {{event.title}}</p>
            <p class="text-sm font-semibold text-gray-600">Location: {{event.address}}</p>
            <p class="text-sm font-semibold text-gray-600">Preparation Date: For {{event.preparation_date}}</p>
          </div>
           
        </div>
  </div>
</template>
<script setup>
definePageMeta({
  layout: 'user'
});
import { useQuery } from "@vue/apollo-composable";
import { onMounted, ref } from "vue";
import {GET_BOOK_MARK_BY_USER_ID} from "../../components/graphql/queries"
import { useAuthStore } from "../../stores/authstore";
const authStore=useAuthStore()
const userId=ref(authStore.userId)
const {result,loading,error,refetch}=useQuery(GET_BOOK_MARK_BY_USER_ID,{user_id:String(userId.value)})
onMounted(()=>{
  refetch()
})

</script>