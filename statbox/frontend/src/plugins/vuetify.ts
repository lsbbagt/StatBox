import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

// 基础颜色配置
const baseColors = {
  // Surfaces
  surface: '#FFFFFF',
  'surface-variant': '#F5F7FA',
  background: '#F8FAFB',
  'background-secondary': '#EEF2F6',

  // Text
  'on-surface': '#2C3E50',
  'on-background': '#2C3E50',

  // Status Colors
  error: '#EF5350',
  info: '#42A5F5',
  success: '#4CAF50',
  warning: '#FF9800',

  // Borders
  border: '#E1E8ED',
  'border-light': '#EEF2F6',
  divider: '#D8E0E7',
}

// 预设主题颜色
const themeColors = {
  // 温暖蓝（默认）
  blue: {
    primary: '#5B9BD5',
    'primary-light': '#8DB8E8',
    'primary-dark': '#3A7BC8',
    secondary: '#7A8B99',
    accent: '#F4A261',
  },
  // 翡翠绿
  green: {
    primary: '#26A69A',
    'primary-light': '#64D8CB',
    'primary-dark': '#00796B',
    secondary: '#78909C',
    accent: '#FFA726',
  },
  // 薰衣草紫
  purple: {
    primary: '#7E57C2',
    'primary-light': '#B085F5',
    'primary-dark': '#4D2C91',
    secondary: '#8D6E63',
    accent: '#FF7043',
  },
  // 珊瑚橙
  orange: {
    primary: '#FF7043',
    'primary-light': '#FFAB91',
    'primary-dark': '#BF360C',
    secondary: '#90A4AE',
    accent: '#FFCA28',
  },
  // 玫瑰红
  rose: {
    primary: '#EC407A',
    'primary-light': '#F48FB1',
    'primary-dark': '#AD1457',
    secondary: '#9E9E9E',
    accent: '#FFEE58',
  },
}

// 创建主题
function createTheme(colorKey: string) {
  const colors = themeColors[colorKey as keyof typeof themeColors]
  return {
    dark: false,
    colors: {
      ...baseColors,
      ...colors,
      'secondary-light': colors.secondary,
      'secondary-dark': colors.secondary,
      'accent-light': colors.accent,
      'accent-dark': colors.accent,
    }
  }
}

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'statboxLight',
    themes: {
      statboxLight: createTheme('blue'),
      statboxGreen: createTheme('green'),
      statboxPurple: createTheme('purple'),
      statboxOrange: createTheme('orange'),
      statboxRose: createTheme('rose'),
    }
  },
  defaults: {
    VBtn: {
      style: 'text-transform: none;',
      rounded: 'md',
      fontWeight: 'medium'
    },
    VCard: {
      rounded: 'lg',
      elevation: 1
    },
    VTextField: {
      rounded: 'md',
      variant: 'outlined',
      density: 'compact'
    }
  }
})

// 导出主题名称映射
export const themeOptions = [
  { title: '温暖蓝', value: 'statboxLight', color: '#5B9BD5' },
  { title: '翡翠绿', value: 'statboxGreen', color: '#26A69A' },
  { title: '薰衣草紫', value: 'statboxPurple', color: '#7E57C2' },
  { title: '珊瑚橙', value: 'statboxOrange', color: '#FF7043' },
  { title: '玫瑰红', value: 'statboxRose', color: '#EC407A' },
]
