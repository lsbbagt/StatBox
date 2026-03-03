<script lang="ts" setup>
import { ref, computed, markRaw } from 'vue'
import TitleBar from './components/TitleBar.vue'
import Sidebar from './components/Sidebar.vue'
import RightPanel from './components/RightPanel.vue'
import BookmarksView from './views/BookmarksView.vue'
import TemplatesView from './views/TemplatesView.vue'
import SettingsView from './views/SettingsView.vue'

type ModuleType = 'bookmarks' | 'templates' | 'settings'

const sidebarVisible = ref(true)
const rightPanelVisible = ref(true)
const activeModule = ref<ModuleType>('bookmarks')

const currentView = computed(() => {
  switch (activeModule.value) {
    case 'bookmarks': return markRaw(BookmarksView)
    case 'templates': return markRaw(TemplatesView)
    case 'settings': return markRaw(SettingsView)
    default: return markRaw(BookmarksView)
  }
})

const toggleSidebar = () => {
  sidebarVisible.value = !sidebarVisible.value
}

const toggleRightPanel = () => {
  rightPanelVisible.value = !rightPanelVisible.value
}

const selectModule = (module: ModuleType) => {
  activeModule.value = module
}
</script>

<template>
  <v-app>
    <!-- 自定义标题栏 -->
    <TitleBar 
      @toggle-sidebar="toggleSidebar"
      @toggle-right-panel="toggleRightPanel"
    />
    
    <!-- 左侧导航栏 -->
    <Sidebar 
      v-model="sidebarVisible"
      :active-module="activeModule"
      @select="selectModule"
    />
    
    <!-- 主内容区 -->
    <v-main class="main-content">
      <v-container fluid class="fill-height pa-0">
        <v-row no-gutters class="fill-height">
          <!-- 中间主面板 -->
          <v-col class="d-flex flex-column">
            <component :is="currentView" :key="activeModule" />
          </v-col>
          
          <!-- 右侧面板 -->
          <RightPanel 
            v-model="rightPanelVisible"
            :module="activeModule"
          />
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
html {
  overflow-y: hidden !important;
}

.main-content {
  height: calc(100vh - 40px);
  overflow: hidden !important;
}

.main-content > .v-container {
  overflow: hidden !important;
}

.main-content .v-row {
  overflow: hidden !important;
}

/* 中间主面板允许滚动 - 使用更强选择器 */
.main-content .v-col:first-child,
.main-content .v-row > .v-col:first-child {
  overflow-y: auto !important;
  overflow-x: hidden !important;
  height: 100%;
}

/* 右侧面板不滚动 */
.main-content .v-col:last-child,
.main-content .v-row > .v-col:last-child {
  overflow: hidden !important;
}

/* 隐藏内部组件的滚动条 */
.v-list {
  overflow: visible !important;
}
</style>
