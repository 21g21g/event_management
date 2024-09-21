

import { defineNuxtPlugin } from '#app';
import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from '@apollo/client/core';
import { DefaultApolloClient } from '@vue/apollo-composable';

export default defineNuxtPlugin(nuxtApp => {
  const authLink = new ApolloLink((operation, forward) => {
    if (!process.client) {
      console.warn('This code should only run on the client side.');
      return forward(operation);
    }

    // Retrieve token from localStorage
    const token = localStorage.getItem('token');
    const userId = localStorage.getItem('userId');
    console.log('Token retrieved:', token);
    console.log('User ID retrieved:', userId);


    // Set the headers for each request
    if(token){
       operation.setContext({
      
      headers: {
        
        Authorization: token ? `Bearer ${token}` : '',
        'x-hasura-admin-secret': nuxtApp.$config.public.hasuraAdminSecret,
        'x-hasura-role': 'user',
         'x-hasura-user-id': userId
        
      },
    });
    }
    else{
      operation.setContext({
      
        headers: {
          
          Authorization: token ? `Bearer ${token}` : '',
          'x-hasura-admin-secret': nuxtApp.$config.public.hasuraAdminSecret,
         
          
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



