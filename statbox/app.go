package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"statbox/internal/services"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App 应用主结构
type App struct {
	ctx            context.Context
	configService  *services.ConfigService
	hotkeyService  *services.HotkeyService
	startupService *services.StartupService
	configDir      string
	startHidden    bool
}

// NewApp 创建应用实例
func NewApp() *App {
	// 获取用户配置目录
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".statbox")

	return &App{
		configDir:     configDir,
		configService: services.NewConfigService(configDir),
		hotkeyService: services.NewHotkeyService(),
	}
}

// startup 应用启动时调用
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 设置上下文
	a.hotkeyService.SetContext(ctx)

	// 如果是静默启动，隐藏窗口
	if a.startHidden {
		wailsRuntime.WindowHide(ctx)
	}

	// 确保配置目录存在
	a.configService.EnsureConfigDir()

	// 初始化开机自启动服务
	startupService, err := services.NewStartupService("StatBox")
	if err != nil {
		fmt.Println("Failed to create startup service:", err)
	} else {
		a.startupService = startupService
	}

	// 加载配置并注册快捷键
	config, err := a.configService.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return
	}

	// 如果启用了开机自启动，确保注册（静默模式）
	if config.Features.StartupWithSystem {
		if a.startupService != nil {
			err := a.startupService.EnableStartup()
			if err != nil {
				fmt.Println("Failed to enable startup:", err)
			}
		}
	}

	// 注册全局快捷键（如果配置为全局）
	if config.Features.CommandScope == "global" {
		err = a.hotkeyService.RegisterHotkey(config.Features.CommandHotkey, func() {
			// 显示窗口
			a.ShowWindow()
		})
		if err != nil {
			fmt.Println("Failed to register hotkey:", err)
		}
	}
}

// shutdown 应用关闭时调用
func (a *App) shutdown(ctx context.Context) {
	// 注销快捷键
	a.hotkeyService.UnregisterHotkey()
}

// ShowWindow 显示窗口
func (a *App) ShowWindow() {
	wailsRuntime.WindowShow(a.ctx)
	wailsRuntime.WindowUnminimise(a.ctx)
}

// HideWindow 隐藏窗口（最小化到托盘）
func (a *App) HideWindow() {
	wailsRuntime.WindowHide(a.ctx)
}

// GetConfig 获取配置
func (a *App) GetConfig() (*services.Config, error) {
	return a.configService.LoadConfig()
}

// SaveConfig 保存配置
func (a *App) SaveConfig(config *services.Config) error {
	return a.configService.SaveConfig(config)
}

// EnableStartup 启用开机自启动
func (a *App) EnableStartup() error {
	if a.startupService == nil {
		return fmt.Errorf("startup service not initialized")
	}
	return a.startupService.EnableStartup()
}

// DisableStartup 禁用开机自启动
func (a *App) DisableStartup() error {
	if a.startupService == nil {
		return fmt.Errorf("startup service not initialized")
	}
	return a.startupService.DisableStartup()
}

// IsStartupEnabled 检查是否启用自启动
func (a *App) IsStartupEnabled() (bool, error) {
	if a.startupService == nil {
		return false, fmt.Errorf("startup service not initialized")
	}
	return a.startupService.IsStartupEnabled()
}

// RegisterHotkey 注册快捷键
func (a *App) RegisterHotkey(key string) error {
	return a.hotkeyService.RegisterHotkey(key, func() {
		a.ShowWindow()
	})
}

// UnregisterHotkey 注销快捷键
func (a *App) UnregisterHotkey() error {
	return a.hotkeyService.UnregisterHotkey()
}

// OpenFileWithDefault 使用系统默认应用打开文件
func (a *App) OpenFileWithDefault(filePath string) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", filePath)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	default: // linux
		cmd = exec.Command("xdg-open", filePath)
	}

	return cmd.Start()
}
