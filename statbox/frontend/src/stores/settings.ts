import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const fontSize = ref(14)
  const theme = ref('statboxLight')
  const startupWithSystem = ref(true)
  const minimizeToTray = ref(true)
  const silentStartup = ref(true)
  const commandHotkey = ref('~')
  const commandScope = ref<'internal' | 'global'>('internal')
  const defaultBrowser = ref<'internal' | 'external'>('internal')

  // 从本地存储加载设置
  const loadSettings = () => {
    try {
      const saved = localStorage.getItem('statbox-settings')
      if (saved) {
        const data = JSON.parse(saved)
        fontSize.value = data.fontSize ?? 14
        theme.value = data.theme ?? 'statboxLight'
        startupWithSystem.value = data.startupWithSystem ?? true
        minimizeToTray.value = data.minimizeToTray ?? true
        silentStartup.value = data.silentStartup ?? true
        commandHotkey.value = data.commandHotkey ?? '~'
        commandScope.value = data.commandScope ?? 'internal'
        defaultBrowser.value = data.defaultBrowser ?? 'internal'
      }
    } catch (error) {
      console.error('加载设置失败:', error)
    }
  }

  // 保存设置到本地存储
  const saveSettings = () => {
    try {
      const data = {
        fontSize: fontSize.value,
        theme: theme.value,
        startupWithSystem: startupWithSystem.value,
        minimizeToTray: minimizeToTray.value,
        silentStartup: silentStartup.value,
        commandHotkey: commandHotkey.value,
        commandScope: commandScope.value,
        defaultBrowser: defaultBrowser.value
      }
      localStorage.setItem('statbox-settings', JSON.stringify(data))
    } catch (error) {
      console.error('保存设置失败:', error)
    }
  }

  // 监听变化自动保存
  watch([fontSize, theme, startupWithSystem, minimizeToTray, silentStartup, commandHotkey, commandScope, defaultBrowser], () => {
    saveSettings()
  })

  return {
    fontSize,
    theme,
    startupWithSystem,
    minimizeToTray,
    silentStartup,
    commandHotkey,
    commandScope,
    defaultBrowser,
    loadSettings,
    saveSettings
  }
})
