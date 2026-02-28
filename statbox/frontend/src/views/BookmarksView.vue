<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useBrowserStore } from '../stores/browser'
import { useBookmarksStore } from '../stores/bookmarks'
import BrowserView from '../components/BrowserView.vue'

const browserStore = useBrowserStore()
const bookmarksStore = useBookmarksStore()

const activeTab = computed(() => browserStore.activeTab)

const openBookmark = (url: string, title: string) => {
  // 如果当前标签是空白页，在当前标签打开
  if (browserStore.currentUrl === 'about:blank') {
    browserStore.updateTabUrl(browserStore.activeTabId, url)
  } else {
    // 否则在新标签页打开
    browserStore.openTab(url, title)
  }
}

const closeTab = (id: number) => {
  browserStore.closeTab(id)
}

const switchTab = (id: number) => {
  browserStore.switchTab(id)
}
</script>

<template>
  <div class="bookmarks-view h-100 d-flex flex-column">
    <!-- 浏览器标签栏 -->
    <v-tabs
      v-model="browserStore.activeTabId"
      bg-color="surface"
      class="browser-tabs"
      show-arrows
    >
      <v-tab
        v-for="tab in browserStore.tabs"
        :key="tab.id"
        :value="tab.id"
        class="text-none"
        @click="switchTab(tab.id)"
      >
        <v-icon start size="small">mdi-web</v-icon>
        {{ tab.title }}
        <v-btn
          v-if="browserStore.tabs.length > 1"
          icon="mdi-close"
          size="x-small"
          variant="text"
          class="ml-2"
          @click.stop="closeTab(tab.id)"
        />
      </v-tab>
      
      <v-btn
        icon="mdi-plus"
        size="small"
        variant="text"
        class="ml-2"
        color="primary"
        @click="browserStore.openTab('about:blank', '新标签页')"
      />
    </v-tabs>
    
    <!-- 浏览器视图 -->
    <BrowserView />
  </div>
</template>

<style scoped>
.browser-tabs {
  border-bottom: 1px solid rgb(var(--v-theme-border));
}
</style>
