

import { defineNuxtPlugin } from '#app';
import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from '@apollo/client/core';
import { DefaultApolloClient } from '@vue/apollo-composable';

export default defineNuxtPlugin(nuxtApp => {
  const authLink = new ApolloLink((operation, forward) => {
    if (!process.client) {
      console.warn('This code should only run on the client side.');
      return forward(operation);
    }

    const token = localStorage.getItem('token');
    console.log('Token retrieved:', token);


    if(token){
       operation.setContext({
      
      headers: {
        Authorization: `Bearer ${token}`,
        'x-hasura-role': 'user',
        
      },
    });
    }
   
   

    return forward(operation);
  });

  const httpLink = new HttpLink({
    uri: 'http://localhost:8080/v1/graphql', // Your API URL
  });

  const apolloClient = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(), // Set up cache
  });

  // Provide Apollo client to the Nuxt application
  nuxtApp.provide('apollo', apolloClient);
  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient);
});



