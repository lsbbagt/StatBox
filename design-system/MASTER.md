# StatBox 设计系统

> 温暖亮色主题 - 专业统计学家工具箱

---

## 🎨 设计原则

### 核心理念
- **简洁**：清晰的信息层级，无冗余装饰
- **温暖**：柔和的色彩，舒适的视觉体验
- **专业**：符合学术工具的专业气质
- **高效**：快速访问，流畅交互

---

## 🌈 色彩系统

### 主色调（Primary）- 温暖蓝
```
Primary: #5B9BD5    // 温暖的天蓝色，专业且亲切
Primary Light: #8DB8E8
Primary Dark: #3A7BC8
```

**使用场景：**
- 主要操作按钮
- 导航栏高亮
- 重要图标
- 链接文本

### 辅助色（Secondary）- 柔和灰蓝
```
Secondary: #7A8B99    // 中性灰蓝，平衡主色
Secondary Light: #9BA8B3
Secondary Dark: #5A6B79
```

**使用场景：**
- 次要按钮
- 标签背景
- 分隔线

### 强调色（Accent）- 温暖橙
```
Accent: #F4A261      // 温暖橙色，增添活力
Accent Light: #F7BA8A
Accent Dark: #E08A3E
```

**使用场景：**
- 重要提示
- 选中状态
- 警告信息

### 背景色系统（Background）

#### 表面颜色
```
Surface: #FFFFFF      // 主内容区域背景
Surface Variant: #F5F7FA  // 次要表面
Surface Elevated: #FFFFFF  // 提升元素（卡片、模态框）
```

#### 背景层级
```
Background: #F8FAFB      // 页面背景
Background Secondary: #EEF2F6  // 侧边栏背景
Background Tertiary: #E8EDF3   // 第三层级背景
```

### 文字颜色（Text）
```
Primary Text: #2C3E50      // 主要文字（深灰蓝）
Secondary Text: #5A6B79    // 次要文字
Muted Text: #8795A1        // 弱化文字
Disabled Text: #A8B5C1     // 禁用文字
```

### 边框与分隔
```
Border: #E1E8ED
Border Light: #EEF2F6
Divider: #D8E0E7
```

### 状态颜色
```
Success: #4CAF50    // 成功 - 柔和绿
Warning: #FF9800    // 警告 - 温暖橙
Error: #EF5350      // 错误 - 柔和红
Info: #42A5F5       // 信息 - 明亮蓝
```

---

## 📝 字体系统

### 字体家族

**主要字体：Inter**
- 现代无衬线字体
- 优秀的可读性
- 适合数据和代码展示

**次要字体：Source Sans Pro**
- 清晰易读
- 适合长文本

**代码字体：JetBrains Mono**
- 等宽字体
- 优秀的代码可读性

### 字体大小

```
Display: 32px / 40px line-height    // 大标题
H1: 24px / 32px line-height         // 一级标题
H2: 20px / 28px line-height         // 二级标题
H3: 18px / 26px line-height         // 三级标题
Body: 14px / 22px line-height       // 正文
Body Small: 12px / 18px line-height // 小号正文
Caption: 11px / 16px line-height    // 说明文字
```

### 字重

```
Light: 300
Regular: 400
Medium: 500
Semi-Bold: 600
Bold: 700
```

---

## 📐 间距系统

### 基础单位：4px

```
Space 1: 4px
Space 2: 8px
Space 3: 12px
Space 4: 16px
Space 5: 20px
Space 6: 24px
Space 8: 32px
Space 10: 40px
Space 12: 48px
Space 16: 64px
```

### 组件内边距

```
Button Padding: 12px 20px
Card Padding: 20px
List Item Padding: 12px 16px
Input Padding: 10px 14px
```

---

## 🎭 圆角系统

```
None: 0px
Small: 4px     // 小组件（标签、徽章）
Medium: 8px    // 中等组件（按钮、输入框）
Large: 12px    // 大组件（卡片）
XLarge: 16px   // 特大组件（模态框）
Full: 9999px   // 圆形
```

---

## 🌟 阴影系统

### 层级阴影

```
Elevation 1: 0 1px 3px rgba(0, 0, 0, 0.06)
Elevation 2: 0 2px 6px rgba(0, 0, 0, 0.08)
Elevation 3: 0 4px 12px rgba(0, 0, 0, 0.10)
Elevation 4: 0 8px 24px rgba(0, 0, 0, 0.12)
```

### 特殊效果

```
Card Shadow: 0 2px 8px rgba(91, 155, 213, 0.08)
Button Shadow: 0 2px 4px rgba(91, 155, 213, 0.12)
Modal Shadow: 0 8px 32px rgba(0, 0, 0, 0.15)
```

---

## 🎬 动画系统

### 时长

```
Instant: 0ms
Fast: 150ms
Normal: 250ms
Slow: 350ms
Very Slow: 500ms
```

### 缓动函数

```
Ease In: cubic-bezier(0.4, 0, 1, 1)
Ease Out: cubic-bezier(0, 0, 0.2, 1)
Ease In Out: cubic-bezier(0.4, 0, 0.2, 1)
Standard: cubic-bezier(0.4, 0.0, 0.2, 1)
```

### 常用动画

**淡入淡出**
```css
transition: opacity 250ms ease-in-out
```

**滑动**
```css
transition: transform 250ms ease-out
```

**缩放**
```css
transition: transform 200ms ease-in-out
```

---

## 🧩 组件规范

### 按钮

**主要按钮**
- 背景：Primary (#5B9BD5)
- 文字：白色
- 圆角：8px
- 内边距：12px 20px
- Hover：背景加深 10%
- 点击：scale(0.98)

**次要按钮**
- 背景：透明
- 边框：Border (#E1E8ED)
- 文字：Primary Text
- Hover：背景 Surface Variant

**文字按钮**
- 背景：透明
- 文字：Primary
- Hover：背景 Surface Variant

### 输入框

- 背景：白色
- 边框：Border (#E1E8ED)
- 圆角：8px
- 内边距：10px 14px
- Focus：边框变为主色，添加柔和阴影

### 卡片

- 背景：白色
- 圆角：12px
- 阴影：Elevation 1
- 内边距：20px
- Hover：阴影提升至 Elevation 2

### 列表项

- 高度：48px
- 内边距：12px 16px
- Hover：背景 Surface Variant
- 选中：背景 Primary Light，左侧 3px 主色边框

---

## 🎯 特殊场景

### 代码编辑器

- 背景：#FAFBFC
- 行号：#9BA8B3
- 当前行：#F0F4F8
- 选中：rgba(91, 155, 213, 0.15)

### 浏览器标签

- 激活：白色，底部主色边框
- 未激活：Surface Variant
- Hover：背景变亮

### 搜索框

- 背景：白色
- 边框：Border
- Focus：边框变为主色，柔和阴影
- 图标：Secondary Text

---

## ✅ 设计检查清单

### 颜色对比度
- [ ] 正文文字对比度 ≥ 4.5:1
- [ ] 大标题对比度 ≥ 3:1
- [ ] 状态颜色在白色背景上可辨识

### 交互反馈
- [ ] 所有可点击元素有 cursor: pointer
- [ ] Hover 状态有视觉反馈
- [ ] 过渡动画流畅（150-300ms）

### 布局一致性
- [ ] 使用统一的间距系统
- [ ] 组件圆角统一
- [ ] 阴影层级清晰

### 无障碍设计
- [ ] 表单元素有 label
- [ ] 图标按钮有 aria-label
- [ ] Focus 状态可见
- [ ] 键盘导航支持

---

## 📦 实施要点

### Vuetify 主题配置

```typescript
const statboxLightTheme = {
  dark: false,
  colors: {
    primary: '#5B9BD5',
    secondary: '#7A8B99',
    accent: '#F4A261',
    surface: '#FFFFFF',
    background: '#F8FAFB',
    error: '#EF5350',
    info: '#42A5F5',
    success: '#4CAF50',
    warning: '#FF9800',
  }
}
```

### CSS 变量

```css
:root {
  /* Colors */
  --color-primary: #5B9BD5;
  --color-secondary: #7A8B99;
  --color-accent: #F4A261;
  
  /* Backgrounds */
  --bg-surface: #FFFFFF;
  --bg-background: #F8FAFB;
  --bg-elevated: #FFFFFF;
  
  /* Text */
  --text-primary: #2C3E50;
  --text-secondary: #5A6B79;
  --text-muted: #8795A1;
  
  /* Borders */
  --border-color: #E1E8ED;
  --border-light: #EEF2F6;
  
  /* Shadows */
  --shadow-sm: 0 1px 3px rgba(0, 0, 0, 0.06);
  --shadow-md: 0 2px 6px rgba(0, 0, 0, 0.08);
  --shadow-lg: 0 4px 12px rgba(0, 0, 0, 0.10);
}
```

---

## 🚫 反模式

### 避免使用

❌ 纯黑色文字（#000000）- 过于生硬  
❌ 高饱和度颜色 - 刺眼  
❌ 过多圆角 - 不专业  
❌ 过重阴影 - 视觉干扰  
❌ Emoji 作为图标 - 不专业  

### 推荐替代

✅ 深灰蓝文字（#2C3E50）- 柔和专业  
✅ 低饱和度色彩 - 舒适耐看  
✅ 统一圆角系统 - 协调一致  
✅ 轻柔阴影 - 清晰层级  
✅ SVG 图标 - 专业清晰  

---

## 📚 参考资源

- [Material Design Color System](https://material.io/design/color)
- [Refactoring UI](https://www.refactoringui.com/)
- [Vuetify Theme Configuration](https://vuetifyjs.com/en/features/theme/)

---

**设计系统创建日期：** 2025-02-28  
**适用项目：** StatBox - 统计学家工具箱
