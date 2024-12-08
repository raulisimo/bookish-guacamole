<template>
  <div class="p-4">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-2xl font-bold">Data Viewer</h2>
      <button
        @click="toggleType"
        class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
      >
        Switch to {{ currentType === 'planets' ? 'People' : 'Planets' }}
      </button>
    </div>
    <DataTable :data="dataStore.data" :onSortChange="onSortChange" />
    <p v-if="dataStore.error" class="text-red-500 mt-2">{{ dataStore.error }}</p>
    <p v-if="dataStore.isLoading" class="text-gray-500 mt-2">Loading...</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useDataStore } from '@/stores/starWarsStore';
import DataTable from '@/components/DataTable.vue';

export default defineComponent({
  components: { DataTable },
  setup() {
    const dataStore = useDataStore();
    const currentType = ref<'people' | 'planets'>('planets');
    const sortParams = ref({ field: '', order: 'asc' });

    const fetchData = () => {
      dataStore.fetchData(currentType.value, sortParams.value.field, sortParams.value.order);
    };

    const toggleType = () => {
      currentType.value = currentType.value === 'planets' ? 'people' : 'planets';
      fetchData();
    };

    const onSortChange = (sort: { field: string; order: string }) => {
      sortParams.value = sort;
      fetchData();
    };

    fetchData();

    return { dataStore, currentType, toggleType, onSortChange };
  },
});
</script>
