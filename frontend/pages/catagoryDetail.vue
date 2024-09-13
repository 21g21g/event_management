<template>
<div class="mt-24 p-4">
<coverpageBanner :image="image" :category="category"/>
<catagoryDetailHome :category="category"/>
</div>
</template>
<script setup>
import { computed } from "vue"
import { useRoute } from 'vue-router';
const route=useRoute()
import catagoryDetailHome from "../components/catagory/catagoryDetailHome.vue"
import coverpageBanner from "../components/catagory/coverpageBanner.vue"
const category=computed(()=>route.query.category)
const image=computed(()=>route.query.image)
</script>
