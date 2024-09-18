<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useQuery, useMutation } from '@vue/apollo-composable';
import Heart from '../../assets/icons/Heart.vue';
import { GET_EVENT_BY_ID, USER_MAKE_BOOK_MARK, GET_BOOK_MARKED_EVENT, CATCH_TICKET,GET_TICKET_USER } from '../../utils/queries';
import { useAuthStore } from '../../stores/authstore';
import  AlertMessage  from '../../components/AlertMessage.vue';
definePageMeta({
  middleware:"after-log"
})
const authStore = useAuthStore();
const route = useRoute();
const router = useRouter();
const id = ref(route.params.id);
const isBookmarked = ref(false);
const isTicket=ref(false)
const user_id = ref(authStore.userId);
const alertMessage = ref('');
const alertVisible = ref(false);
const alertType = ref('success');
const { result, loading, error, refetch } = useQuery(GET_EVENT_BY_ID, {
  id: String(id.value),
});

const { mutate: bookMark } = useMutation(USER_MAKE_BOOK_MARK);
const { mutate: ticket } = useMutation(CATCH_TICKET);

const count = ref(1);
const increaseCount = () => count.value += 1;
const decreaseCount = () => {
  if (count.value > 1) {
    count.value -= 1;
  }
};

const totalPrice = computed(() => {
  if (eventData.value?.price === 'paid') {
    return (parseFloat(eventData.value.specific_price) * count.value).toFixed(2);
  }
  return eventData.value?.price || '';
});
const showAlert = (message, type = 'success') => {
  alertMessage.value = message;
  alertType.value = type;
  alertVisible.value = true;
  setTimeout(() => {
    alertVisible.value = false;
  }, 3000); 
};

const {result:ticketResult,loading:ticketLoading,error:ticketError,refetch:ticketRefetch}=useQuery(GET_TICKET_USER,{user_id:user_id.value,event_id:id.value,})

const reserveTicket = async () => {
  if (!localStorage.getItem('token')) {
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    router.replace('/auth/login');
    return;
  }
  if(isTicket.value===false){
    try{
  await ticket({
    user_id: user_id.value,
    event_id: id.value,
    quantity: count.value,
    catchedTicket:true
  });

showAlert(`You reserved ${count.value} tickets.`, 'success');  
 localStorage.setItem('recentlyTicketed', 'true');
setTimeout(()=>{
  router.replace("/user/ticketView");

},4000
)

}catch(error){
  console.log('Unexpected error:', error);
    showAlert('Something went wrong while catching the ticket.', 'error');

}
  }
   else{
      showAlert('You have already catch ticket of an event!', 'success');

    setTimeout(() => {
      router.replace('/user/ticketView');
    }, 2000); 
  }

  
};

const { result: bookmarkResult, loading: bookmarkLoading, error: bookmarkError, refetch: bookRefetch } = useQuery(GET_BOOK_MARKED_EVENT, {
  user_id: String(user_id.value),
  event_id: String(id.value),
});
const toggleBookmark = async () => {
  if (!localStorage.getItem('token')) {
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    router.replace('/auth/login');
    return;
  }
    if(isBookmarked.value===false){
        try {
    await bookMark({
      user_id: String(user_id.value),
      event_id: String(id.value),
      isbookMarked: true
    });
    showAlert('Event bookmarked successfully!', 'success');
    localStorage.setItem('recentlyBookmarked', 'true');
    
    setTimeout(() => {
      router.replace('/user/bookMark');
    }, 2000); 
  } catch (error) {
    console.log('Unexpected error:', error);
    showAlert('Something went wrong while bookmarking the event.', 'error');
  }

    }
    else{
          showAlert('You have already bookmarked an event!', 'success');  
      setTimeout(() => {
        router.replace('/user/bookMark');
      }, 2000); 
    }

};
const eventData = computed(() => {
  if (loading.value) return null;
  if (error.value) return { error: true };
  return result.value?.events_by_pk || {};
});
    
     
onMounted(async () => {
  await bookRefetch();

  if (!bookmarkLoading.value && bookmarkResult?.value) {
    const bookmarks = bookmarkResult?.value?.bookmarks || [];
    console.log(bookmarks[0]?.isbookMarked)
       if (bookmarks.length > 0 && bookmarks[0]?.isbookMarked === true) {
         isBookmarked.value = true;
      
    } else {
      isBookmarked.value = false;
    }
    }
  
  else if (bookmarkLoading.value) {
    console.log('Loading bookmark data...');
  }

  if (bookmarkError.value) {
    console.log('Error checking bookmark status:', bookmarkError.value);
  }


  await ticketRefetch()
  if(!ticketLoading.value && ticketResult.value){
    const tickets=ticketResult.value?.tickets||[]
  console.log(tickets[0]?.catchedTicket)  
    if(tickets.length >0 && tickets[0]?.catchedTicket===true){
      isTicket.value=true
      
    }
    else{
      isTicket.value=false
    }
  }else if(ticketResult.value){
    console.log("loading tickets")
  }
   if (ticketError.value) {
    console.log('Error checking bookmark status:', ticketError.value);
  }

});
</script>

<template>
  <div class="flex flex-col mt-20 bg-gradient-to-b from-gray-100 to-gray-300 py-10 px-4">
    <!-- Alert Message -->
    <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

    <!-- Carousel Section -->
    <div class="w-full max-w-6xl mx-auto mt-10 relative shadow-2xl bg-white rounded-lg overflow-hidden">
      <v-carousel :items-to-show="1" :items-to-scroll="1" cycle autoplay class="h-auto rounded-lg shadow-lg">
        <v-carousel-item
          v-for="(item, i) in eventData?.imagestores"
          :key="i"
          class="flex items-center justify-center w-full hover:scale-105 transition-transform duration-300"
        >
          <img
            :src="item.url"
            :alt="item.alt"
            class="object-cover w-full h-80 md:h-96 lg:h-120 rounded-md"
          />
        </v-carousel-item>
      </v-carousel>
      <!-- Bookmark Button -->
      <Heart
        @click="toggleBookmark"
        :color="isBookmarked ? 'red' : 'orange'"
        class="w-12 h-12 md:w-16 md:h-16 lg:w-20 lg:h-24 absolute top-5 right-5 cursor-pointer transition-transform duration-300 transform hover:scale-110"
      />
    </div>

    <!-- Event Details Section -->
    <div class="flex flex-col md:flex-row justify-between w-full max-w-6xl mx-auto p-6 mt-8 bg-white rounded-lg shadow-lg space-y-6 md:space-y-0">
      <!-- Left Content -->
      <div class="flex flex-col w-full md:w-2/3 space-y-6">
        <h1 class="text-2xl md:text-3xl lg:text-4xl font-bold text-gray-900">
          Start Date: {{ eventData?.preparation_date }}
        </h1>
        <p class="text-lg md:text-xl font-semibold text-gray-800">
          Title: {{ eventData?.title }}
        </p>
        <h2 class="text-md md:text-lg font-medium text-gray-700">
          Location: {{ eventData?.address }}
        </h2>
        <h3 class="text-md md:text-lg font-medium text-gray-700">About the Event</h3>
        <p class="text-gray-600 text-sm md:text-base">{{ eventData?.description }}</p>
      </div>

      <!-- Right Content - Ticket Info -->
      <div class="flex flex-col w-full md:w-1/3 p-6 space-y-6 rounded-md bg-gray-50 border-2 border-gray-200">
        <div class="flex flex-col space-y-4">
          <!-- Ticket Count Control -->
          <div class="flex items-center justify-between">
            <p class="text-md md:text-lg font-semibold text-gray-800">Number of Tickets</p>
            <div class="flex items-center space-x-2">
              <button
                @click="decreaseCount"
                :disabled="count <= 1"
                class="bg-gray-300 hover:bg-gray-400 rounded-md p-2 shadow-sm transition duration-200"
              >
                <span class="text-2xl font-bold">-</span>
              </button>
              <p class="text-lg md:text-xl font-semibold">{{ count }}</p>
              <button
                @click="increaseCount"
                class="bg-gray-300 hover:bg-gray-400 rounded-md p-2 shadow-sm transition duration-200"
              >
                <span class="text-2xl font-bold">+</span>
              </button>
            </div>
          </div>

          <!-- Price Information -->
          <p v-if="eventData?.price === 'paid'" class="text-md md:text-lg font-bold text-gray-800">
            Total Price: ${{ totalPrice }}
          </p>
          <p v-else class="text-md md:text-lg font-semibold text-gray-800">{{ eventData?.price }}</p>
        </div>

        <!-- Reserve Ticket Button -->
        <button
          @click="reserveTicket"
          :class="{
            'bg-green-500': isTicket,
            'bg-slate-400': !isTicket
          }"
          class="bg-gradient-to-r from-blue-500 to-green-500 text-white hover:from-blue-600 hover:to-green-600 rounded-md py-3 px-6 text-lg font-semibold shadow-md transition-transform duration-300 transform hover:scale-105"
        >
          <span v-if="isTicket">Ticket Reserved</span>
          <span v-if="!isTicket">Reserve Ticket</span>
        </button>
      </div>
    </div>
  </div>
</template>

