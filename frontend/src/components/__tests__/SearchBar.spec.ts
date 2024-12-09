import { mount } from '@vue/test-utils'
import { describe, it, expect, vi } from 'vitest'
import SearchBar from '@/components/SearchBar.vue'

describe('SearchBar Component', () => {
  it('should update searchQuery and emit update:searchQuery event on input change', async () => {
    // Create a mock function for the emit
    const emit = vi.fn()

    // Mount the component
    const wrapper = mount(SearchBar, {
      global: {
        mocks: {
          emit
        }
      }
    })

    // Find the input element
    const input = wrapper.find('input')

    // Simulate typing in the input
    await input.setValue('New search query')

    // Assert that the searchQuery is updated
    expect(wrapper.vm.searchQuery).toBe('New search query')

    // Assert that the correct event was emitted with the updated query
    expect(emit).toHaveBeenCalledWith('update:searchQuery', 'New search query')
  })
})
