<script lang="ts" setup>
import { ref } from 'vue'

const settings = ref({
  startupWithSystem: true,
  minimizeToTray: true,
  commandHotkey: '~',
  commandScope: 'internal',
  defaultBrowser: 'internal',
  theme: 'light',
  fontSize: 14
})
</script>

<template>
  <div class="settings-view h-100">
    <v-container class="pa-4" style="max-width: 800px;">
      <v-card elevation="0">
        <v-card-title class="text-h5 pb-2">
          <v-icon icon="mdi-cog" color="primary" class="mr-2" />
          设置
        </v-card-title>
        
        <v-divider class="mb-4" />
        
        <!-- 常规设置 -->
        <v-card-text class="pa-0">
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
          
          <v-divider class="my-4" />
          
          <!-- 快捷键设置 -->
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
          <h3 class="text-subtitle-1 mb-3">外观</h3>
          
          <v-list nav density="compact" class="pa-0">
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
          
          <v-divider class="my-4" />
          
          <!-- 操作按钮 -->
          <div class="d-flex justify-end">
            <v-btn variant="outlined" class="mr-2">重置</v-btn>
            <v-btn color="primary" variant="flat">保存设置</v-btn>
          </div>
        </v-card-text>
      </v-card>
    </v-container>
  </div>
</template>
