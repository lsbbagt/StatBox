<script lang="ts" setup>
import { ref } from 'vue'
import type { ModuleType, ModuleItem } from '../types'

const props = defineProps<{
  modelValue: boolean
  activeModule: ModuleType
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'select', module: ModuleType): void
}>()

const collapsed = ref(false)

const modules: ModuleItem[] = [
  { id: 'bookmarks', name: '收藏夹', icon: 'mdi-folder-star' },
  { id: 'templates', name: '代码模板', icon: 'mdi-code-tags' },
  { id: 'commands', name: '命令库', icon: 'mdi-keyboard' },
  { id: 'settings', name: '设置', icon: 'mdi-cog' }
]

const selectModule = (module: ModuleType) => {
  console.log('Selecting module:', module)
  emit('select', module)
}

const toggleCollapse = () => {
  collapsed.value = !collapsed.value
}

const hideSidebar = () => {
  emit('update:modelValue', false)
}
</script>

<template>
  <div v-if="modelValue" class="sidebar" :style="{ width: collapsed ? '60px' : '200px' }">
    <!-- 模块列表 -->
    <div class="module-list">
      <div
        v-for="module in modules"
        :key="module.id"
        :class="['module-item', { active: activeModule === module.id }]"
        @click="selectModule(module.id)"
      >
        <span :class="['mdi', module.icon, 'module-icon']"></span>
        <span v-if="!collapsed" class="module-name">{{ module.name }}</span>
      </div>
    </div>
    
    <!-- 底部按钮 -->
    <div class="sidebar-footer">
      <button class="footer-btn" @click="toggleCollapse">
        <span :class="['mdi', collapsed ? 'mdi-chevron-right' : 'mdi-chevron-left']"></span>
        <span v-if="!collapsed">收起</span>
      </button>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  height: 100%;
  background-color: #EEF2F6;
  border-right: 1px solid #E1E8ED;
  display: flex;
  flex-direction: column;
  transition: width 0.2s ease;
}

.module-list {
  flex: 1;
  padding: 8px;
  overflow-y: auto;
}

.module-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  margin: 4px 0;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  color: #2C3E50;
}

.module-item:hover {
  background-color: rgba(91, 155, 213, 0.1);
}

.module-item.active {
  background-color: rgba(91, 155, 213, 0.15);
  color: #5B9BD5;
}

.module-icon {
  font-size: 20px;
  flex-shrink: 0;
}

.module-name {
  font-size: 14px;
  font-weight: 500;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.sidebar-footer {
  border-top: 1px solid #E1E8ED;
  padding: 8px;
}

.footer-btn {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 10px;
  border: none;
  background: transparent;
  border-radius: 8px;
  cursor: pointer;
  color: #7A8B99;
  font-size: 14px;
  transition: all 0.2s ease;
}

.footer-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
  color: #5B9BD5;
}
</style>
