import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'statboxLight',
    themes: {
      statboxLight: {
        dark: false,
        colors: {
          // Primary - 温暖蓝
          primary: '#5B9BD5',
          'primary-light': '#8DB8E8',
          'primary-dark': '#3A7BC8',
          
          // Secondary - 柔和灰蓝
          secondary: '#7A8B99',
          'secondary-light': '#9BA8B3',
          'secondary-dark': '#5A6B79',
          
          // Accent - 温暖橙
          accent: '#F4A261',
          'accent-light': '#F7BA8A',
          'accent-dark': '#E08A3E',
          
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
      }
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
