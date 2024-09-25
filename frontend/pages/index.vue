<script setup>
//  definePageMeta({
//   middleware:"after-log"
// })
import eventCatagory from "../components/homecomponents/eventCatagory.vue"
import homecover from "../components/homecomponents/homecover.vue"
import allLocation from "../components/homecomponents/allLocation.vue"

</script>
<template>
   <div class="bg-yellow-50">
       <homecover/>
       <eventCatagory/>  
       <allLocation />   
        
   </div> 
</template>
