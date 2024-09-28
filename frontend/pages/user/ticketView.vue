<script setup>
definePageMeta({
  layout: 'user',
    middleware:"auth-log"
});

import { GET_TICKET_USER_BY_USER_ID,GET_TRANSACTIONS } from "../../utils/queries";
import { useQuery } from "@vue/apollo-composable";
import { onMounted, ref } from "vue";

// const userId = ref(localStorage.getItem("userId"));
const { result, error, loading,refetch } = useQuery(GET_TRANSACTIONS);

// Helper function to format date
const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  });
};
onMounted(()=>{
  refetch()
})
</script>

<template>
  <div class="mt-20">
    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    
    <div v-else-if="error" class="text-center text-red-500">There is an error fetching your tickets.</div>
    
    <div v-else-if="result?.transactions.length === 0" class="text-center text-red-600">No tickets found.</div>

    <div v-else>
      <div class="overflow-x-auto">
        <table class="min-w-full bg-white border border-gray-200 rounded-lg shadow-md">
          <thead class="bg-gradient-to-r from-gray-100 to-gray-300">
            <tr>
              <th class="px-6 py-3 text-left text-gray-700 font-bold uppercase tracking-wider border-b">checkout_url</th>
              <th class="px-6 py-3 text-left text-gray-700 font-bold uppercase tracking-wider border-b">Amount</th>
              <th class="px-6 py-3 text-left text-gray-700 font-bold uppercase tracking-wider border-b">Event Date</th>
              <th class="px-6 py-3 text-left text-gray-700 font-bold uppercase tracking-wider border-b">PhoneNumber</th>
            </tr>
          </thead>

          <tbody>
            <tr
              v-for="(ticket, index) in result?.transactions"
              :key="ticket.id"
              :class="index % 2 === 0 ? 'bg-white' : 'bg-gray-50'" 
              class="hover:bg-gray-200 transition duration-300 ease-in-out"
            >
<td class="px-6 py-4 border-b text-gray-800">
  <a :href="ticket?.checkout_url" target="_blank" class="text-blue-500 hover:underline">
    {{ ticket?.checkout_url }}
  </a>
</td>              <td class="px-6 py-4 border-b text-gray-800">{{ ticket?.amount }}</td>
              <td class="px-6 py-4 border-b text-gray-800">{{ formatDate(ticket?.event.preparation_date) }}</td>
              <td class="px-6 py-4 border-b text-gray-800">{{ ticket?.phoneNumber }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>