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
	ctx             context.Context
	configService   *services.ConfigService
	startupService  *services.StartupService
	templateService *services.TemplateService
	configDir       string // 配置文件目录 (~/.statbox)
	dataDir         string // 数据目录 (可自定义)
	startHidden     bool
}

// NewApp 创建应用实例
func NewApp() *App {
	// 获取用户配置目录
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".statbox")
	
	// 创建配置服务
	configService := services.NewConfigService(configDir)
	
	// 获取数据目录（可能是自定义的）
	dataDir := configService.GetDataDir()
	templatesDir := filepath.Join(dataDir, "templates")

	return &App{
		configDir:       configDir,
		dataDir:         dataDir,
		configService:   configService,
		templateService: services.NewTemplateService(templatesDir),
	}
}

// startup 应用启动时调用
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 如果是静默启动，隐藏窗口
	if a.startHidden {
		wailsRuntime.WindowHide(ctx)
	}

	// 确保配置目录存在
	a.configService.EnsureConfigDir()

	// 确保模板目录存在
	a.templateService.EnsureTemplatesDir()
	
	// 确保模板目录存在
	a.templateService.EnsureTemplatesDir()

	// 初始化开机自启动服务
	startupService, err := services.NewStartupService("StatBox")
	if err != nil {
		fmt.Println("Failed to create startup service:", err)
	} else {
		a.startupService = startupService
	}

	// 加载配置
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
}

// shutdown 应用关闭时调用
func (a *App) shutdown(ctx context.Context) {
	// 清理资源
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

// GetTemplatePath 获取模板文件的完整路径
func (a *App) GetTemplatePath(relativePath string) string {
	return a.templateService.GetFilePath(relativePath)
}

// GetTemplateModules 获取所有模板模块
func (a *App) GetTemplateModules() ([]services.TemplateModule, error) {
	return a.templateService.GetModules()
}

// CreateTemplateModule 创建新模板模块
func (a *App) CreateTemplateModule(name string) error {
	return a.templateService.CreateModule(name)
}

// DeleteTemplateModule 删除模板模块
func (a *App) DeleteTemplateModule(name string) error {
	return a.templateService.DeleteModule(name)
}

// DeleteTemplateFile 删除模板文件
func (a *App) DeleteTemplateFile(relativePath string) error {
	return a.templateService.DeleteFile(relativePath)
}

// CreateTemplateFile 创建模板文件
func (a *App) CreateTemplateFile(module string, fileName string, content string) error {
	return a.templateService.CreateFile(module, fileName, content)
}

// OpenFolderInExplorer 在资源管理器中打开文件夹
func (a *App) OpenFolderInExplorer(folderPath string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", folderPath)
	case "darwin":
		cmd = exec.Command("open", folderPath)
	default: // linux
		cmd = exec.Command("xdg-open", folderPath)
	}
	return cmd.Start()
}

// OpenTemplateModuleFolder 在资源管理器中打开模块文件夹
func (a *App) OpenTemplateModuleFolder(moduleName string) error {
	folderPath := a.templateService.GetModulePath(moduleName)
	return a.OpenFolderInExplorer(folderPath)
}

// OpenTemplateFileFolder 在资源管理器中打开文件所在文件夹
func (a *App) OpenTemplateFileFolder(relativePath string) error {
	filePath := a.templateService.GetFilePath(relativePath)
	folderPath := filepath.Dir(filePath)
	return a.OpenFolderInExplorer(folderPath)
}

// OpenDirectoryDialog 打开目录选择对话框
func (a *App) OpenDirectoryDialog(title string) (string, error) {
	return wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		Title: title,
	})
}

// GetConfigDir 获取配置目录路径
func (a *App) GetConfigDir() string {
	return a.configDir
}

// GetTemplatesDir 获取模板目录路径
func (a *App) GetTemplatesDir() string {
	return a.templateService.GetModulePath("")
}

// GetDataDir 获取数据目录路径
func (a *App) GetDataDir() string {
	return a.dataDir
}

// SetDataDir 设置数据目录路径
func (a *App) SetDataDir(newDataDir string) error {
	// 验证新路径
	if newDataDir == "" {
		return fmt.Errorf("数据目录路径不能为空")
	}

	// 检查新路径是否与当前路径相同
	if newDataDir == a.dataDir {
		return fmt.Errorf("新路径与当前路径相同")
	}

	// 检查选择的是否是 templates 文件夹本身
	// 如果选择的目录名是 "templates"，则使用其父目录作为数据目录
	if filepath.Base(newDataDir) == "templates" {
		newDataDir = filepath.Dir(newDataDir)
	}

	// 确保新目录存在
	if err := os.MkdirAll(newDataDir, 0755); err != nil {
		return fmt.Errorf("无法创建目录: %v", err)
	}

	// 确保必要的子目录存在
	templatesDir := filepath.Join(newDataDir, "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return fmt.Errorf("无法创建模板目录: %v", err)
	}

	// 更新配置
	config, err := a.configService.LoadConfig()
	if err != nil {
		return err
	}
	config.DataDir = newDataDir
	if err := a.configService.SaveConfig(config); err != nil {
		return err
	}

	// 更新应用状态
	a.dataDir = newDataDir
	a.templateService = services.NewTemplateService(templatesDir)

	return nil
}
