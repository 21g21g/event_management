import { defineNuxtPlugin } from '#app';
import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core';
import { DefaultApolloClient } from '@vue/apollo-composable';

export default defineNuxtPlugin((nuxtApp) => {
  // Direct HTTP connection to the API with headers
  const token = process.client ? localStorage.getItem('token') : null; // Retrieve token only on the client side
  const role = 'user'; // Set a static role

  const httpLink = new HttpLink({
    uri: 'http://localhost:8080/v1/graphql', // API URL
    headers: {
      'x-hasura-admin-secret':'assegagebeyehu212121'
    //   authorization: token ? `Bearer ${token}` : '', // Set authorization header if token is available
    //   'X-Hasura-Role': role, // Set the Hasura role header
    },
  });

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(), // Set up cache
  });

  // Provide Apollo client to the Nuxt application
  nuxtApp.provide('apollo', apolloClient);
  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient);
});