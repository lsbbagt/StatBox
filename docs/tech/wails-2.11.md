# Wails v2.11.0 技术文档

> 本文档基于系统实际安装的 Wails v2.11.0 版本

---

## 📋 版本信息

- **版本号：** v2.11.0
- **发布日期：** 2024年
- **Go版本要求：** 1.18+
- **支持平台：** Windows, macOS, Linux

---

## ✨ 主要特性

### 1. 原生窗口控制
- ✅ 自定义标题栏（Frameless窗口）
- ✅ 窗口最小化、最大化、关闭
- ✅ 系统托盘支持
- ✅ 窗口透明度和置顶

### 2. 前端框架支持
- Vue 3
- React
- Svelte
- Vanilla JS/TS

### 3. 运行时API
- 事件系统
- 日志记录
- 对话框
- 菜单系统
- 剪贴板操作

### 4. 构建优化
- 单文件可执行程序
- 小体积打包
- 跨平台编译

---

## 🎯 StatBox 中的关键应用

### 自定义标题栏实现

```go
// main.go
func main() {
    err := wails.Run(&options.App{
        Title:      "StatBox",
        Width:      1200,
        Height:     800,
        Frameless:  true,  // 关键：启用无边框窗口
        AssetServer: &assetserver.Options{
            Assets: assets,
        },
        BackgroundColour: &options.RGBA{R: 18, G: 18, B: 18, A: 1},
        OnStartup:        app.startup,
        OnShutdown:       app.shutdown,
    })
}
```

**前端窗口控制：**

```typescript
import { WindowMinimise, WindowMaximise, WindowUnmaximise, WindowClose } from '../../wailsjs/runtime'

// 最小化
WindowMinimise()

// 切换最大化
if (isMaximized.value) {
    await WindowUnmaximise()
} else {
    await WindowMaximise()
}

// 关闭窗口
WindowClose()
```

### 系统托盘集成

```go
// 托盘服务
func (s *TrayService) CreateTrayMenu() *menu.Menu {
    trayMenu := menu.NewMenu()
    
    trayMenu.AddText("显示窗口", keys.CmdOrCtrl("s"), func(_ *menu.CallbackData) {
        s.ShowApp()
    })
    
    trayMenu.AddSeparator()
    
    trayMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
        s.onExit()
    })
    
    return trayMenu
}
```

### 全局快捷键

```go
// 注册全局快捷键
runtime.EventsOn(ctx, "hotkey", func(optionalData ...interface{}) {
    // 处理快捷键事件
})
```

---

## 🔧 最佳实践

### 1. 应用生命周期

```go
type App struct {
    ctx           context.Context
    configService *services.ConfigService
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    // 初始化服务
    a.configService.EnsureConfigDir()
}

func (a *App) shutdown(ctx context.Context) {
    // 清理资源
    a.configService.SaveConfig()
}
```

### 2. 前后端通信

**后端暴露方法：**

```go
func (a *App) GetConfig() *models.AppConfig {
    config, _ := a.configService.LoadConfig()
    return config
}

func (a *App) SaveConfig(config *models.AppConfig) error {
    return a.configService.SaveConfig(config)
}
```

**前端调用：**

```typescript
import { GetConfig, SaveConfig } from '../../wailsjs/go/main/App'

const config = await GetConfig()
await SaveConfig(newConfig)
```

### 3. 事件系统

**后端发送事件：**

```go
runtime.EventsEmit(ctx, "config-updated", newConfig)
```

**前端监听事件：**

```typescript
import { EventsOn, EventsOff } from '../../wailsjs/runtime'

EventsOn("config-updated", (config) => {
    console.log("Config updated:", config)
})

// 组件卸载时移除监听
EventsOff("config-updated")
```

---

## ⚠️ 已知问题与解决方案

### 1. 窗口拖拽问题

**问题：** 无边框窗口无法拖拽移动

**解决方案：** 在标题栏元素上添加 CSS

```css
.title-bar {
    -webkit-app-region: drag;
}

.title-bar :deep(.v-btn),
.title-bar :deep(.v-text-field) {
    -webkit-app-region: no-drag;  /* 按钮和输入框可点击 */
}
```

### 2. 滚动条隐藏

**问题：** 需要隐藏默认滚动条保持界面简洁

**解决方案：**

```css
.content-wrapper {
    scrollbar-width: none;  /* Firefox */
    -ms-overflow-style: none;  /* IE/Edge */
}

.content-wrapper::-webkit-scrollbar {
    display: none;  /* Chrome/Safari */
}
```

---

## 📚 API参考

### 窗口控制

| 方法 | 说明 |
|------|------|
| `WindowMinimise()` | 最小化窗口 |
| `WindowMaximise()` | 最大化窗口 |
| `WindowUnmaximise()` | 取消最大化 |
| `WindowClose()` | 关闭窗口 |
| `WindowShow()` | 显示窗口 |
| `WindowHide()` | 隐藏窗口 |
| `WindowToggleMaximise()` | 切换最大化状态 |

### 运行时

| 方法 | 说明 |
|------|------|
| `EventsOn(event, callback)` | 监听事件 |
| `EventsOff(event)` | 移除事件监听 |
| `EventsEmit(event, data)` | 发送事件 |
| `LogDebug(msg)` | 调试日志 |
| `LogInfo(msg)` | 信息日志 |
| `LogError(msg)` | 错误日志 |

---

## 🔗 相关资源

- [Wails 官方文档](https://wails.io/docs/introduction)
- [Wails GitHub](https://github.com/wailsapp/wails)
- [示例项目](https://github.com/wailsapp/wails/tree/master/samples)

---

## 📅 更新记录

- 2025-02-28: 创建文档，基于 v2.11.0
