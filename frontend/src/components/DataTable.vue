<template>
    <div>
      <input
        v-model="searchQuery"
        type="text"
        placeholder="Search..."
        class="border rounded px-2 py-1 mb-4 w-full"
      />
      <table class="table-auto w-full border-collapse border border-gray-300">
        <thead>
          <tr>
            <th
              v-for="(key, index) in headers"
              :key="index"
              class="border px-4 py-2 cursor-pointer"
              @click="toggleSort(key)"
            >
              {{ key }}
              <span v-if="isSortable(key)">
                <span v-if="currentSort.field === key">
                  {{ currentSort.order === 'asc' ? '⬆️' : '⬇️' }}
                </span>
                <span v-else>⬍</span>
              </span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in filteredData" :key="item.id" class="hover:bg-gray-100">
            <td v-for="(value, index) in Object.values(item)" :key="index" class="border px-4 py-2">
              {{ value }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent, computed, ref, watch } from 'vue';
  
  export default defineComponent({
    name: 'DataTable',
    props: {
      data: {
        type: Array,
        required: true,
      },
      onSortChange: {
        type: Function,
        required: true,
      },
    },
    setup(props) {
      const searchQuery = ref('');
      const currentSort = ref({ field: '', order: 'asc' });
  
      const headers = computed(() => (props.data[0] ? Object.keys(props.data[0]) : []));
  
      const filteredData = computed(() => {
        if (!searchQuery.value) return props.data;
        return props.data.filter((item) =>
          Object.values(item).some((value) =>
            String(value).toLowerCase().includes(searchQuery.value.toLowerCase())
          )
        );
      });
  
      const isSortable = (key: string) => ['name', 'created'].includes(key);
  
      const toggleSort = (key: string) => {
        if (!isSortable(key)) return;
  
        if (currentSort.value.field === key) {
          currentSort.value.order = currentSort.value.order === 'asc' ? 'desc' : 'asc';
        } else {
          currentSort.value.field = key;
          currentSort.value.order = 'asc';
        }
  
        // Notify parent about the sorting change
        props.onSortChange(currentSort.value);
      };
  
      return { searchQuery, headers, filteredData, currentSort, toggleSort, isSortable };
    },
  });
  </script>
  