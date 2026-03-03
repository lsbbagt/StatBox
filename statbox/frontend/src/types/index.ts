export type ModuleType = 'bookmarks' | 'templates' | 'settings'

export interface ModuleItem {
  id: ModuleType
  name: string
  icon: string
}

export interface BookmarkFolder {
  id: string
  name: string
  icon?: string
  order: number
  items: BookmarkItem[]
}

export interface BookmarkItem {
  id: string
  name: string
  url: string
  icon?: string
  tags?: string[]
  createdAt: string
  visitCount: number
}

export interface TemplateFile {
  name: string
  path: string
  language: 'R' | 'Python' | 'Julia'
  description?: string
  tags?: string[]
}

export interface CommandFile {
  name: string
  path: string
  description: string
  type: string
  category?: string
}

export interface AppSettings {
  startupWithSystem: boolean
  minimizeToTray: boolean
  defaultBrowser: 'internal' | 'external'
  theme: 'light' | 'dark'
  fontSize: number
}
