<script setup>
definePageMeta({
  layout: 'user',
    middleware:"auth-log"
});
import { useMutation, useQuery } from "@vue/apollo-composable";
import { computed, onMounted, ref } from "vue";
import {GET_BOOK_MARK_BY_USER_ID} from "../../utils/queries"
import {REMOVE_BOOKMARK_FROM_FAVOURITE} from "../../utils/queries"
import AlertMessage from "../../components/AlertMessage.vue"
import Unfav from "../../assets/icons/Unfav.vue"
const hoverStates = ref({});
const alertMessage = ref('');
const alertVisible = ref(false);
const alertType = ref('success');
const setHovered = (eventId, state) => {
hoverStates.value[eventId] = state;
     };
const userId=ref(localStorage.getItem("userId"))
const {result,loading,error,refetch}=useQuery(GET_BOOK_MARK_BY_USER_ID)
const {mutate:deleteBookmark}=useMutation(REMOVE_BOOKMARK_FROM_FAVOURITE)
const showAlert = (message, type = 'success') => {
  alertMessage.value = message;
  alertType.value = type;
  alertVisible.value = true;
  setTimeout(() => {
    alertVisible.value = false;
  }, 3000); 
};
const EventData = computed(() => {
  return result.value?.bookmarks || []; 
});
const unMarked=async(id)=>{
  try{

    const {data}=await deleteBookmark({id:String(id)})
    if(data.delete_bookmarks_by_pk.id===id){
      showAlert("Succesfully unmarked!","success")
      setTimeout(()=>{
        refetch()
      },4000)
       }
 }catch(error){
    showAlert("something went wrong!","error")
  }

}

onMounted(()=>{
  refetch()
})
</script>
<template>
  <div class="bookmarks-list max-w-4xl mx-auto p-5 bg-white rounded-lg shadow-md mt-20 gap-4">
    <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    
    <div v-else-if="error" class="text-center text-red-500">There is an error fetching your bookmarks.</div>
    
    <div v-else-if="EventData.length === 0" class="text-center text-red-600">No bookmarks found.</div>
    
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="event in EventData" 
        :key="event.id"
        class="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 ease-in-out"
      >
        <div class="relative">
          <img :src="event?.event.featured_image" alt="Event image" class="w-full h-48 object-cover"/>
          <Unfav class="absolute top-3 right-2 cursor-pointer" 
                 @click="unMarked(event.id)" 
                 @mouseover="setHovered(event.id, true)" 
                 @mouseleave="setHovered(event.id, false)"/>
          <span 
            v-if="hoverStates[event.id]" 
            class="absolute top-16 right-2 bg-gray-800 text-white text-sm px-2 py-1 rounded-md shadow-lg"
          >
            Unmark
          </span>
        </div>
        
        <div class="p-4">
          <p class="text-lg font-semibold text-gray-800 mb-2">Event Name: {{event?.event.title}}</p>
          <p class="text-sm font-semibold text-gray-600">Location: {{event?.event.address}}</p>
          <p class="text-sm font-semibold text-gray-600">Preparation Date: For {{event?.event.preparation_date}}</p>
        </div>
      </div>
    </div>
  </div>
</template>
