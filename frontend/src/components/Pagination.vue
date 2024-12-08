<template>
    <div class="flex justify-center items-center mt-4">
      <button
        :disabled="currentPage <= 1"
        @click="$emit('update:page', currentPage - 1)"
        class="px-4 py-2 mr-2 bg-gray-300 rounded disabled:bg-gray-200"
      >
        Previous
      </button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
      <button
        :disabled="currentPage >= totalPages"
        @click="$emit('update:page', currentPage + 1)"
        class="px-4 py-2 ml-2 bg-gray-300 rounded disabled:bg-gray-200"
      >
        Next
      </button>
    </div>
  </template>
  
  <script setup lang="ts">
  import { computed } from 'vue'
  
  const props = defineProps({
    currentPage: { type: Number, required: true },
    totalItems: { type: Number, required: true },
    itemsPerPage: { type: Number, required: true },
  })
  
  const totalPages = computed(() => Math.ceil(props.totalItems / props.itemsPerPage))
  </script>
  