<template>
  <div v-if="visible" :class="alertClass" class="fixed top-20 right-4 px-6 py-4 rounded shadow-lg transition-all z-50">
   <p>{{ message }}</p>
  </div>
</template>
<script setup>
import { computed } from 'vue';
import { defineProps } from 'vue';
const props = defineProps({
  message: {
    type: String,
    required: true,
  },
  type: {
    type: String,
    default: 'success',
  },
  visible: {
    type: Boolean,
    default: false,
  },
});

const alertClass = computed(() => {
  switch (props.type) {
    case 'error':
      return 'bg-red-500 text-white';
    case 'success':
      return 'bg-green-500 text-white';
    default:
      return 'bg-gray-500 text-white';
  }
});
</script>