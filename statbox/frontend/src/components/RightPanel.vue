<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useBookmarksStore } from '../stores/bookmarks'
import { useBrowserStore } from '../stores/browser'
import { useSelectionStore } from '../stores/selection'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime'
import type { ModuleType, BookmarkItem, BookmarkFolder } from '../types'

const props = defineProps<{
  modelValue: boolean
  module: ModuleType
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const bookmarksStore = useBookmarksStore()
const browserStore = useBrowserStore()
const selectionStore = useSelectionStore()

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
  { id: 't1', name: '数据清洗.R', language: 'R', path: '/R/数据清洗.R', code: 'library(dplyr)\nlibrary(tidyr)\n\n# 处理缺失值\ndata <- data %>% drop_na()' },
  { id: 't2', name: '回归分析.R', language: 'R', path: '/R/回归分析.R', code: '# 回归分析模板\nmodel <- lm(y ~ x, data = df)\nsummary(model)' },
  { id: 't3', name: 'pandas处理.py', language: 'Python', path: '/Python/pandas处理.py', code: 'import pandas as pd\nimport numpy as np\n\n# 数据处理\ndf = pd.read_csv("data.csv")' },
])

const commands = ref([
  { id: 'c1', name: '最优编译.md', description: 'LaTeX编译命令', path: '/LaTeX/最优编译.md', content: 'pdflatex -interaction=nonstopmode -shell-escape main.tex\nbiber main\npdflatex main.tex' },
  { id: 'c2', name: '数据流程.md', description: '数据分析流程', path: '/常用/数据流程.md', content: '# 数据分析标准流程\n1. 数据清洗\n2. 探索性分析\n3. 建模\n4. 验证' },
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

// 无法在 iframe 中加载的网站列表
const externalOnlyDomains = [
  'kaggle.com',
  'github.com',
  'stackoverflow.com',
  'google.com',
  'youtube.com',
]

// 检查是否需要在外部浏览器打开
const needsExternalBrowser = (url: string): boolean => {
  try {
    const hostname = new URL(url).hostname.replace('www.', '')
    return externalOnlyDomains.some(domain => hostname.includes(domain))
  } catch {
    return false
  }
}

// 打开收藏项
const openBookmark = (item: BookmarkItem) => {
  console.log('Opening bookmark:', item)

  // 检查是否需要在外部浏览器打开
  if (needsExternalBrowser(item.url)) {
    openExternal(item.url)
    return
  }

  // 在应用内浏览器打开
  const activeTab = browserStore.activeTab
  if (activeTab && activeTab.url === 'about:blank') {
    browserStore.updateTabUrl(browserStore.activeTabId, item.url)
  } else {
    browserStore.openTab(item.url, item.name)
  }
}

// 选择模板
const selectTemplate = (template: any) => {
  selectionStore.selectItem('template', template)
}

// 选择命令
const selectCommand = (cmd: any) => {
  selectionStore.selectItem('command', cmd)
}

// 选择设置
const selectSetting = (setting: any) => {
  selectionStore.selectItem('setting', setting)
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

// ========== 增删改功能 ==========
const showAddDialog = ref(false)
const showEditDialog = ref(false)
const showAddFolderDialog = ref(false)

const editingItem = ref<BookmarkItem | null>(null)
const editingFolderId = ref<string>('')
const newItem = ref({ name: '', url: '', tags: '' })

// 添加收藏项
const openAddDialog = (folderId: string) => {
  editingFolderId.value = folderId
  newItem.value = { name: '', url: '', tags: '' }
  showAddDialog.value = true
}

const addBookmark = () => {
  if (!newItem.value.name || !newItem.value.url) return
  
  const item: BookmarkItem = {
    id: `${editingFolderId.value}-${Date.now()}`,
    name: newItem.value.name,
    url: newItem.value.url.startsWith('http') ? newItem.value.url : `https://${newItem.value.url}`,
    icon: '🔗',
    tags: newItem.value.tags.split(',').map(t => t.trim()).filter(Boolean),
    createdAt: new Date().toISOString().split('T')[0],
    visitCount: 0
  }
  
  bookmarksStore.addBookmark(editingFolderId.value, item)
  showAddDialog.value = false
}

// 编辑收藏项
const openEditDialog = (folderId: string, item: BookmarkItem) => {
  editingFolderId.value = folderId
  editingItem.value = { ...item }
  newItem.value = {
    name: item.name,
    url: item.url,
    tags: item.tags?.join(', ') || ''
  }
  showEditDialog.value = true
}

const editBookmark = () => {
  if (!editingItem.value || !newItem.value.name || !newItem.value.url) return
  
  const url = newItem.value.url.startsWith('http') ? newItem.value.url : `https://${newItem.value.url}`
  const tags = newItem.value.tags.split(',').map(t => t.trim()).filter(Boolean)
  
  bookmarksStore.updateBookmark(editingFolderId.value, editingItem.value.id, {
    name: newItem.value.name,
    url,
    tags
  })
  
  showEditDialog.value = false
}

// 删除收藏项
const deleteBookmark = (folderId: string, itemId: string) => {
  if (confirm('确定要删除这个收藏吗？')) {
    bookmarksStore.removeBookmark(folderId, itemId)
  }
}

// 新建文件夹
const newFolderName = ref('')
const openAddFolderDialog = () => {
  newFolderName.value = ''
  showAddFolderDialog.value = true
}

const addFolder = () => {
  if (!newFolderName.value.trim()) return
  
  const folder: BookmarkFolder = {
    id: `folder-${Date.now()}`,
    name: newFolderName.value.trim(),
    icon: '📁',
    order: bookmarksStore.bookmarks.length,
    items: []
  }
  
  bookmarksStore.bookmarks.push(folder)
  showAddFolderDialog.value = false
}

// ========== 右键菜单 ==========
const contextMenu = ref({
  show: false,
  x: 0,
  y: 0,
  folderId: '',
  item: null as BookmarkItem | null
})

const showContextMenu = (event: MouseEvent, folderId: string, item: BookmarkItem) => {
  contextMenu.value = {
    show: true,
    x: event.clientX,
    y: event.clientY,
    folderId,
    item
  }
}

const closeContextMenu = () => {
  contextMenu.value.show = false
}

const contextMenuEdit = () => {
  if (contextMenu.value.item) {
    openEditDialog(contextMenu.value.folderId, contextMenu.value.item)
  }
  closeContextMenu()
}

const contextMenuDelete = () => {
  if (contextMenu.value.item) {
    deleteBookmark(contextMenu.value.folderId, contextMenu.value.item.id)
  }
  closeContextMenu()
}

const contextMenuOpenExternal = () => {
  if (contextMenu.value.item) {
    openExternal(contextMenu.value.item.url)
  }
  closeContextMenu()
}
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
        <!-- 操作按钮 -->
        <div class="d-flex ga-1 mb-2 px-1">
          <v-btn 
            size="small" 
            variant="tonal" 
            color="primary" 
            prepend-icon="mdi-folder-plus"
            @click="openAddFolderDialog"
            class="flex-grow-1"
          >
            新建文件夹
          </v-btn>
        </div>
        
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
            >
              <template #append>
                <v-btn
                  icon="mdi-plus"
                  size="x-small"
                  variant="text"
                  @click.stop="openAddDialog(folder.id)"
                />
              </template>
            </v-list-item>
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
            @click="openBookmark(item)"
            @contextmenu.prevent="showContextMenu($event, folder.id, item)"
          >
            <template #append>
              <v-btn
                icon="mdi-open-in-new"
                size="x-small"
                variant="text"
                @click.stop="openExternal(item.url)"
              />
              <v-btn
                icon="mdi-delete"
                size="x-small"
                variant="text"
                color="error"
                @click.stop="deleteBookmark(folder.id, item.id)"
              />
            </template>
          </v-list-item>
          
          <!-- 空文件夹提示 -->
          <v-list-item v-if="folder.items.length === 0" class="ml-4 text-medium-emphasis">
            <v-list-item-title>空文件夹</v-list-item-title>
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
            :key="file.id"
            :title="file.name"
            rounded="lg"
            class="ml-4"
            link
            @click="selectTemplate(file)"
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
            :key="file.id"
            :title="file.name"
            rounded="lg"
            class="ml-4"
            link
            @click="selectTemplate(file)"
          />
        </v-list-group>
      </v-list>

      <!-- 命令库列表 -->
      <v-list v-if="module === 'commands'" nav density="compact">
        <v-list-item
          v-for="cmd in commands"
          :key="cmd.id"
          :title="cmd.name"
          :subtitle="cmd.description"
          prepend-icon="mdi-file-document-outline"
          rounded="lg"
          class="mb-1"
          link
          @click="selectCommand(cmd)"
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
          @click="selectSetting(setting)"
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
  
  <!-- 添加收藏对话框 -->
  <v-dialog v-model="showAddDialog" max-width="400">
    <v-card>
      <v-card-title class="text-h6">添加收藏</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="newItem.name"
          label="名称"
          variant="outlined"
          density="compact"
          class="mb-3"
        />
        <v-text-field
          v-model="newItem.url"
          label="网址"
          variant="outlined"
          density="compact"
          class="mb-3"
          placeholder="https://example.com"
        />
        <v-text-field
          v-model="newItem.tags"
          label="标签（用逗号分隔）"
          variant="outlined"
          density="compact"
          placeholder="文档, 教程"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn variant="text" @click="showAddDialog = false">取消</v-btn>
        <v-btn variant="flat" color="primary" @click="addBookmark">添加</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  
  <!-- 编辑收藏对话框 -->
  <v-dialog v-model="showEditDialog" max-width="400">
    <v-card>
      <v-card-title class="text-h6">编辑收藏</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="newItem.name"
          label="名称"
          variant="outlined"
          density="compact"
          class="mb-3"
        />
        <v-text-field
          v-model="newItem.url"
          label="网址"
          variant="outlined"
          density="compact"
          class="mb-3"
        />
        <v-text-field
          v-model="newItem.tags"
          label="标签（用逗号分隔）"
          variant="outlined"
          density="compact"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn variant="text" @click="showEditDialog = false">取消</v-btn>
        <v-btn variant="flat" color="primary" @click="editBookmark">保存</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  
  <!-- 新建文件夹对话框 -->
  <v-dialog v-model="showAddFolderDialog" max-width="400">
    <v-card>
      <v-card-title class="text-h6">新建文件夹</v-card-title>
      <v-card-text>
        <v-text-field
          v-model="newFolderName"
          label="文件夹名称"
          variant="outlined"
          density="compact"
          placeholder="例如：学习资源"
        />
      </v-card-text>
      <v-card-actions>
        <v-spacer />
        <v-btn variant="text" @click="showAddFolderDialog = false">取消</v-btn>
        <v-btn variant="flat" color="primary" @click="addFolder">创建</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
  
  <!-- 右键菜单 -->
  <v-menu
    v-model="contextMenu.show"
    :position-x="contextMenu.x"
    :position-y="contextMenu.y"
    absolute
    offset-y
  >
    <v-list density="compact" nav>
      <v-list-item @click="contextMenuEdit">
        <template #prepend>
          <v-icon icon="mdi-pencil" />
        </template>
        <v-list-item-title>编辑</v-list-item-title>
      </v-list-item>
      <v-list-item @click="contextMenuOpenExternal">
        <template #prepend>
          <v-icon icon="mdi-open-in-new" />
        </template>
        <v-list-item-title>外部浏览器打开</v-list-item-title>
      </v-list-item>
      <v-divider />
      <v-list-item @click="contextMenuDelete" class="text-error">
        <template #prepend>
          <v-icon icon="mdi-delete" color="error" />
        </template>
        <v-list-item-title>删除</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
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
