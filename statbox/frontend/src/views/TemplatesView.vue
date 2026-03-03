<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { useSelectionStore } from '../stores/selection'
import { BrowserOpenURL } from '../../wailsjs/runtime/runtime'

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

const terminalOutput = ref('点击"运行"按钮执行代码...')
const terminalHeight = ref(180)
const isDragging = ref(false)

// 运行代码
const runCode = async () => {
  terminalOutput.value = `> 运行 ${fileName.value}...\n\n正在执行...\n\n[完成]`
}

// 在外部IDE打开
const openInExternalIDE = () => {
  // 在实际应用中，这里会调用 Go 后端打开文件
  console.log('Opening in external IDE:', fileName.value)
  alert('此功能需要配置文件路径，目前为演示模式')
}

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

// 清空终端
const clearTerminal = () => {
  terminalOutput.value = ''
}

// 拖动调整终端高度
const startDrag = (e: MouseEvent) => {
  isDragging.value = true
  const startY = e.clientY
  const startHeight = terminalHeight.value

  const onMouseMove = (e: MouseEvent) => {
    if (!isDragging.value) return
    const diff = startY - e.clientY
    terminalHeight.value = Math.max(100, Math.min(400, startHeight + diff))
  }

  const onMouseUp = () => {
    isDragging.value = false
    document.removeEventListener('mousemove', onMouseMove)
    document.removeEventListener('mouseup', onMouseUp)
  }

  document.addEventListener('mousemove', onMouseMove)
  document.addEventListener('mouseup', onMouseUp)
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
      <v-btn size="small" variant="outlined" class="mr-2" prepend-icon="mdi-play" @click="runCode">
        运行
      </v-btn>
      <v-btn size="small" variant="outlined" class="mr-2" prepend-icon="mdi-content-copy" @click="copyCode">
        复制
      </v-btn>
      <v-btn size="small" variant="outlined" prepend-icon="mdi-open-in-new" @click="openInExternalIDE">
        外部IDE
      </v-btn>
    </div>

    <!-- 代码编辑器 -->
    <div class="code-editor flex-grow-1 overflow-auto" style="background: rgb(var(--v-theme-background)); padding: 16px;">
      <pre style="font-family: 'Consolas', 'Monaco', monospace; font-size: 14px; line-height: 1.7; color: rgb(var(--v-theme-on-surface));">{{ code || '// 选择左侧模板查看代码' }}</pre>
    </div>

    <!-- 终端输出区域 -->
    <div
      class="terminal-output"
      :style="{
        background: 'rgb(var(--v-theme-surface-variant))',
        borderTop: '1px solid rgb(var(--v-theme-border))',
        height: terminalHeight + 'px'
      }"
    >
      <!-- 可拖动的分隔线 -->
      <div
        class="terminal-dragger"
        @mousedown="startDrag"
        :style="{ cursor: isDragging ? 'ns-resize' : 'row-resize' }"
      />

      <div class="pa-3">
        <div class="d-flex align-center mb-2">
          <v-icon icon="mdi-console" size="small" color="secondary" class="mr-2" />
          <span class="text-caption font-weight-medium">终端输出</span>
          <v-spacer />
          <v-btn size="x-small" variant="text" @click="clearTerminal">清空</v-btn>
        </div>
        <div
          class="output-content overflow-auto"
          :style="{
            height: (terminalHeight - 60) + 'px',
            fontFamily: 'Consolas, Monaco, monospace',
            fontSize: '13px',
            color: 'rgb(var(--v-theme-on-surface))'
          }"
        >
          <pre style="margin: 0; white-space: pre-wrap;">{{ terminalOutput }}</pre>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.code-editor pre {
  margin: 0;
  white-space: pre-wrap;
}

.terminal-dragger {
  height: 6px;
  background: rgb(var(--v-theme-border));
  cursor: row-resize;
  transition: background 0.2s;
}

.terminal-dragger:hover {
  background: rgb(var(--v-theme-primary));
}

.output-content pre {
  margin: 0;
}
</style>
