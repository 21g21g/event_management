<script setup>
definePageMeta({
  layout: 'user',
    middleware:"auth-log"

  
});
import { useQuery } from "@vue/apollo-composable";
import { onMounted, ref } from "vue";
import {GET_BOOK_MARK_BY_USER_ID} from "../../utils/queries"
import { useAuthStore } from "../../stores/authstore";
const authStore=useAuthStore()
const userId=ref(authStore.userId)
const {result,loading,error,refetch}=useQuery(GET_BOOK_MARK_BY_USER_ID,{user_id:String(userId.value)})
onMounted(()=>{
  refetch()
})

</script>
<template>
   
  <div class="bookmarks-list max-w-xl mx-auto p-5 bg-white  rounded-lg mt-20">
    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    
    <div v-else-if="error" class="text-center text-red-500">There is an error fetching your bookmarks.</div>
    
    <div v-else-if="result?.bookmarks.length === 0" class="text-center text-red-600">No bookmarks found.</div>
      <div  v-else>

         <div 
         
          v-for="event in result?.bookmarks" 
          :key="event.id"
          class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out"
        >
          <img :src="event.event.featured_image" alt="Event image" class="w-full h-48 object-cover cursor-pointer"/>
          <div class="p-4">
            <p class="text-lg font-semibold text-gray-800 mb-2">Event Name: {{event.event.title}}</p>
            <p class="text-sm font-semibold text-gray-600">Location: {{event.event.address}}</p>
            <p class="text-sm font-semibold text-gray-600">Preparation Date: For {{event.event.preparation_date}}</p>
          </div>
           
        </div>
    </div>
  
  </div>
</template>
