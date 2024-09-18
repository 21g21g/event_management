<script setup>

   defineProps({
    color:{
        type:String,
        default:'red'
    }
   })
</script>

<template>
     <svg   :style="{ fill: color }" 
 xmlns="http://www.w3.org/2000/svg" class="cursor-pointer" title="unmark" fill="none" viewBox="0 0 24 24" stroke="currentColor" width="60px" height="60px">
  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3l18 18M19.5 13.5a5.978 5.978 0 00-1.44-4.754l-2.21-2.21a6.018 6.018 0 00-8.35 0l-.5.5m-2.26 2.26L12 19.5l2.26-2.26" />
</svg>
</template>
