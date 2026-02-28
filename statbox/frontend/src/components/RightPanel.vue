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
  <v-navigation-drawer
    :model-value="modelValue"
    @update:model-value="$emit('update:modelValue', $event)"
    location="right"
    width="280"
    class="right-panel"
    :permanent="true"
    color="background"
  >
    <!-- 搜索框 -->
    <v-text-field
      v-model="searchQuery"
      density="compact"
      variant="outlined"
      :placeholder="searchPlaceholder"
      prepend-inner-icon="mdi-magnify"
      class="ma-3 search-field"
      hide-details
      clearable
      bg-color="surface"
    />
    
    <v-divider />
    
    <!-- 列表内容 -->
    <div class="list-container pa-2">
      <!-- 收藏夹列表 -->
      <v-list v-if="module === 'bookmarks'" nav density="compact">
        <v-list-group
          v-for="folder in filteredBookmarks"
          :key="folder.id"
          :value="folder.id"
        >
          <template #activator="{ props: activatorProps }">
            <v-list-item
              v-bind="activatorProps"
              :prepend-icon="folder.icon || 'mdi-folder'"
              :title="folder.name"
            />
          </template>
          
          <v-list-item
            v-for="item in folder.items"
            :key="item.id"
            :title="item.name"
            :subtitle="item.url"
            prepend-icon="mdi-link"
            rounded="lg"
            class="ml-4 mb-1"
            link
            @click="onSelect(item)"
          >
            <template #append>
              <v-btn
                icon="mdi-open-in-new"
                size="x-small"
                variant="text"
                @click.stop="openExternal(item.url)"
              />
            </template>
          </v-list-item>
        </v-list-group>
      </v-list>
      
      <!-- 代码模板列表 -->
      <v-list v-if="module === 'templates'" nav density="compact">
        <v-list-group value="R">
          <template #activator="{ props: activatorProps }">
            <v-list-item
              v-bind="activatorProps"
              prepend-icon="mdi-code-braces"
              title="R"
            />
          </template>
          <v-list-item
            v-for="file in templates.filter(t => t.language === 'R')"
            :key="file.path"
            :title="file.name"
            rounded="lg"
            class="ml-4"
            link
            @click="onSelect(file)"
          />
        </v-list-group>
        
        <v-list-group value="Python">
          <template #activator="{ props: activatorProps }">
            <v-list-item
              v-bind="activatorProps"
              prepend-icon="mdi-language-python"
              title="Python"
            />
          </template>
          <v-list-item
            v-for="file in templates.filter(t => t.language === 'Python')"
            :key="file.path"
            :title="file.name"
            rounded="lg"
            class="ml-4"
            link
            @click="onSelect(file)"
          />
        </v-list-group>
      </v-list>
      
      <!-- 命令库列表 -->
      <v-list v-if="module === 'commands'" nav density="compact">
        <v-list-item
          v-for="cmd in commands"
          :key="cmd.path"
          :title="cmd.name"
          :subtitle="cmd.description"
          prepend-icon="mdi-file-document-outline"
          rounded="lg"
          class="mb-1"
          link
          @click="onSelect(cmd)"
        />
      </v-list>
      
      <!-- 设置分类 -->
      <v-list v-if="module === 'settings'" nav density="compact">
        <v-list-item
          v-for="setting in settings"
          :key="setting.id"
          :title="setting.name"
          :prepend-icon="setting.icon"
          rounded="lg"
          class="mb-1"
          link
          @click="onSelect(setting)"
        />
      </v-list>
    </div>
    
    <!-- 折叠按钮 -->
    <template #append>
      <v-divider />
      <v-btn
        block
        variant="text"
        prepend-icon="mdi-chevron-right"
        @click="$emit('update:modelValue', false)"
        class="collapse-btn"
        color="secondary"
        size="small"
      >
        隐藏面板
      </v-btn>
    </template>
  </v-navigation-drawer>
  
  <!-- 展开按钮（当隐藏时显示） -->
  <v-btn
    v-if="!modelValue"
    icon="mdi-chevron-left"
    size="small"
    variant="outlined"
    class="expand-btn"
    color="primary"
    @click="$emit('update:modelValue', true)"
  />
</template>

<style scoped>
.right-panel {
  border-left: 1px solid rgb(var(--v-theme-border));
}

.search-field {
  font-size: 14px;
}

.list-container {
  height: calc(100% - 140px);
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.list-container::-webkit-scrollbar {
  display: none;
}

.expand-btn {
  position: fixed;
  right: 16px;
  top: 60px;
  z-index: 100;
}
</style>
