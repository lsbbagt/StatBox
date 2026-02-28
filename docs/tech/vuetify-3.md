# Vuetify 3 技术文档

> 本文档用于 StatBox 项目前端开发参考

---

## 📋 版本信息

- **版本号：** ^3.5.0
- **支持框架：** Vue 3
- **设计系统：** Material Design

---

## ✨ 主要特性

### 1. 组件库
- 80+ 预制组件
- 响应式设计
- 无障碍支持
- RTL支持

### 2. 主题系统
- 亮色/暗色主题
- 自定义颜色
- 动态主题切换

### 3. 图标支持
- Material Design Icons
- Font Awesome
- 自定义SVG图标

### 4. 布局系统
- Grid系统
- Flex布局
- 响应式断点

---

## 🎯 StatBox 中的关键应用

### 1. 主题配置

```typescript
// frontend/src/plugins/vuetify.ts
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
    themes: {
      dark: {
        colors: {
          primary: '#2196F3',    // 主色调
          secondary: '#424242',   // 次要颜色
          accent: '#FF4081',      // 强调色
          surface: '#1e1e1e',     // 表面颜色
          background: '#121212',  // 背景颜色
        }
      },
      light: {
        colors: {
          primary: '#1976D2',
          secondary: '#757575',
          accent: '#E91E63',
          surface: '#FFFFFF',
          background: '#FAFAFA',
        }
      }
    }
  }
})
```

### 2. 布局组件

**应用布局：**

```vue
<template>
  <v-app>
    <TitleBar />
    <v-main class="main-content">
      <div class="d-flex h-100">
        <Sidebar />
        <div class="flex-grow-1 content-wrapper">
          <router-view />
        </div>
        <RightPanel />
      </div>
    </v-main>
  </v-app>
</template>
```

**导航抽屉：**

```vue
<v-navigation-drawer
  v-model="visible"
  width="200"
  :rail="collapsed"
  rail-width="60"
  permanent
>
  <v-list nav density="compact">
    <v-list-item
      v-for="module in modules"
      :key="module.id"
      :prepend-icon="module.icon"
      :title="module.name"
    />
  </v-list>
</v-navigation-drawer>
```

### 3. 树形组件

```vue
<v-treeview
  :items="bookmarks"
  item-title="name"
  item-value="id"
  activatable
  @update:active="onSelect"
>
  <template #prepend="{ item }">
    <v-icon :icon="item.icon || 'mdi-folder'" />
  </template>
</v-treeview>
```

### 4. 搜索输入框

```vue
<v-text-field
  v-model="searchQuery"
  density="compact"
  variant="outlined"
  placeholder="搜索..."
  prepend-inner-icon="mdi-magnify"
  hide-details
  clearable
  @update:model-value="onSearch"
/>
```

### 5. 右键菜单

```vue
<v-menu>
  <template #activator="{ props }">
    <div v-bind="props" @contextmenu.prevent="showMenu">
      右键点击我
    </div>
  </template>
  <v-list>
    <v-list-item title="查看" @click="onView" />
    <v-list-item title="编辑" @click="onEdit" />
    <v-list-item title="删除" @click="onDelete" />
  </v-list>
</v-menu>
```

---

## 🔧 最佳实践

### 1. 组件组合

**使用 Composition API：**

```vue
<script setup lang="ts">
import { ref, computed } from 'vue'
import type { ModuleType } from '../types'

const props = defineProps<{
  modelValue: boolean
  activeModule: ModuleType
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'select', module: ModuleType): void
}>()

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})
</script>
```

### 2. 响应式设计

**断点系统：**

```vue
<template>
  <v-container>
    <v-row>
      <v-col cols="12" md="6" lg="4">
        <!-- 小屏幕12列，中屏幕6列，大屏幕4列 -->
      </v-col>
    </v-row>
  </v-container>
</template>
```

**使用 useDisplay：**

```typescript
import { useDisplay } from 'vuetify'

const { mobile, mdAndUp } = useDisplay()

if (mobile.value) {
  // 移动端逻辑
}
```

### 3. 状态管理

**使用 Pinia：**

```typescript
// stores/bookmarks.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useBookmarksStore = defineStore('bookmarks', () => {
  const bookmarks = ref([])
  const searchQuery = ref('')
  
  const filteredBookmarks = computed(() => {
    return bookmarks.value.filter(b => 
      b.name.includes(searchQuery.value)
    )
  })
  
  return { bookmarks, searchQuery, filteredBookmarks }
})
```

### 4. 图标使用

```vue
<!-- Material Design Icons -->
<v-icon icon="mdi-folder" />
<v-icon icon="mdi-code-tags" />

<!-- 在按钮中 -->
<v-btn prepend-icon="mdi-plus">新建</v-btn>

<!-- 在列表项中 -->
<v-list-item prepend-icon="mdi-folder" title="收藏夹" />
```

---

## 🎨 主题定制

### CSS 变量

```css
:root {
  --v-theme-primary: #2196F3;
  --v-theme-secondary: #424242;
  --v-theme-background: #121212;
}
```

### 动态主题切换

```typescript
import { useTheme } from 'vuetify'

const theme = useTheme()

// 切换主题
theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'

// 自定义主题颜色
theme.themes.value.dark.colors.primary = '#FF5722'
```

---

## 📚 常用组件速查

### 布局组件

| 组件 | 用途 |
|------|------|
| `v-app` | 应用根容器 |
| `v-main` | 主内容区域 |
| `v-navigation-drawer` | 导航抽屉 |
| `v-app-bar` | 应用栏 |
| `v-container` | 容器 |
| `v-row` / `v-col` | 栅格系统 |

### 交互组件

| 组件 | 用途 |
|------|------|
| `v-btn` | 按钮 |
| `v-text-field` | 文本输入框 |
| `v-textarea` | 多行文本框 |
| `v-select` | 下拉选择框 |
| `v-checkbox` | 复选框 |
| `v-switch` | 开关 |

### 数据展示

| 组件 | 用途 |
|------|------|
| `v-list` | 列表 |
| `v-treeview` | 树形视图 |
| `v-table` | 表格 |
| `v-card` | 卡片 |
| `v-chip` | 标签 |

### 反馈组件

| 组件 | 用途 |
|------|------|
| `v-dialog` | 对话框 |
| `v-snackbar` | 消息提示 |
| `v-menu` | 菜单 |
| `v-tooltip` | 工具提示 |
| `v-progress-circular` | 圆形进度条 |

---

## ⚠️ 注意事项

### 1. 样式隔离

使用 `scoped` 样式避免全局污染：

```vue
<style scoped>
.my-component {
  /* 局部样式 */
}
</style>
```

### 2. 深度选择器

修改子组件样式：

```css
:deep(.v-list-item) {
  margin: 4px 8px;
}
```

### 3. 滚动条隐藏

```css
.no-scrollbar {
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.no-scrollbar::-webkit-scrollbar {
  display: none;
}
```

---

## 🔗 相关资源

- [Vuetify 官方文档](https://vuetifyjs.com/en/)
- [Material Design 规范](https://material.io/design)
- [Material Design Icons](https://materialdesignicons.com/)
- [Vuetify 示例](https://vuetifyjs.com/en/getting-started/wireframes/)

---

## 📅 更新记录

- 2025-02-28: 创建文档，用于 StatBox 项目
