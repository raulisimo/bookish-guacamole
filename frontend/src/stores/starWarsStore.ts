import { defineStore } from 'pinia'
import { HttpClient } from '@/services/httpClient'

const apiClient = new HttpClient(import.meta.env.VITE_API_BASE_URL)

export const useDataStore = defineStore('data', {
  state: () => ({
    data: [] as Array<any>,
    isLoading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchData(type: 'people' | 'planets', sort: string = '', order: string = 'asc') {
      this.isLoading = true
      this.error = null

      try {
        const queryParams = new URLSearchParams()
        if (sort) queryParams.append('sort', sort)
        if (order) queryParams.append('order', order)

        const response = await apiClient.get<any[]>(`/${type}?${queryParams.toString()}`)
        this.data = response
      } catch (error: any) {
        this.error = error.message || 'An unexpected error occurred.'
      } finally {
        this.isLoading = false
      }
    },
  },
})
