import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { BookmarkFolder, BookmarkItem } from '../types'

export const useBookmarksStore = defineStore('bookmarks', () => {
  const bookmarks = ref<BookmarkFolder[]>([
    {
      id: '1',
      name: '统计学资源',
      icon: '📊',
      order: 0,
      items: [
        {
          id: '1-1',
          name: 'R Documentation',
          url: 'https://www.r-project.org/',
          icon: '🔗',
          tags: ['R', '文档'],
          createdAt: '2024-01-01',
          visitCount: 15
        },
        {
          id: '1-2',
          name: 'Python Docs',
          url: 'https://docs.python.org/',
          icon: '🔗',
          tags: ['Python', '文档'],
          createdAt: '2024-01-02',
          visitCount: 10
        }
      ]
    },
    {
      id: '2',
      name: '数据分析',
      icon: '📈',
      order: 1,
      items: [
        {
          id: '2-1',
          name: 'Kaggle',
          url: 'https://www.kaggle.com/',
          icon: '🔗',
          tags: ['数据科学', '竞赛'],
          createdAt: '2024-01-03',
          visitCount: 20
        }
      ]
    }
  ])

  const searchQuery = ref('')

  const filteredBookmarks = computed(() => {
    if (!searchQuery.value) return bookmarks.value

    const query = searchQuery.value.toLowerCase()
    return bookmarks.value.map(folder => ({
      ...folder,
      items: folder.items.filter(item =>
        item.name.toLowerCase().includes(query) ||
        item.url.toLowerCase().includes(query) ||
        item.tags?.some(tag => tag.toLowerCase().includes(query))
      )
    })).filter(folder => folder.items.length > 0)
  })

  const addBookmark = (folderId: string, item: BookmarkItem) => {
    const folder = bookmarks.value.find(f => f.id === folderId)
    if (folder) {
      folder.items.push(item)
    }
  }

  const updateBookmark = (folderId: string, itemId: string, updates: Partial<BookmarkItem>) => {
    const folder = bookmarks.value.find(f => f.id === folderId)
    if (folder) {
      const item = folder.items.find(i => i.id === itemId)
      if (item) {
        Object.assign(item, updates)
      }
    }
  }

  const removeBookmark = (folderId: string, itemId: string) => {
    const folder = bookmarks.value.find(f => f.id === folderId)
    if (folder) {
      folder.items = folder.items.filter(i => i.id !== itemId)
    }
  }

  return {
    bookmarks,
    searchQuery,
    filteredBookmarks,
    addBookmark,
    updateBookmark,
    removeBookmark
  }
})
