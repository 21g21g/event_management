<template>
<div class="mt-24 ml-9">
    <h1 class="text-black">{{counters.count}}</h1> 
   <button @click="counters.increment">increase</button>
   </div>
   </template>
<script setup>
import { ref } from 'vue';
import { usecounterStore } from '../stores/counter';
const counters=usecounterStore()
</script>