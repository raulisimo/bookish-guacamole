<template>
  <div>
    <!-- Items Per Page Selector -->
    <div class="items-per-page-selector mb-4 flex justify-end items-center">
      <label for="itemsPerPage" class="mr-2 text-white font-semibold">Items per page:</label>
      <select
        id="itemsPerPage"
        v-model="itemsPerPage"
        class="px-3 py-1 rounded-lg bg-gray-800 text-yellow-300 border border-yellow-400 hover:bg-gray-700"
      >
        <option v-for="option in itemsPerPageOptions" :key="option" :value="option">
          {{ option }}
        </option>
      </select>
    </div>

    <!-- Table -->
    <div class="overflow-x-auto shadow-md rounded-lg bg-opacity-90 backdrop-blur">
      <table class="table-auto w-full border-collapse border border-yellow-400 rounded-lg">
        <thead class="bg-gradient-to-r from-gray-900 via-gray-800 to-gray-900">
          <tr>
            <th
              v-for="(key, index) in headers"
              :key="index"
              class="border px-4 py-2 cursor-pointer text-yellow-300 text-left tracking-wider hover:text-yellow-400 transition-all"
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
          <tr
            v-for="item in paginatedData"
            :key="item.url"
            class="hover:bg-gray-800 hover:scale-105 transition-transform duration-300"
          >
            <td
              v-for="(value, index) in Object.values(item)"
              :key="index"
              class="border px-4 py-2 text-white text-center font-light"
            >
              {{ value }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Pagination Controls -->
    <div class="pagination mt-6 flex justify-center items-center gap-4">
      <button
        :disabled="currentPage === 1"
        @click="previousPage"
        class="px-5 py-2 rounded-full text-white bg-gradient-to-r from-blue-600 to-blue-800 hover:from-blue-500 hover:to-blue-700 disabled:opacity-50 transition-all"
      >
        Previous
      </button>
      <span class="text-white font-semibold text-lg"> {{ currentPage }} / {{ totalPages }} </span>
      <button
        :disabled="currentPage === totalPages"
        @click="nextPage"
        class="px-5 py-2 rounded-full text-white bg-gradient-to-r from-blue-600 to-blue-800 hover:from-blue-500 hover:to-blue-700 disabled:opacity-50 transition-all"
      >
        Next
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref, watch } from 'vue'
import { Planet } from '@/models/planet'
import { Person } from '@/models/person'

type DataItem = Person | Planet

export default defineComponent({
  name: 'DataTable',
  props: {
    data: {
      type: Array as () => DataItem[],
      required: true,
    },
    onSortChange: {
      type: Function,
      required: true,
    },
    searchQuery: {
      type: String,
      default: '',
    },
  },
  setup(props) {
    const currentPage = ref(1) // Track current page
    const itemsPerPage = ref(15) // Track number of items per page
    const itemsPerPageOptions = [10, 15, 20, 30] // Options for items per page

    const currentSort = ref({ field: '', order: 'asc' })

    const headers = computed(() => (props.data[0] ? Object.keys(props.data[0]) : []))

    // Filter data based on search query
    const filteredData = computed(() => {
      if (!props.searchQuery) return props.data
      return props.data.filter((item) =>
        Object.values(item).some((value) =>
          String(value).toLowerCase().includes(props.searchQuery.toLowerCase()),
        ),
      )
    })

    // Paginated data: Slice the filtered data to only show the current page
    const paginatedData = computed(() => {
      const start = (currentPage.value - 1) * itemsPerPage.value
      const end = start + itemsPerPage.value
      return filteredData.value.slice(start, end)
    })

    // Total number of pages
    const totalPages = computed(() => {
      return Math.ceil(filteredData.value.length / itemsPerPage.value)
    })

    const isSortable = (key: string) => ['name', 'created'].includes(key)

    const toggleSort = (key: string) => {
      if (!isSortable(key)) return

      if (currentSort.value.field === key) {
        currentSort.value.order = currentSort.value.order === 'asc' ? 'desc' : 'asc'
      } else {
        currentSort.value.field = key
        currentSort.value.order = 'asc'
      }

      // Notify parent about the sorting change
      props.onSortChange(currentSort.value)
    }

    // Pagination functions
    const previousPage = () => {
      if (currentPage.value > 1) {
        currentPage.value--
      }
    }

    const nextPage = () => {
      if (currentPage.value < totalPages.value) {
        currentPage.value++
      }
    }

    // Watch for changes in itemsPerPage and reset pagination
    watch(itemsPerPage, () => {
      currentPage.value = 1 // Reset to first page
    })

    return {
      headers,
      paginatedData,
      currentSort,
      toggleSort,
      isSortable,
      currentPage,
      totalPages,
      previousPage,
      nextPage,
      itemsPerPage,
      itemsPerPageOptions,
    }
  },
})
</script>

<style scoped>
.table-auto {
  border: 2px solid #444;
  background: radial-gradient(circle at top left, #1b1b2f, #1a1a1d);
  border-radius: 20px;
  overflow: hidden;
}

.table-auto th {
  font-size: 1rem;
  letter-spacing: 0.05rem;
  background-image: linear-gradient(45deg, #333, #222);
}

.table-auto td {
  font-size: 0.9rem;
  text-align: center;
  line-height: 1.5;
}

.table-auto tr {
  transition: all 0.2s ease-in-out;
  background-color: rgba(50, 50, 70, 0.5);
}

.table-auto tr:hover {
  background-color: rgba(80, 80, 100, 0.8);
  transform: translateY(-2px);
}

.pagination button {
  font-size: 1rem;
  font-weight: bold;
  text-transform: uppercase;
}

.pagination span {
  font-size: 1.2rem;
  color: #ffdd00;
}
</style>
