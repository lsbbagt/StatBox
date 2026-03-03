<script lang="ts" setup>
import { ref, watch, computed } from 'vue'
import { useSelectionStore } from '../stores/selection'
import { useSettingsStore } from '../stores/settings'

const selectionStore = useSelectionStore()
const settingsStore = useSettingsStore()

const commandName = ref('选择一个命令')
const commandDescription = ref('')
const commandContent = ref('')

// 监听选择变化
watch(() => selectionStore.selectedItem, (item) => {
  if (item?.type === 'command' && item.data) {
    commandName.value = item.data.name
    commandDescription.value = item.data.description
    commandContent.value = item.data.content || '暂无命令内容'
  }
}, { immediate: true })

// 复制命令
const copyCommand = async () => {
  try {
    await navigator.clipboard.writeText(commandContent.value)
    alert('命令已复制到剪贴板！')
  } catch {
    // 备用方案
    const textarea = document.createElement('textarea')
    textarea.value = commandContent.value
    document.body.appendChild(textarea)
    textarea.select()
    document.execCommand('copy')
    document.body.removeChild(textarea)
    alert('命令已复制到剪贴板！')
  }
}

const hasContent = computed(() => commandContent.value && commandContent.value !== '暂无命令内容')
</script>

<template>
  <div class="commands-view h-100 overflow-auto">
    <v-card class="ma-4" elevation="0">
      <v-card-title class="pb-2">
        {{ commandName }}
      </v-card-title>

      <v-divider />

      <v-card-text>
        <v-alert
          v-if="commandDescription"
          type="info"
          variant="tonal"
          class="mb-4"
        >
          <strong>描述：</strong> {{ commandDescription }}
        </v-alert>

        <v-card v-if="hasContent" class="pa-4" color="surface-variant" elevation="0">
          <div class="d-flex align-center mb-2">
            <span class="text-overline text-secondary">命令内容</span>
            <v-spacer />
            <v-btn
              size="small"
              variant="outlined"
              color="primary"
              @click="copyCommand"
            >
              复制命令
            </v-btn>
          </div>
          <pre :style="{ fontFamily: 'Consolas, Monaco, monospace', fontSize: settingsStore.fontSize + 'px', lineHeight: 1.6, color: 'rgb(var(--v-theme-on-surface))', whiteSpace: 'pre-wrap' }">{{ commandContent }}</pre>
        </v-card>

        <v-card v-else class="pa-8 text-center" color="surface-variant" elevation="0">
          <p class="text-body-1 text-secondary">
            请从右侧列表选择一个命令查看详情
          </p>
        </v-card>

        <v-card class="mt-4 pa-4" color="surface" elevation="0">
          <span class="text-caption text-secondary">
            提示：按 <kbd>~</kbd> 键快速唤醒命令库，输入命令名称快速搜索
          </span>
        </v-card>
      </v-card-text>
    </v-card>
  </div>
</template>

<style scoped>
pre {
  margin: 0;
}

kbd {
  background: rgb(var(--v-theme-background));
  padding: 2px 8px;
  border-radius: 4px;
  border: 1px solid rgb(var(--v-theme-border));
  font-family: monospace;
  font-size: 12px;
}
</style>
