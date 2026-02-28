<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { 
  WindowMinimise, 
  WindowToggleMaximise, 
  WindowIsMaximised,
  BrowserOpenURL,
  Quit 
} from '../../wailsjs/runtime/runtime'

defineEmits<{
  (e: 'toggle-sidebar'): void
  (e: 'toggle-right-panel'): void
}>()

const searchQuery = ref('')
const isMaximised = ref(false)

// 检查窗口状态
const checkWindowState = async () => {
  try {
    isMaximised.value = await WindowIsMaximised()
  } catch (e) {
    console.error('Failed to check window state:', e)
  }
}

onMounted(() => {
  checkWindowState()
})

const minimizeWindow = () => {
  console.log('Minimizing...')
  WindowMinimise()
}

const toggleMaximize = async () => {
  console.log('Toggling maximize...')
  WindowToggleMaximise()
  setTimeout(() => {
    checkWindowState()
  }, 100)
}

const closeWindow = () => {
  console.log('Closing...')
  Quit()
}

const openExternal = (url: string) => {
  BrowserOpenURL(url)
}

const onSearch = () => {
  console.log('Search:', searchQuery.value)
}
</script>

<template>
  <div class="title-bar">
    <!-- 拖动区域 -->
    <div class="drag-area">
      <!-- 左侧：菜单按钮和Logo -->
      <div class="left-section">
        <button class="icon-btn" @click="$emit('toggle-sidebar')">
          <span class="mdi mdi-menu"></span>
        </button>
        
        <span class="mdi mdi-chart-box logo-icon"></span>
        <span class="app-title">StatBox</span>
      </div>
      
      <!-- 中间：全局搜索 -->
      <div class="center-section">
        <div class="search-wrapper">
          <span class="mdi mdi-magnify search-icon"></span>
          <input 
            v-model="searchQuery"
            type="text"
            placeholder="全局搜索..."
            class="search-input"
            @keyup.enter="onSearch"
          />
        </div>
      </div>
      
      <!-- 右侧：窗口控制按钮 -->
      <div class="right-section">
        <button class="window-btn" @click="minimizeWindow" title="最小化">
          <span class="mdi mdi-minus"></span>
        </button>
        <button class="window-btn" @click="toggleMaximize" :title="isMaximised ? '还原' : '最大化'">
          <span :class="isMaximised ? 'mdi mdi-fullscreen-exit' : 'mdi mdi-fullscreen'"></span>
        </button>
        <button class="window-btn close-btn" @click="closeWindow" title="关闭">
          <span class="mdi mdi-close"></span>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.title-bar {
  height: 40px;
  background-color: #FFFFFF;
  border-bottom: 1px solid #E1E8ED;
  display: flex;
  align-items: center;
  user-select: none;
}

.drag-area {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  height: 100%;
  padding: 0 8px;
}

.left-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.logo-icon {
  font-size: 20px;
  color: #5B9BD5;
}

.app-title {
  font-size: 14px;
  font-weight: 500;
  color: #5B9BD5;
}

.center-section {
  flex: 1;
  display: flex;
  justify-content: center;
  padding: 0 20px;
}

.search-wrapper {
  display: flex;
  align-items: center;
  background-color: #F8FAFB;
  border: 1px solid #E1E8ED;
  border-radius: 8px;
  padding: 4px 12px;
  width: 100%;
  max-width: 400px;
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
  width: 100%;
  color: #2C3E50;
}

.search-input::placeholder {
  color: #9BA8B3;
}

.right-section {
  display: flex;
  align-items: center;
  gap: 4px;
}

.icon-btn {
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  border-radius: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #5B9BD5;
  font-size: 20px;
  transition: background-color 0.2s;
}

.icon-btn:hover {
  background-color: rgba(91, 155, 213, 0.1);
}

.window-btn {
  width: 40px;
  height: 32px;
  border: none;
  background: transparent;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #2C3E50;
  font-size: 18px;
  transition: background-color 0.2s;
}

.window-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.close-btn:hover {
  background-color: #EF5350;
  color: white;
}
</style>
