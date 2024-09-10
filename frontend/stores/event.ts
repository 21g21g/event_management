import {defineStore} from "pinia"
import { useQuery } from '@vue/apollo-composable'
import gql from "graphql-tag"


const GET_EVENTS = gql`
  query GetEvents {
    events {
      id
      title
      date
    }
  }
`

export const useEventStore=defineStore("event",{
    state: () => ({
        events: [] as Array<{ id: string; title: string; date: string }>, // Adjust the type based on your schema
        loading: false,
        error: null as string | null,
      }),
    
      actions: {
        async fetchEvents() {
          this.loading = true
          this.error = null
    
          try {
            // Execute the GraphQL query using Apollo Client
            const { result } = await useQuery(GET_EVENTS)
    
            // Update the state with the fetched events
            if (result.value) {
              this.events = result.value.events
            } else {
              this.events = []
            }
          } catch (err) {
            // Handle any errors
            this.error = err instanceof Error ? err.message : 'Failed to fetch events'
          } finally {
            this.loading = false
          }
        },
      },

})

// <template>
//   <div>
//     <h1>Events</h1>
//     <div v-if="loading">Loading...</div>
//     <div v-if="error">{{ error }}</div>
//     <ul v-if="!loading && !error">
//       <li v-for="event in events" :key="event.id">{{ event.title }} - {{ event.date }}</li>
//     </ul>
//   </div>
// </template>

// <script setup lang="ts">
// import { useEventsStore } from '~/stores/events'

// // Access the Pinia store
// const eventsStore = useEventsStore()

// // Fetch events on component mount
// eventsStore.fetchEvents()

// // Access the state
// const { events, loading, error } = eventsStore
// </script>