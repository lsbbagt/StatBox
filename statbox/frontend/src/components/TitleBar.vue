<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { 
  WindowMinimise, 
  WindowToggleMaximise, 
  WindowIsMaximised,
  Quit 
} from '../../wailsjs/runtime/runtime'

defineEmits<{
  (e: 'toggle-sidebar'): void
  (e: 'toggle-right-panel'): void
}>()

const searchQuery = ref('')
const isMaximised = ref(false)

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

const onSearch = () => {
  console.log('Search:', searchQuery.value)
}
</script>

<template>
  <v-app-bar
    height="40"
    color="surface"
    class="title-bar"
    flat
    :elevation="0"
  >
    <!-- 左侧：菜单按钮和Logo -->
    <v-btn
      icon="mdi-menu"
      size="small"
      variant="text"
      color="primary"
      @click="$emit('toggle-sidebar')"
    />
    
    <v-icon icon="mdi-chart-box" class="ml-2" color="primary" size="20" />
    <span class="text-subtitle-2 ml-2 font-weight-medium text-primary">StatBox</span>
    
    <v-spacer />
    
    <!-- 中间：全局搜索 -->
    <v-text-field
      v-model="searchQuery"
      density="compact"
      variant="outlined"
      placeholder="全局搜索..."
      prepend-inner-icon="mdi-magnify"
      class="mx-4 search-field"
      hide-details
      bg-color="background"
      @keyup.enter="onSearch"
    />
    
    <v-spacer />
    
    <!-- 右侧：窗口控制按钮 -->
    <v-btn
      icon="mdi-minus"
      size="small"
      variant="text"
      @click="minimizeWindow"
    />
    <v-btn
      :icon="isMaximised ? 'mdi-fullscreen-exit' : 'mdi-fullscreen'"
      size="small"
      variant="text"
      @click="toggleMaximize"
    />
    <v-btn
      icon="mdi-close"
      size="small"
      variant="text"
      color="error"
      @click="closeWindow"
    />
  </v-app-bar>
</template>

<style scoped>
.title-bar {
  user-select: none;
  border-bottom: 1px solid rgb(var(--v-theme-border));
}

.search-field {
  max-width: 400px;
}
</style>
