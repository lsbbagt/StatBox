import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface BrowserTab {
  id: number
  title: string
  url: string
  favicon?: string
}

export const useBrowserStore = defineStore('browser', () => {
  const tabs = ref<BrowserTab[]>([
    { id: 1, title: '欢迎使用', url: 'about:blank' }
  ])
  
  const activeTabId = ref(1)
  const currentUrl = ref('')

  const activeTab = computed(() => 
    tabs.value.find(t => t.id === activeTabId.value)
  )

  const openTab = (url: string, title?: string) => {
    const id = Date.now()
    tabs.value.push({
      id,
      title: title || url,
      url
    })
    activeTabId.value = id
    currentUrl.value = url
  }

  const closeTab = (id: number) => {
    const index = tabs.value.findIndex(t => t.id === id)
    if (index > -1) {
      tabs.value.splice(index, 1)
      
      // 如果关闭的是当前标签，切换到前一个或后一个
      if (activeTabId.value === id && tabs.value.length > 0) {
        const newIndex = Math.min(index, tabs.value.length - 1)
        activeTabId.value = tabs.value[newIndex].id
        currentUrl.value = tabs.value[newIndex].url
      }
    }
  }

  const switchTab = (id: number) => {
    activeTabId.value = id
    const tab = tabs.value.find(t => t.id === id)
    if (tab) {
      currentUrl.value = tab.url
    }
  }

  const updateTabUrl = (id: number, url: string) => {
    const tab = tabs.value.find(t => t.id === id)
    if (tab) {
      tab.url = url
      tab.title = url // 简化处理，实际应该从页面获取标题
    }
    currentUrl.value = url
  }

  return {
    tabs,
    activeTabId,
    currentUrl,
    activeTab,
    openTab,
    closeTab,
    switchTab,
    updateTabUrl
  }
})
