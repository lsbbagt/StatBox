<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import { useTheme } from 'vuetify'
import { themeOptions } from '../plugins/vuetify'
import { useSelectionStore } from '../stores/selection'
import { useSettingsStore } from '../stores/settings'
import { GetConfigDir, GetTemplatesDir, GetDataDir, SetDataDir, OpenFolderInExplorer, OpenDirectoryDialog } from '../../wailsjs/go/main/App'

const theme = useTheme()
const selectionStore = useSelectionStore()
const settingsStore = useSettingsStore()

const configDir = ref('')
const templatesDir = ref('')
const dataDir = ref('')
const showDataDirDialog = ref(false)
const newDataDir = ref('')
const isMigrating = ref(false)

// 加载路径配置
const loadPaths = async () => {
  try {
    configDir.value = await GetConfigDir()
    templatesDir.value = await GetTemplatesDir()
    dataDir.value = await GetDataDir()
  } catch (error) {
    console.error('加载路径配置失败:', error)
  }
}

// 打开文件夹
const openFolder = async (path: string) => {
  try {
    await OpenFolderInExplorer(path)
  } catch (error) {
    console.error('打开文件夹失败:', error)
    alert('打开文件夹失败: ' + error)
  }
}

// 选择新数据目录
const selectDataDir = async () => {
  try {
    const selectedPath = await OpenDirectoryDialog('选择数据存储位置')
    if (selectedPath) {
      newDataDir.value = selectedPath
      showDataDirDialog.value = true
    }
  } catch (error) {
    console.error('选择目录失败:', error)
  }
}

// 确认更改数据目录
const confirmChange = async () => {
  if (!newDataDir.value) return
  
  isMigrating.value = true
  try {
    await SetDataDir(newDataDir.value)
    showDataDirDialog.value = false
    alert('数据目录已更改！应用将重新加载以使用新的数据目录。')
    // 重新加载页面以应用新的数据目录
    window.location.reload()
  } catch (error) {
    console.error('更改数据目录失败:', error)
    alert('更改数据目录失败: ' + error)
  } finally {
    isMigrating.value = false
  }
}

// 取消更改
const cancelChange = () => {
  showDataDirDialog.value = false
  newDataDir.value = ''
}

// 监听选择变化，滚动到对应设置区域
watch(() => selectionStore.selectedItem, (item) => {
  if (item?.type === 'setting' && item.data) {
    const element = document.getElementById(item.data.id)
    if (element) {
      element.scrollIntoView({ behavior: 'smooth', block: 'start' })
    }
  }
})

// 监听主题变化
watch(() => settingsStore.theme, (newTheme) => {
  theme.global.name.value = newTheme
})

// 组件挂载时加载路径和设置
onMounted(() => {
  loadPaths()
  settingsStore.loadSettings()
})

// 重置设置
const handleReset = () => {
  if (confirm('确定要重置所有设置为默认值吗？')) {
    settingsStore.resetSettings()
    // 更新主题
    theme.global.name.value = settingsStore.theme
  }
}

// 保存设置
const handleSave = () => {
  settingsStore.saveSettings()
  alert('设置已保存')
}
</script>

<template>
  <div class="settings-view">
    <div class="settings-content">
      <div class="settings-header">
        <v-icon icon="mdi-cog" color="primary" class="mr-2" />
        <span class="text-h5">设置</span>
      </div>

      <v-divider class="mb-4" />

      <!-- 常规设置 -->
      <div id="general" class="settings-section">
        <h3 class="text-subtitle-1 mb-3">常规设置</h3>

        <v-list nav density="compact" class="pa-0">
          <v-list-item class="px-0">
            <v-list-item-title>开机自启动</v-list-item-title>
            <v-list-item-subtitle>系统启动时自动运行 StatBox（静默启动）</v-list-item-subtitle>
            <template #append>
              <v-switch
                v-model="settingsStore.startupWithSystem"
                color="primary"
                hide-details
              />
            </template>
          </v-list-item>

          <v-list-item class="px-0">
            <v-list-item-title>关闭后最小化到托盘</v-list-item-title>
            <v-list-item-subtitle>关闭窗口时最小化到系统托盘而不是退出</v-list-item-subtitle>
            <template #append>
              <v-switch
                v-model="settingsStore.minimizeToTray"
                color="primary"
                hide-details
              />
            </template>
          </v-list-item>
        </v-list>
      </div>

      <v-divider class="my-4" />

      <!-- 数据存储位置 -->
      <div id="storage" class="settings-section">
        <h3 class="text-subtitle-1 mb-3">数据存储位置</h3>

        <v-list nav density="compact" class="pa-0">
          <v-list-item class="px-0">
            <v-list-item-title>数据目录</v-list-item-title>
            <v-list-item-subtitle>存储所有应用数据（配置、模板、收藏夹等）</v-list-item-subtitle>
            <template #append>
              <v-btn
                size="small"
                variant="outlined"
                prepend-icon="mdi-folder-open"
                @click="openFolder(dataDir)"
                :disabled="!dataDir"
                class="mr-2"
              >
                打开
              </v-btn>
              <v-btn
                size="small"
                variant="outlined"
                prepend-icon="mdi-folder-edit"
                @click="selectDataDir"
              >
                修改
              </v-btn>
            </template>
          </v-list-item>
          <div class="px-4 pb-2">
            <v-chip size="small" color="surface-variant">{{ dataDir || '加载中...' }}</v-chip>
          </div>
        </v-list>
      </div>

      <v-divider class="my-4" />

      <!-- 外观设置 -->
      <div id="appearance" class="settings-section">
        <h3 class="text-subtitle-1 mb-3">外观</h3>

        <v-list nav density="compact" class="pa-0">
          <!-- 颜色主题 -->
          <v-list-item class="px-0">
            <v-list-item-title>颜色主题</v-list-item-title>
            <v-list-item-subtitle>选择你喜欢的主题颜色</v-list-item-subtitle>
          </v-list-item>

          <div class="d-flex flex-wrap ga-2 mb-4 px-1">
            <v-btn
              v-for="option in themeOptions"
              :key="option.value"
              :style="{
                border: settingsStore.theme === option.value ? '3px solid ' + option.color : '2px solid #E1E8ED',
                backgroundColor: settingsStore.theme === option.value ? option.color + '20' : 'transparent'
              }"
              variant="outlined"
              size="small"
              class="text-none"
              @click="settingsStore.theme = option.value"
            >
              <v-icon :color="option.color" class="mr-1">mdi-palette</v-icon>
              {{ option.title }}
            </v-btn>
          </div>

          <v-list-item class="px-0">
            <v-list-item-title>字体大小</v-list-item-title>
            <v-list-item-subtitle>编辑器和终端的字体大小</v-list-item-subtitle>
            <template #append>
              <v-slider
                v-model="settingsStore.fontSize"
                :min="10"
                :max="20"
                :step="1"
                thumb-label
                color="primary"
                style="width: 200px;"
                hide-details
              />
            </template>
          </v-list-item>
        </v-list>
      </div>

      <v-divider class="my-4" />

      <!-- 操作按钮 -->
      <div class="d-flex justify-end pb-4">
        <v-btn variant="outlined" class="mr-2" @click="handleReset">重置</v-btn>
        <v-btn color="primary" variant="flat" @click="handleSave">保存设置</v-btn>
      </div>
    </div>

    <!-- 确认更改数据目录对话框 -->
    <v-dialog v-model="showDataDirDialog" max-width="500" persistent>
      <v-card>
        <v-card-title class="text-h6">
          <v-icon icon="mdi-folder-edit" class="mr-2" />
          确认更改数据目录
        </v-card-title>
        <v-card-text>
          <div class="mb-4">
            <div class="text-subtitle-2 mb-2">当前数据目录：</div>
            <v-chip size="small" color="surface-variant">{{ dataDir }}</v-chip>
          </div>
          <div class="mb-4">
            <div class="text-subtitle-2 mb-2">新数据目录：</div>
            <v-chip size="small" color="primary">{{ newDataDir }}</v-chip>
          </div>
          <v-alert type="info" variant="tonal" density="compact">
            <div class="text-body-2">
              应用将使用新目录中的数据。如果新目录中没有模板数据，将创建空的模板目录。
            </div>
          </v-alert>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            variant="outlined"
            @click="cancelChange"
            :disabled="isMigrating"
          >
            取消
          </v-btn>
          <v-btn
            color="primary"
            variant="flat"
            @click="confirmChange"
            :loading="isMigrating"
          >
            确认更改
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<style scoped>
.settings-view {
  height: 100%;
  width: 100%;
}

.settings-content {
  max-width: 800px;
  padding: 16px;
  margin: 0 auto;
}

.settings-header {
  display: flex;
  align-items: center;
  padding-bottom: 8px;
}

.settings-section {
  margin-bottom: 8px;
}

/* 隐藏v-list内部滚动条 */
.settings-view :deep(.v-list) {
  overflow: visible !important;
}

.settings-view :deep(.v-list::-webkit-scrollbar) {
  display: none;
}
</style>