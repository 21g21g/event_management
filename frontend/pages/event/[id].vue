
<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useQuery, useMutation } from '@vue/apollo-composable';
import Heart from '../../assets/icons/Heart.vue';
import { GET_EVENT_BY_ID, USER_MAKE_BOOK_MARK, GET_BOOK_MARKED_EVENT, CATCH_TICKET, GET_TICKET_USER } from '../../utils/queries';
import { useAuthStore } from '../../stores/authstore';
import AlertMessage from '../../components/AlertMessage.vue';
import Map from "../../components/homecomponents/Map.vue";

definePageMeta({
  middleware: "after-log"
});

const authStore = useAuthStore();
const route = useRoute();
const router = useRouter();
const id = ref(route.params.id);
const isBookmarked = ref(false);
const isTicket = ref(false);
const user_id = ref(localStorage.getItem("userId"));
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

const { result: ticketResult, loading: ticketLoading, error: ticketError, refetch: ticketRefetch } = useQuery(GET_TICKET_USER, { user_id: String(user_id.value), event_id: String(id.value) });

const reserveTicket = async () => {
  if (!localStorage.getItem('token')) {
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    showAlert("First you must be logged in!", "info");
    setTimeout(() => {
      router.replace('/auth/login');
    }, 4000);
    return;
  }
  if (isTicket.value === false) {
    try {
      await ticket({
        user_id: String(user_id.value),
        event_id: String(id.value),
        quantity: count.value,
        catchedTicket: true
      });
      showAlert(`You reserved ${count.value} tickets.`, 'success');
      localStorage.setItem('recentlyTicketed', 'true');
      setTimeout(() => {
        router.replace("/user/ticketView");
      }, 4000);
    } catch (error) {
      console.log('Unexpected error:', error);
      showAlert('Something went wrong while catching the ticket.', 'error');
    }
  } else {
    showAlert('You have already caught a ticket for this event!', 'success');
    setTimeout(() => {
      router.replace('/user/ticketView');
    }, 4000);
  }
};

const { result: bookmarkResult, loading: bookmarkLoading, error: bookmarkError, refetch: bookRefetch } = useQuery(GET_BOOK_MARKED_EVENT, { user_id: String(user_id.value), event_id: String(id.value) });
const toggleBookmark = async () => {
  if (!localStorage.getItem('token')) {
    showAlert("First you must be logged in!", "info");
    localStorage.setItem('redirectAfterLogin', window.location.pathname);
    setTimeout(() => {
      router.replace('/auth/login');
    }, 4000);
    return;
  }
  if (isBookmarked.value === false) {
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
  } else {
    showAlert('You have already bookmarked this event!', 'success');
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
    if (bookmarks.length > 0 && bookmarks[0]?.isbookMarked === true) {
      isBookmarked.value = true;
    } else {
      isBookmarked.value = false;
    }
  }
  await ticketRefetch();
  if (!ticketLoading.value && ticketResult.value) {
    const tickets = ticketResult.value?.tickets || [];
    if (tickets.length > 0 && tickets[0]?.catchedTicket === true) {
      isTicket.value = true;
    } else {
      isTicket.value = false;
    }
  }
  refetch()
});
</script>


<template>
  <div class="container mx-auto mt-20 px-6">
    <div class="flex flex-col lg:flex-row gap-10">
      <div class="lg:w-2/3 bg-white rounded-lg p-6">
        <AlertMessage :message="alertMessage" :type="alertType" :visible="alertVisible" />

        <div class="relative">
       <v-carousel v-if="eventData?.imagestores && eventData.imagestores.length > 0" :items-to-show="1" class="rounded-lg">
  <v-carousel-item
    v-for="(item, i) in eventData.imagestores"
    :key="i"
    class="flex items-center justify-center"
  >
    <img
      :src="item.url"
      :alt="item.alt"
      class="object-cover w-full h-72 rounded-lg"
    />
  </v-carousel-item>
</v-carousel>
<div v-else>
  
  <p>Loading images...</p>
</div>

          <Heart
            @click="toggleBookmark"
            :color="isBookmarked ? 'red' : 'orange'"
            class="w-8 h-8 cursor-pointer absolute top-4 right-4 lg:top-6 lg:right-6"
          />
        </div>

        <div class="flex justify-between items-center mt-4">
          <h1 class="text-3xl font-bold">{{ eventData?.title }}</h1>
        </div>

        <p class="text-gray-500 mt-2">{{ eventData?.preparation_date }}</p>

        <h2 class="text-xl font-semibold mt-6">Event Details</h2>
        <p class="text-gray-600 mt-2">{{ eventData?.description }}</p>

        <h2 class="text-xl font-semibold mt-6">Location</h2>
        <p class="text-gray-600">{{ eventData?.address }}</p>
      </div>

      <div class="lg:w-1/3 flex flex-col gap-6">
        <div class="bg-white shadow-lg rounded-lg p-6">
          <h2 class="text-xl font-semibold mb-4">Reserve Your Tickets</h2>

          <div class="flex justify-between items-center mb-4">
            <p class="text-lg">Number of Tickets</p>
            <div class="flex items-center gap-4">
              <button @click="decreaseCount" class="bg-gray-200 p-2 rounded-lg">-</button>
              <span class="text-lg">{{ count }}</span>
              <button @click="increaseCount" class="bg-gray-200 p-2 rounded-lg">+</button>
            </div>
          </div>

          <div v-if="eventData?.price === 'paid'" class="mb-4">
            <p class="text-lg font-bold">Total Price: ${{ totalPrice }}</p>
          </div>

          <button
            @click="reserveTicket"
            :class="isTicket ? 'bg-green-500' : 'bg-blue-500'"
            class="w-full text-white py-3 rounded-lg transition-transform hover:scale-105"
          >
            <span v-if="isTicket">Ticket Reserved</span>
            <span v-else>Reserve Ticket</span>
          </button>
        </div>

        <!-- Map Section -->
        <div class="bg-white shadow-lg rounded-lg p-6">
          <h2 class="text-xl font-semibold mb-4">Event Location</h2>
          <Map class="w-full h-72 rounded-lg shadow-sm"  :address="eventData?.address"/>
        </div>
      </div>
    </div>
  </div>
</template>