import { defineNuxtPlugin } from '#app';
import { ApolloClient, InMemoryCache } from '@apollo/client/core';
import { DefaultApolloClient } from '@vue/apollo-composable';

export default defineNuxtPlugin((nuxtApp) => {
  const apolloClient = new ApolloClient({
    uri: 'http://localhost:8080/v1/graphql', 
    cache: new InMemoryCache(),
    headers:{
      'x-hasura-admin-secret': 'assegagebeyehu212121', 
    }
  });

  nuxtApp.provide('apollo', apolloClient);
  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient);
});

