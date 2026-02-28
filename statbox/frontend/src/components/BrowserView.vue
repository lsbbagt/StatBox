<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { useBrowserStore } from '../stores/browser'

const browserStore = useBrowserStore()

const addressBarUrl = ref('')

// 同步地址栏和当前URL
watch(() => browserStore.currentUrl, (newUrl) => {
  addressBarUrl.value = newUrl
}, { immediate: true })

const navigateToUrl = () => {
  if (addressBarUrl.value) {
    let url = addressBarUrl.value
    // 自动添加 https://
    if (!url.startsWith('http://') && !url.startsWith('https://')) {
      url = 'https://' + url
    }
    browserStore.updateTabUrl(browserStore.activeTabId, url)
  }
}

const goBack = () => {
  console.log('Go back')
}

const goForward = () => {
  console.log('Go forward')
}

const refresh = () => {
  console.log('Refresh')
}

const openInExternal = () => {
  if (browserStore.currentUrl) {
    window.open(browserStore.currentUrl, '_blank')
  }
}
</script>

<template>
  <div class="browser-view h-100 d-flex flex-column">
    <!-- 工具栏 -->
    <div class="browser-toolbar pa-2 d-flex align-center" style="background: rgb(var(--v-theme-surface)); border-bottom: 1px solid rgb(var(--v-theme-border));">
      <!-- 导航按钮 -->
      <v-btn icon="mdi-arrow-left" size="small" variant="text" @click="goBack" />
      <v-btn icon="mdi-arrow-right" size="small" variant="text" @click="goForward" />
      <v-btn icon="mdi-refresh" size="small" variant="text" @click="refresh" />
      
      <!-- 地址栏 -->
      <v-text-field
        v-model="addressBarUrl"
        density="compact"
        variant="outlined"
        placeholder="输入网址..."
        class="mx-2 flex-grow-1"
        hide-details
        @keyup.enter="navigateToUrl"
      >
        <template #prepend-inner>
          <v-icon size="small" color="secondary">mdi-lock</v-icon>
        </template>
      </v-text-field>
      
      <!-- 外部浏览器按钮 -->
      <v-btn
        icon="mdi-open-in-new"
        size="small"
        variant="text"
        color="primary"
        @click="openInExternal"
      />
    </div>
    
    <!-- 浏览器内容 -->
    <div class="browser-content flex-grow-1" style="background: rgb(var(--v-theme-background));">
      <iframe
        v-if="browserStore.currentUrl && browserStore.currentUrl !== 'about:blank'"
        :src="browserStore.currentUrl"
        class="w-100 h-100"
        style="border: none;"
        sandbox="allow-same-origin allow-scripts allow-popups allow-forms"
      />
      <div v-else class="d-flex align-center justify-center h-100">
        <v-card class="pa-8 text-center" max-width="500" elevation="0">
          <v-icon icon="mdi-web" size="64" color="primary" class="mb-4" />
          <h2 class="text-h5 mb-2">欢迎使用浏览器</h2>
          <p class="text-body-1 text-secondary">
            在地址栏输入网址开始浏览，或点击右侧收藏夹快速打开
          </p>
        </v-card>
      </div>
    </div>
  </div>
</template>

<style scoped>
.browser-toolbar {
  min-height: 48px;
}
</style>
