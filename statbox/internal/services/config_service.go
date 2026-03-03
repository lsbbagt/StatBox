package services

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// WindowConfig 窗口配置
type WindowConfig struct {
	Width             int  `json:"width"`
	Height            int  `json:"height"`
	SidebarVisible    bool `json:"sidebarVisible"`
	RightPanelVisible bool `json:"rightPanelVisible"`
}

// FeatureConfig 功能配置
type FeatureConfig struct {
	StartupWithSystem bool   `json:"startupWithSystem"`
	MinimizeToTray    bool   `json:"minimizeToTray"`
	CommandScope      string `json:"commandScope"`      // "global" or "internal"
	DefaultBrowser    string `json:"defaultBrowser"`    // "internal" or "external"
}

// EditorConfig 编辑器配置
type EditorConfig struct {
	Theme    string `json:"theme"`
	FontSize int    `json:"fontSize"`
	WordWrap bool   `json:"wordWrap"`
}

// Config 应用配置
type Config struct {
	Version  string        `json:"version"`
	Window   WindowConfig  `json:"window"`
	Features FeatureConfig `json:"features"`
	Editor   EditorConfig  `json:"editor"`
}

// ConfigService 配置服务
type ConfigService struct {
	configDir  string
	configFile string
}

// NewConfigService 创建配置服务
func NewConfigService(configDir string) *ConfigService {
	return &ConfigService{
		configDir:  configDir,
		configFile: filepath.Join(configDir, "config.json"),
	}
}

// LoadConfig 加载配置
func (s *ConfigService) LoadConfig() (*Config, error) {
	// 检查配置文件是否存在
	if _, err := os.Stat(s.configFile); os.IsNotExist(err) {
		// 返回默认配置
		return s.defaultConfig(), nil
	}

	data, err := os.ReadFile(s.configFile)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig 保存配置
func (s *ConfigService) SaveConfig(config *Config) error {
	// 确保目录存在
	if err := os.MkdirAll(s.configDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.configFile, data, 0644)
}

// EnsureConfigDir 确保配置目录存在
func (s *ConfigService) EnsureConfigDir() error {
	return os.MkdirAll(s.configDir, 0755)
}

// defaultConfig 默认配置
func (s *ConfigService) defaultConfig() *Config {
	return &Config{
		Version: "1.0",
		Window: WindowConfig{
			Width:             1200,
			Height:            800,
			SidebarVisible:    true,
			RightPanelVisible: true,
		},
		Features: FeatureConfig{
			StartupWithSystem: true,
			MinimizeToTray:    true,
			CommandScope:      "internal",
			DefaultBrowser:    "internal",
		},
		Editor: EditorConfig{
			Theme:    "light",
			FontSize: 14,
			WordWrap: true,
		},
	}
}
