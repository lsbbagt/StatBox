import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface SelectedItem {
  type: 'bookmark' | 'template' | 'command' | 'setting'
  data: any
}

export const useSelectionStore = defineStore('selection', () => {
  const selectedItem = ref<SelectedItem | null>(null)

  const selectItem = (type: SelectedItem['type'], data: any) => {
    selectedItem.value = { type, data }
  }

  const clearSelection = () => {
    selectedItem.value = null
  }

  return {
    selectedItem,
    selectItem,
    clearSelection
  }
})
