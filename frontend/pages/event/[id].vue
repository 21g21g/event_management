<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useQuery, useMutation } from '@vue/apollo-composable';
import Heart from '../../assets/icons/Heart.vue';
import { GET_EVENT_BY_ID, USER_MAKE_BOOK_MARK, GET_BOOK_MARKED_EVENT, CATCH_TICKET,GET_TICKET_USER } from '../../utils/queries';
import { useAuthStore } from '../../stores/authstore';

const authStore = useAuthStore();
const route = useRoute();
const router = useRouter();
const id = ref(route.params.id);
const isBookmarked = ref(false);
const isTicket=ref(false)
const user_id = ref(authStore.userId);

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

const {result:ticketResult,loading:ticketLoading,error:ticketError,refetch:ticketRefetch}=useQuery(GET_TICKET_USER,{user_id:user_id.value,event_id:id.value,})

const reserveTicket = async () => {
  if (!localStorage.getItem('token')) {
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    router.replace('/auth/login');
    return;
  }
try{
  await ticket({
    user_id: user_id.value,
    event_id: id.value,
    quantity: count.value,
    catchedTicket:true
  });
  isTicket.value=true

  alert(`You reserved ${count.value} tickets.`);
   localStorage.setItem('recentlyTicketed', 'true');

  router.replace("/user/ticketView");

}catch(error){
  console.log('Unexpected error:', error);
    alert('Something went wrong while catch the ticket.');

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

  try {
    await bookMark({
      user_id: String(user_id.value),
      event_id: String(id.value),
      isbookMarked: true
    });
    isBookmarked.value = true;
    alert('Event bookmarked successfully!');
    
    localStorage.setItem('recentlyBookmarked', 'true');
    
    router.replace("/user/bookMark"); 
  } catch (error) {
    console.log('Unexpected error:', error);
    alert('Something went wrong while bookmarking the event.');
  }
};
const eventData = computed(() => {
  if (loading.value) return null;
  if (error.value) return { error: true };
  return result.value?.events_by_pk || {};
});

onMounted(async () => {
  await bookRefetch();

  if (!bookmarkLoading.value && bookmarkResult.value) {
    const bookmarks = bookmarkResult.value?.bookmarks || [];
    isBookmarked.value = bookmarks.length > 0;
  } else if (bookmarkLoading.value) {
    console.log('Loading bookmark data...');
  }

  if (bookmarkError.value) {
    console.log('Error checking bookmark status:', bookmarkError.value);
  }

  if (localStorage.getItem('recentlyBookmarked') === 'true') {
    localStorage.removeItem('recentlyBookmarked');

    alert('You have already bookmarked an event!');
    setTimeout(() => {
      router.replace('/user/bookMark');
    }, 1000); 
  }
  await ticketRefetch()
  if(!ticketLoading.value && ticketResult.value){
    const tickets=ticketResult.value?.tickets||[]
    isTicket.value=tickets.length > 0
  }else if(ticketResult.value){
    console.log("loading tickets")
  }
   if (ticketError.value) {
    console.log('Error checking bookmark status:', ticketError.value);
  }

  if (localStorage.getItem('recentlyTicketed') === 'true') {
    localStorage.removeItem('recentlyTicketed');

    alert('You have already catch ticket of an event!');
    setTimeout(() => {
      router.replace('/user/ticketView');
    }, 1000); 
  }
});
</script>

<template>
  <div class="flex flex-col mt-20 bg-gradient-to-b from-gray-100 to-gray-300 py-10 px-4">
    <div class="w-full max-w-4xl mx-auto mt-10 relative shadow-xl bg-white rounded-lg overflow-hidden">
      <v-carousel :items-to-show="2" :items-to-scroll="1" cycle class="h-auto rounded-lg shadow-lg">
        <v-carousel-item
          v-for="(item, i) in eventData?.imagestores"
          :key="i"
          class="flex items-center justify-center w-full hover:scale-105 transition-transform duration-300"
        >
          <img
            :src="item.url"
            :alt="item.alt"
            class="object-cover w-full h-72 rounded-md "
          />
        </v-carousel-item>
      </v-carousel>
    </div>

    <div class="flex justify-center items-center mt-4">
      <button
        @click="toggleBookmark"
        :disabled="isBookmarked"
        class="cursor-pointer transition-transform duration-300 transform hover:scale-125"
        :class="{
          'bg-gray-950 text-white': isBookmarked,
          'bg-slate-400 text-gray-500': !isBookmarked,
        }"
      >
        <Heart />
        <span v-if="!isBookmarked">Bookmark Event</span>
        <span v-if="isBookmarked">Event Bookmarked</span>
      </button>
    </div>

    <div class="flex flex-col md:flex-row justify-between w-full p-4 mt-5 bg-white rounded-lg">
      <h1 v-if="loading" class="text-2xl font-bold">Loading...</h1>
      <h1 v-else-if="error" class="text-2xl font-bold text-red-500">There is an Error</h1>
      <div v-else class="flex flex-col w-full md:w-2/3 space-y-6">
        <h1 class="text-3xl font-bold text-gray-900">
          Start Date: {{ eventData?.preparation_date }}
        </h1>
        <p class="text-xl font-semibold text-gray-800">
          Title: {{ eventData?.title }}
        </p>
        <h2 class="text-lg font-medium text-gray-700">
          Location: {{ eventData?.address }}
        </h2>
        <h3 class="text-lg font-medium text-gray-700">About the Event</h3>
        <p class="text-gray-600">{{ eventData?.description }}</p>
      </div>

      <div class="border-2 border-gray-200 flex flex-col w-full md:w-1/3 p-4 space-y-4 rounded-md bg-gray-50">
        <div class="flex flex-col space-y-2">
          <div class="flex items-center justify-between">
            <p class="text-lg font-semibold text-gray-800">Number of Tickets</p>
            <div class="flex items-center space-x-2">
              <button
                @click="decreaseCount"
                :disabled="count <= 1"
                class="bg-gray-300 hover:bg-gray-400 rounded-md p-2 shadow-sm"
              >
                <span class="text-2xl font-bold">-</span>
              </button>
              <p class="text-xl font-semibold">{{ count }}</p>
              <button
                @click="increaseCount"
                class="bg-gray-300 hover:bg-gray-400 rounded-md p-2 shadow-sm"
              >
                <span class="text-2xl font-bold">+</span>
              </button>
            </div>
          </div>

          <p v-if="eventData?.price === 'paid'" class="text-xl font-bold text-gray-800">
            Total Price: ${{ totalPrice }}
          </p>
          <p v-else class="text-xl font-semibold text-gray-800">{{ eventData?.price }}</p>
        </div>

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