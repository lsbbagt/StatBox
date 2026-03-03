<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import { useTheme } from 'vuetify'
import { themeOptions } from '../plugins/vuetify'
import { useSelectionStore } from '../stores/selection'
import { GetConfigDir, GetTemplatesDir, OpenFolderInExplorer } from '../../wailsjs/go/main/App'

const theme = useTheme()
const selectionStore = useSelectionStore()

const settings = ref({
  startupWithSystem: true,
  minimizeToTray: true,
  silentStartup: true,
  commandHotkey: '~',
  commandScope: 'internal',
  defaultBrowser: 'internal',
  theme: 'statboxLight',
  fontSize: 14
})

const configDir = ref('')
const templatesDir = ref('')

// 加载路径配置
const loadPaths = async () => {
  try {
    configDir.value = await GetConfigDir()
    templatesDir.value = await GetTemplatesDir()
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
watch(() => settings.value.theme, (newTheme) => {
  theme.global.name.value = newTheme
})

// 组件挂载时加载路径
onMounted(() => {
  loadPaths()
})
</script>

<template>
  <div class="settings-view h-100 overflow-auto">
    <v-container class="pa-4" style="max-width: 800px;">
      <v-card elevation="0">
        <v-card-title class="text-h5 pb-2">
          <v-icon icon="mdi-cog" color="primary" class="mr-2" />
          设置
        </v-card-title>

        <v-divider class="mb-4" />

        <!-- 常规设置 -->
        <v-card-text class="pa-0">
          <div id="general">
            <h3 class="text-subtitle-1 mb-3">常规设置</h3>

            <v-list nav density="compact" class="pa-0">
              <v-list-item class="px-0">
                <v-list-item-title>开机自启动</v-list-item-title>
                <v-list-item-subtitle>系统启动时自动运行 StatBox</v-list-item-subtitle>
                <template #append>
                  <v-switch
                    v-model="settings.startupWithSystem"
                    color="primary"
                    hide-details
                  />
                </template>
              </v-list-item>

              <v-list-item class="px-0">
                <v-list-item-title>静默启动</v-list-item-title>
                <v-list-item-subtitle>开机启动时最小化到托盘，不显示窗口</v-list-item-subtitle>
                <template #append>
                  <v-switch
                    v-model="settings.silentStartup"
                    color="primary"
                    hide-details
                  />
                </template>
              </v-list-item>

              <v-list-item class="px-0">
                <v-list-item-title>最小化到托盘</v-list-item-title>
                <v-list-item-subtitle>关闭窗口时最小化到系统托盘</v-list-item-subtitle>
                <template #append>
                  <v-switch
                    v-model="settings.minimizeToTray"
                    color="primary"
                    hide-details
                  />
                </template>
              </v-list-item>
            </v-list>
          </div>

          <v-divider class="my-4" />

          <!-- 路径设置 -->
          <div>
            <h3 class="text-subtitle-1 mb-3">数据存储位置</h3>

            <v-list nav density="compact" class="pa-0">
              <v-list-item class="px-0">
                <v-list-item-title>配置目录</v-list-item-title>
                <v-list-item-subtitle>存储应用配置和设置</v-list-item-subtitle>
                <template #append>
                  <v-btn
                    size="small"
                    variant="outlined"
                    prepend-icon="mdi-folder-open"
                    @click="openFolder(configDir)"
                    :disabled="!configDir"
                  >
                    打开
                  </v-btn>
                </template>
              </v-list-item>
              <div class="px-4 pb-2">
                <v-chip size="small" color="surface-variant">{{ configDir || '加载中...' }}</v-chip>
              </div>

              <v-list-item class="px-0">
                <v-list-item-title>模板目录</v-list-item-title>
                <v-list-item-subtitle>存储代码模板文件</v-list-item-subtitle>
                <template #append>
                  <v-btn
                    size="small"
                    variant="outlined"
                    prepend-icon="mdi-folder-open"
                    @click="openFolder(templatesDir)"
                    :disabled="!templatesDir"
                  >
                    打开
                  </v-btn>
                </template>
              </v-list-item>
              <div class="px-4 pb-2">
                <v-chip size="small" color="surface-variant">{{ templatesDir || '加载中...' }}</v-chip>
              </div>
            </v-list>
          </div>

          <v-divider class="my-4" />

          <!-- 快捷键设置 -->
          <div id="shortcuts">
            <h3 class="text-subtitle-1 mb-3">快捷键</h3>

            <v-list nav density="compact" class="pa-0">
              <v-list-item class="px-0">
                <v-list-item-title>命令库快捷键</v-list-item-title>
                <v-list-item-subtitle>按此键唤醒命令库</v-list-item-subtitle>
                <template #append>
                  <v-text-field
                    v-model="settings.commandHotkey"
                    density="compact"
                    variant="outlined"
                    style="width: 100px;"
                    hide-details
                  />
                </template>
              </v-list-item>

              <v-list-item class="px-0">
                <v-list-item-title>唤醒范围</v-list-item-title>
                <v-list-item-subtitle>命令库在何处可用</v-list-item-subtitle>
                <template #append>
                  <v-select
                    v-model="settings.commandScope"
                    :items="[
                      { title: '仅应用内', value: 'internal' },
                      { title: '全局', value: 'global' }
                    ]"
                    density="compact"
                    variant="outlined"
                    style="width: 150px;"
                    hide-details
                  />
                </template>
              </v-list-item>
            </v-list>
          </div>

          <v-divider class="my-4" />

          <!-- 浏览器设置 -->
          <h3 class="text-subtitle-1 mb-3">浏览器</h3>

          <v-list nav density="compact" class="pa-0">
            <v-list-item class="px-0">
              <v-list-item-title>默认浏览器</v-list-item-title>
              <v-list-item-subtitle>打开收藏时的默认行为</v-list-item-subtitle>
              <template #append>
                <v-select
                  v-model="settings.defaultBrowser"
                  :items="[
                    { title: '应用内浏览器', value: 'internal' },
                    { title: '系统默认浏览器', value: 'external' }
                  ]"
                  density="compact"
                  variant="outlined"
                  style="width: 180px;"
                  hide-details
                />
              </template>
            </v-list-item>
          </v-list>

          <v-divider class="my-4" />

          <!-- 外观设置 -->
          <div id="appearance">
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
                    border: settings.theme === option.value ? '3px solid ' + option.color : '2px solid #E1E8ED',
                    backgroundColor: settings.theme === option.value ? option.color + '20' : 'transparent'
                  }"
                  variant="outlined"
                  size="small"
                  class="text-none"
                  @click="settings.theme = option.value"
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
                    v-model="settings.fontSize"
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
            <v-btn variant="outlined" class="mr-2">重置</v-btn>
            <v-btn color="primary" variant="flat">保存设置</v-btn>
          </div>
        </v-card-text>
      </v-card>
    </v-container>
  </div>
</template>

<style scoped>
.settings-view {
  overflow-y: auto;
}
</style>
