<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useBookmarksStore } from '../stores/bookmarks'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime'
import type { ModuleType } from '../types'

const props = defineProps<{
  modelValue: boolean
  module: ModuleType
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const bookmarksStore = useBookmarksStore()
const searchQuery = ref('')

const searchPlaceholder = computed(() => {
  switch (props.module) {
    case 'bookmarks': return '搜索收藏...'
    case 'templates': return '搜索模板...'
    case 'commands': return '搜索命令...'
    default: return '搜索...'
  }
})

// 模拟数据
const templates = ref([
  { name: '数据清洗.R', language: 'R', path: '/R/数据清洗.R' },
  { name: '回归分析.R', language: 'R', path: '/R/回归分析.R' },
  { name: 'pandas处理.py', language: 'Python', path: '/Python/pandas处理.py' },
])

const commands = ref([
  { name: '最优编译.md', description: 'LaTeX编译命令', path: '/LaTeX/最优编译.md' },
  { name: '数据流程.md', description: '数据分析流程', path: '/常用/数据流程.md' },
])

const settings = ref([
  { id: 'general', name: '常规设置', icon: 'mdi-cog' },
  { id: 'appearance', name: '外观', icon: 'mdi-palette' },
  { id: 'shortcuts', name: '快捷键', icon: 'mdi-keyboard' },
])

// 展开的文件夹
const expandedFolders = ref<Set<string>>(new Set(['1', 'R', 'Python']))

const toggleFolder = (folderId: string) => {
  if (expandedFolders.value.has(folderId)) {
    expandedFolders.value.delete(folderId)
  } else {
    expandedFolders.value.add(folderId)
  }
}

const openExternal = (url: string) => {
  console.log('Opening external:', url)
  BrowserOpenURL(url)
}

const onSelect = (item: any) => {
  console.log('Selected:', item)
}

// 过滤收藏夹
const filteredBookmarks = computed(() => {
  if (!searchQuery.value) return bookmarksStore.bookmarks
  
  const query = searchQuery.value.toLowerCase()
  return bookmarksStore.bookmarks.map(folder => ({
    ...folder,
    items: folder.items.filter(item =>
      item.name.toLowerCase().includes(query) ||
      item.url.toLowerCase().includes(query)
    )
  })).filter(folder => folder.items.length > 0)
})
</script>

<template>
  <div v-if="modelValue" class="right-panel">
    <!-- 搜索框 -->
    <div class="search-container">
      <span class="mdi mdi-magnify search-icon"></span>
      <input 
        v-model="searchQuery"
        type="text"
        :placeholder="searchPlaceholder"
        class="search-input"
      />
    </div>
    
    <div class="divider"></div>
    
    <!-- 列表内容 -->
    <div class="list-container">
      <!-- 收藏夹列表 -->
      <div v-if="module === 'bookmarks'" class="list-section">
        <div v-for="folder in filteredBookmarks" :key="folder.id" class="folder">
          <div class="folder-header" @click="toggleFolder(folder.id)">
            <span :class="['mdi', expandedFolders.has(folder.id) ? 'mdi-chevron-down' : 'mdi-chevron-right', 'expand-icon']"></span>
            <span class="mdi mdi-folder folder-icon"></span>
            <span class="folder-name">{{ folder.name }}</span>
          </div>
          <div v-if="expandedFolders.has(folder.id)" class="folder-items">
            <div 
              v-for="item in folder.items" 
              :key="item.id" 
              class="list-item"
              @click="onSelect(item)"
            >
              <span class="mdi mdi-link item-icon"></span>
              <div class="item-content">
                <span class="item-name">{{ item.name }}</span>
                <span class="item-url">{{ item.url }}</span>
              </div>
              <button class="external-btn" @click.stop="openExternal(item.url)" title="在浏览器中打开">
                <span class="mdi mdi-open-in-new"></span>
              </button>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 代码模板列表 -->
      <div v-if="module === 'templates'" class="list-section">
        <div class="folder">
          <div class="folder-header" @click="toggleFolder('R')">
            <span :class="['mdi', expandedFolders.has('R') ? 'mdi-chevron-down' : 'mdi-chevron-right', 'expand-icon']"></span>
            <span class="mdi mdi-language-r folder-icon"></span>
            <span class="folder-name">R</span>
          </div>
          <div v-if="expandedFolders.has('R')" class="folder-items">
            <div 
              v-for="file in templates.filter(t => t.language === 'R')" 
              :key="file.path" 
              class="list-item"
              @click="onSelect(file)"
            >
              <span class="mdi mdi-file-document-outline item-icon"></span>
              <span class="item-name">{{ file.name }}</span>
            </div>
          </div>
        </div>
        
        <div class="folder">
          <div class="folder-header" @click="toggleFolder('Python')">
            <span :class="['mdi', expandedFolders.has('Python') ? 'mdi-chevron-down' : 'mdi-chevron-right', 'expand-icon']"></span>
            <span class="mdi mdi-language-python folder-icon"></span>
            <span class="folder-name">Python</span>
          </div>
          <div v-if="expandedFolders.has('Python')" class="folder-items">
            <div 
              v-for="file in templates.filter(t => t.language === 'Python')" 
              :key="file.path" 
              class="list-item"
              @click="onSelect(file)"
            >
              <span class="mdi mdi-file-document-outline item-icon"></span>
              <span class="item-name">{{ file.name }}</span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- 命令库列表 -->
      <div v-if="module === 'commands'" class="list-section">
        <div 
          v-for="cmd in commands" 
          :key="cmd.path" 
          class="list-item"
          @click="onSelect(cmd)"
        >
          <span class="mdi mdi-file-document-outline item-icon"></span>
          <div class="item-content">
            <span class="item-name">{{ cmd.name }}</span>
            <span class="item-desc">{{ cmd.description }}</span>
          </div>
        </div>
      </div>
      
      <!-- 设置分类 -->
      <div v-if="module === 'settings'" class="list-section">
        <div 
          v-for="setting in settings" 
          :key="setting.id" 
          class="list-item"
          @click="onSelect(setting)"
        >
          <span :class="['mdi', setting.icon, 'item-icon']"></span>
          <span class="item-name">{{ setting.name }}</span>
        </div>
      </div>
    </div>
    
    <!-- 隐藏按钮 -->
    <div class="panel-footer">
      <button class="hide-btn" @click="emit('update:modelValue', false)">
        <span class="mdi mdi-chevron-right"></span>
        隐藏面板
      </button>
    </div>
  </div>
  
  <!-- 展开按钮 -->
  <button v-if="!modelValue" class="expand-btn" @click="emit('update:modelValue', true)">
    <span class="mdi mdi-chevron-left"></span>
  </button>
</template>

<style scoped>
.right-panel {
  width: 280px;
  height: 100%;
  background-color: #F8FAFB;
  border-left: 1px solid #E1E8ED;
  display: flex;
  flex-direction: column;
}

.search-container {
  display: flex;
  align-items: center;
  background-color: #FFFFFF;
  border: 1px solid #E1E8ED;
  border-radius: 8px;
  margin: 12px;
  padding: 8px 12px;
}

.search-icon {
  color: #7A8B99;
  font-size: 18px;
  margin-right: 8px;
}

.search-input {
  border: none;
  background: transparent;
  outline: none;
  font-size: 14px;
  flex: 1;
  color: #2C3E50;
}

.search-input::placeholder {
  color: #9BA8B3;
}

.divider {
  height: 1px;
  background-color: #E1E8ED;
}

.list-container {
  flex: 1;
  overflow-y: auto;
  padding: 8px;
}

.list-section {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.folder {
  margin-bottom: 4px;
}

.folder-header {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 8px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.folder-header:hover {
  background-color: rgba(0, 0, 0, 0.04);
}

.expand-icon {
  font-size: 18px;
  color: #7A8B99;
}

.folder-icon {
  font-size: 18px;
  color: #5B9BD5;
}

.folder-name {
  font-size: 14px;
  font-weight: 500;
  color: #2C3E50;
}

.folder-items {
  margin-left: 16px;
}

.list-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.list-item:hover {
  background-color: rgba(91, 155, 213, 0.1);
}

.item-icon {
  font-size: 18px;
  color: #7A8B99;
  flex-shrink: 0;
}

.item-content {
  flex: 1;
  overflow: hidden;
}

.item-name {
  font-size: 14px;
  color: #2C3E50;
  display: block;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.item-url, .item-desc {
  font-size: 12px;
  color: #9BA8B3;
  display: block;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.external-btn {
  width: 28px;
  height: 28px;
  border: none;
  background: transparent;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #7A8B99;
  transition: all 0.2s;
}

.external-btn:hover {
  background-color: rgba(91, 155, 213, 0.15);
  color: #5B9BD5;
}

.panel-footer {
  border-top: 1px solid #E1E8ED;
  padding: 8px;
}

.hide-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px;
  border: none;
  background: transparent;
  border-radius: 8px;
  cursor: pointer;
  color: #7A8B99;
  font-size: 14px;
  transition: all 0.2s;
}

.hide-btn:hover {
  background-color: rgba(0, 0, 0, 0.04);
  color: #5B9BD5;
}

.expand-btn {
  position: fixed;
  right: 16px;
  top: 60px;
  width: 36px;
  height: 36px;
  border: 1px solid #E1E8ED;
  background-color: #FFFFFF;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #5B9BD5;
  font-size: 20px;
  z-index: 100;
  transition: all 0.2s;
}

.expand-btn:hover {
  background-color: #5B9BD5;
  color: white;
}
</style>
