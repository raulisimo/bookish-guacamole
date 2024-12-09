<template>
  <div
    class="min-h-screen bg-cover bg-center bg-fixed"
    style="background-image: url('/galaxy.jpg')"
  >
    <Navbar />
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Navbar from '@/components/StarWarsNavbar.vue'

export default defineComponent({
  name: 'App',
  components: { Navbar },
})
</script>
