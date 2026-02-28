<script lang="ts" setup>
import { ref, computed } from 'vue'
import TitleBar from './components/TitleBar.vue'
import Sidebar from './components/Sidebar.vue'
import RightPanel from './components/RightPanel.vue'
import BookmarksView from './views/BookmarksView.vue'
import TemplatesView from './views/TemplatesView.vue'
import CommandsView from './views/CommandsView.vue'
import SettingsView from './views/SettingsView.vue'
import type { ModuleType } from './types'

const sidebarVisible = ref(true)
const rightPanelVisible = ref(true)
const activeModule = ref<ModuleType>('bookmarks')

const currentView = computed(() => {
  switch (activeModule.value) {
    case 'bookmarks': return BookmarksView
    case 'templates': return TemplatesView
    case 'commands': return CommandsView
    case 'settings': return SettingsView
    default: return BookmarksView
  }
})
</script>

<template>
  <v-app>
    <!-- 自定义标题栏 -->
    <TitleBar 
      @toggle-sidebar="sidebarVisible = !sidebarVisible"
      @toggle-right-panel="rightPanelVisible = !rightPanelVisible"
    />
    
    <!-- 主内容区域 -->
    <v-main class="main-content">
      <div class="d-flex h-100">
        <!-- 左侧模块选择器 -->
        <Sidebar
          v-model="sidebarVisible"
          :active-module="activeModule"
          @select="activeModule = $event"
        />
        
        <!-- 中间主面板 -->
        <div class="flex-grow-1 content-wrapper">
          <component :is="currentView" />
        </div>
        
        <!-- 右侧列表面板 -->
        <RightPanel
          v-model="rightPanelVisible"
          :module="activeModule"
        />
      </div>
    </v-main>
  </v-app>
</template>

<style>
/* 隐藏滚动条 */
html {
  overflow-y: hidden !important;
}

.main-content {
  height: calc(100vh - 40px);
  overflow: hidden;
  background: rgb(var(--v-theme-background));
}

.content-wrapper {
  height: 100%;
  overflow-y: auto;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.content-wrapper::-webkit-scrollbar {
  display: none;
}
</style>
