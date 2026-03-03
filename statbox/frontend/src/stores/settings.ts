import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const fontSize = ref(14)
  const theme = ref('statboxLight')
  const startupWithSystem = ref(true)
  const minimizeToTray = ref(true)

  // 默认设置
  const defaultSettings = {
    fontSize: 14,
    theme: 'statboxLight',
    startupWithSystem: true,
    minimizeToTray: true
  }

  // 从本地存储加载设置
  const loadSettings = () => {
    try {
      const saved = localStorage.getItem('statbox-settings')
      if (saved) {
        const data = JSON.parse(saved)
        fontSize.value = data.fontSize ?? defaultSettings.fontSize
        theme.value = data.theme ?? defaultSettings.theme
        startupWithSystem.value = data.startupWithSystem ?? defaultSettings.startupWithSystem
        minimizeToTray.value = data.minimizeToTray ?? defaultSettings.minimizeToTray
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
        minimizeToTray: minimizeToTray.value
      }
      localStorage.setItem('statbox-settings', JSON.stringify(data))
      console.log('设置已保存')
    } catch (error) {
      console.error('保存设置失败:', error)
    }
  }

  // 重置设置为默认值
  const resetSettings = () => {
    fontSize.value = defaultSettings.fontSize
    theme.value = defaultSettings.theme
    startupWithSystem.value = defaultSettings.startupWithSystem
    minimizeToTray.value = defaultSettings.minimizeToTray
    saveSettings()
    console.log('设置已重置为默认值')
  }

  // 监听变化自动保存
  watch([fontSize, theme, startupWithSystem, minimizeToTray], () => {
    saveSettings()
  })

  return {
    fontSize,
    theme,
    startupWithSystem,
    minimizeToTray,
    loadSettings,
    saveSettings,
    resetSettings
  }
})
