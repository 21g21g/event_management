<template>
   
  <div class="bookmarks-list max-w-xl mx-auto p-5 bg-white shadow-md rounded-lg">
    <!-- <h2 class="text-2xl font-bold text-center mb-5">Your Bookmarked Events</h2>
    <div v-if="loading" class="text-center text-gray-500">Loading...</div>
    <div v-else-if="error" class="text-center text-red-500">
      Error fetching bookmarks: {{ error.message }}
    </div>
    <div v-else>
      <ul class="space-y-5">
        <li
          v-for="bookmark in bookmarks"
          :key="bookmark.id"
          class="p-4 border border-gray-300 rounded-lg"
        >
          <h3 class="text-xl font-semibold">{{ bookmark.event.name }}</h3>
          <p class="text-sm text-gray-600">{{ bookmark.event.date }} - {{ bookmark.event.location }}</p>
          <p class="text-gray-700 mt-2">{{ bookmark.event.description }}</p>
          <img
            :src="bookmark.event.imageUrl"
            alt="Event Image"
            class="mt-3 w-full h-auto rounded-md"
          />
        </li>
      </ul>
    </div> -->
    <h1 class="mt-20">gebeyehu</h1>
  </div>
</template>
<script setup>
// import { ref } from 'vue';
// import { useQuery } from '@vue/apollo-composable';
// import gql from 'graphql-tag';

// const GET_USER_BOOKMARKS = gql`
//   query GetUserBookmarks($userId: uuid!) {
//     bookmarks(where: { user_id: { _eq: $userId } }) {
//       id
//       event {
//         id
//         name
//         date
//         location
//         description
//         imageUrl
//       }
//     }
//   }
// `;

// const userId = ref('your-authenticated-user-id');

// const { result, loading, error } = useQuery(GET_USER_BOOKMARKS, {userId });

// const bookmarks = ref([]);
// if (result.value) {
//   bookmarks.value = result.value.bookmarks;
// }
definePageMeta({
  layout: 'user'
});
</script>