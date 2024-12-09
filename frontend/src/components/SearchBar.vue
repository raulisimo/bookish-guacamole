<template>
  <div class="space-search-container mb-4">
    <input
      v-model="searchQuery"
      type="text"
      placeholder="Search for planets or people..."
      class="search-input"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue'

export default defineComponent({
  name: 'SearchBar',
  emits: ['update:searchQuery'],
  setup(_, { emit }) {
    const searchQuery = ref('')

    // Emit searchQuery change to the parent component
    watch(searchQuery, (newQuery) => {
      emit('update:searchQuery', newQuery)
    })

    return { searchQuery }
  },
})
</script>

<style scoped>
/* Container for centering the search bar */
.space-search-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 1.5rem; /* Add some space at the bottom */
  position: relative;
  padding: 20px;
}

/* Style for the input field */
.search-input {
  width: 100%;
  max-width: 600px; /* Limit the width to prevent it from becoming too wide */
  padding: 15px 25px; /* Larger padding for a bigger input field */
  font-size: 1.25rem; /* Bigger font size */
  border-radius: 50px; /* Rounded corners for a sleek look */
  background-color: #1e1e1e; /* Dark background */
  color: #fff; /* White text color */
  border: 2px solid #444; /* Subtle border */
  outline: none; /* Remove default outline */
  transition: all 0.3s ease-in-out; /* Smooth transition for interactions */
  box-shadow: 0 0 10px rgba(0, 255, 255, 1); /* Futuristic glow effect */
}

/* Focus state for the input */
.search-input:focus {
  border-color: #00b3b3; /* Change border color to cyan */
  box-shadow: 0 0 15px rgba(0, 255, 255, 1); /* Intense cyan glow */
}

/* Placeholder styling */
.search-input::placeholder {
  color: #bbb; /* Light gray color for placeholder */
  font-style: italic;
}

/* Add a little animation on focus */
.search-input:focus::placeholder {
  opacity: 0.6; /* Fade out placeholder when focused */
  transition: opacity 0.2s ease-in-out;
}

/* Hover effect */
.search-input:hover {
  border-color: #ffcc00; /* Change border to gold when hovering */
  box-shadow: 0 0 15px rgba(255, 204, 0, 1); /* Gold glow effect */
}
</style>
