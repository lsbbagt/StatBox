<script lang="ts" setup>
import { ref, computed } from 'vue'
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

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const modules: ModuleItem[] = [
  { id: 'bookmarks', name: '收藏夹', icon: 'mdi-folder-star' },
  { id: 'templates', name: '代码模板', icon: 'mdi-code-tags' },
  { id: 'commands', name: '命令库', icon: 'mdi-keyboard' },
  { id: 'settings', name: '设置', icon: 'mdi-cog' }
]

const selectModule = (module: ModuleType) => {
  emit('select', module)
}

const toggleCollapse = () => {
  collapsed.value = !collapsed.value
}
</script>

<template>
  <v-navigation-drawer
    v-model="visible"
    width="200"
    class="sidebar"
    :rail="collapsed"
    rail-width="60"
    permanent
    color="background-secondary"
  >
    <v-list nav density="compact" class="py-2">
      <v-list-item
        v-for="module in modules"
        :key="module.id"
        :value="module.id"
        :active="activeModule === module.id"
        :prepend-icon="module.icon"
        :title="collapsed ? '' : module.name"
        @click="selectModule(module.id)"
        class="module-item my-1"
        :color="activeModule === module.id ? 'primary' : 'default'"
        rounded="lg"
      />
    </v-list>
    
    <template #append>
      <v-divider />
      <v-btn
        block
        variant="text"
        :icon="collapsed ? 'mdi-chevron-right' : 'mdi-chevron-left'"
        @click="toggleCollapse"
        class="collapse-btn"
        color="secondary"
      />
    </template>
  </v-navigation-drawer>
</template>

<style scoped>
.sidebar {
  border-right: 1px solid rgb(var(--v-theme-border));
}

.module-item {
  margin: 4px 8px;
}
</style>
