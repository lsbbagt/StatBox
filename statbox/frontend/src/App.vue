<script lang="ts" setup>
import { ref } from 'vue'
import TitleBar from './components/TitleBar.vue'
import Sidebar from './components/Sidebar.vue'
import RightPanel from './components/RightPanel.vue'

type ModuleType = 'bookmarks' | 'templates' | 'commands' | 'settings'

const sidebarVisible = ref(true)
const rightPanelVisible = ref(true)
const activeModule = ref<ModuleType>('bookmarks')

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
            <v-card flat tile class="fill-height">
              <v-card-title class="text-h6 py-3 px-4 border-b">
                {{ activeModule === 'bookmarks' ? '收藏夹' : 
                   activeModule === 'templates' ? '代码模板' : 
                   activeModule === 'commands' ? '命令库' : '设置' }}
              </v-card-title>
              <v-card-text class="flex-grow-1 overflow-auto">
                <div class="text-center py-8">
                  <v-icon :icon="activeModule === 'bookmarks' ? 'mdi-folder-star' : 
                                 activeModule === 'templates' ? 'mdi-code-tags' : 
                                 activeModule === 'commands' ? 'mdi-keyboard' : 'mdi-cog'" 
                          size="64" color="primary" class="mb-4" />
                  <h2 class="text-h5 mb-2">
                    {{ activeModule === 'bookmarks' ? '收藏夹管理' : 
                       activeModule === 'templates' ? '代码模板管理' : 
                       activeModule === 'commands' ? '命令库' : '系统设置' }}
                  </h2>
                  <p class="text-body-1 text-medium-emphasis">
                    {{ activeModule === 'bookmarks' ? '管理您的网址收藏，支持多标签页浏览器' : 
                       activeModule === 'templates' ? '管理 R / Python / Julia 代码模板' : 
                       activeModule === 'commands' ? '常用命令快捷调用（按 ~ 唤醒）' : '配置应用设置' }}
                  </p>
                </div>
              </v-card-text>
            </v-card>
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
}
</style>
