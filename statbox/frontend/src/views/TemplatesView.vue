<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { useSelectionStore } from '../stores/selection'
import { OpenFileWithDefault, GetTemplatePath } from '../../wailsjs/go/main/App'

const selectionStore = useSelectionStore()

const code = ref('')
const fileName = ref('选择一个模板')
const fileLanguage = ref('R')

// 监听选择变化
watch(() => selectionStore.selectedItem, (item) => {
  if (item?.type === 'template' && item.data) {
    fileName.value = item.data.name
    fileLanguage.value = item.data.language
    code.value = item.data.code || '// 暂无代码内容'
  }
}, { immediate: true })

// 复制代码
const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(code.value)
  } catch {
    // 备用方案
    const textarea = document.createElement('textarea')
    textarea.value = code.value
    document.body.appendChild(textarea)
    textarea.select()
    document.execCommand('copy')
    document.body.removeChild(textarea)
  }
}

// 在外部IDE打开（使用系统默认应用）
const openInExternalIDE = async () => {
  const item = selectionStore.selectedItem
  if (item?.type === 'template' && item.data?.path) {
    try {
      // 获取完整文件路径
      const fullPath = await GetTemplatePath(item.data.path)
      await OpenFileWithDefault(fullPath)
    } catch (error) {
      console.error('打开文件失败:', error)
      alert('打开文件失败: ' + error)
    }
  } else {
    alert('模板文件路径未配置')
  }
}

const languageIcon = computed(() => {
  switch (fileLanguage.value) {
    case 'R': return 'mdi-code-braces'
    case 'Python': return 'mdi-language-python'
    case 'Julia': return 'mdi-language-julia'
    default: return 'mdi-code-tags'
  }
})
</script>

<template>
  <div class="templates-view h-100 d-flex flex-column">
    <!-- 编辑器头部 -->
    <div class="editor-header pa-3 d-flex align-center" style="background: rgb(var(--v-theme-surface)); border-bottom: 1px solid rgb(var(--v-theme-border));">
      <v-icon :icon="languageIcon" color="primary" class="mr-2" />
      <span class="text-subtitle-1 font-weight-medium">{{ fileName }}</span>
      <v-spacer />
      <v-btn size="small" variant="outlined" class="mr-2" prepend-icon="mdi-content-copy" @click="copyCode">
        复制
      </v-btn>
      <v-btn size="small" variant="outlined" prepend-icon="mdi-open-in-new" @click="openInExternalIDE">
        外部打开
      </v-btn>
    </div>

    <!-- 代码编辑器 -->
    <div class="code-editor flex-grow-1 overflow-auto" style="background: rgb(var(--v-theme-background)); padding: 16px;">
      <pre style="font-family: 'Consolas', 'Monaco', monospace; font-size: 14px; line-height: 1.7; color: rgb(var(--v-theme-on-surface));">{{ code || '// 选择左侧模板查看代码' }}</pre>
    </div>
  </div>
</template>

<style scoped>
.code-editor pre {
  margin: 0;
  white-space: pre-wrap;
}
</style>
