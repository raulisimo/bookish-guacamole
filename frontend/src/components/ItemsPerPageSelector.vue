<template>
    <div class="mb-4">
      <label for="itemsPerPage" class="mr-2">Items per page:</label>
      <select id="itemsPerPage" v-model="selected" @change="emitChange">
        <option v-for="option in options" :key="option" :value="option">{{ option }}</option>
      </select>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, defineProps, defineEmits } from 'vue'
  
  const props = defineProps({
    options: {
      type: Array,
      default: () => [10, 15, 20, 30],
    },
  })
  
  const emit = defineEmits(['update:limit'])
  
  const selected = ref(10)
  
  function emitChange() {
    emit('update:limit', selected.value) // Correctly reference `emit`
  }
  </script>
  