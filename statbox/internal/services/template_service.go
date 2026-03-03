package services

import (
	"os"
	"path/filepath"
	"strings"
)

// TemplateFile 模板文件
type TemplateFile struct {
	Name     string `json:"name"`
	Path     string `json:"path"`     // 相对路径
	Language string `json:"language"` // R, Python, Julia, etc.
	Code     string `json:"code"`
}

// TemplateModule 模板模块
type TemplateModule struct {
	Name  string         `json:"name"`
	Files []TemplateFile `json:"files"`
}

// TemplateService 模板管理服务
type TemplateService struct {
	templatesDir string
}

// NewTemplateService 创建模板管理服务
func NewTemplateService(templatesDir string) *TemplateService {
	return &TemplateService{
		templatesDir: templatesDir,
	}
}

// EnsureTemplatesDir 确保模板目录存在
func (s *TemplateService) EnsureTemplatesDir() error {
	return os.MkdirAll(s.templatesDir, 0755)
}

// GetModules 获取所有模块
func (s *TemplateService) GetModules() ([]TemplateModule, error) {
	// 确保目录存在
	if err := s.EnsureTemplatesDir(); err != nil {
		return nil, err
	}

	// 读取目录
	entries, err := os.ReadDir(s.templatesDir)
	if err != nil {
		return nil, err
	}

	modules := []TemplateModule{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		moduleName := entry.Name()
		module := TemplateModule{
			Name:  moduleName,
			Files: []TemplateFile{},
		}

		// 读取模块下的文件
		modulePath := filepath.Join(s.templatesDir, moduleName)
		files, err := s.getTemplateFiles(modulePath, moduleName)
		if err != nil {
			continue
		}

		module.Files = files
		modules = append(modules, module)
	}

	return modules, nil
}

// getTemplateFiles 递归获取模板文件
func (s *TemplateService) getTemplateFiles(dir string, module string) ([]TemplateFile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	files := []TemplateFile{}
	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			// 递归读取子目录
			subFiles, err := s.getTemplateFiles(fullPath, module)
			if err != nil {
				continue
			}
			files = append(files, subFiles...)
			continue
		}

		// 检查文件扩展名
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if !isValidTemplateExt(ext) {
			continue
		}

		// 读取文件内容
		content, err := os.ReadFile(fullPath)
		if err != nil {
			continue
		}

		// 计算相对路径
		relPath, err := filepath.Rel(s.templatesDir, fullPath)
		if err != nil {
			continue
		}

		// 推断语言类型
		language := getLanguageFromExt(ext)

		files = append(files, TemplateFile{
			Name:     entry.Name(),
			Path:     filepath.ToSlash(relPath),
			Language: language,
			Code:     string(content),
		})
	}

	return files, nil
}

// CreateModule 创建新模块
func (s *TemplateService) CreateModule(name string) error {
	modulePath := filepath.Join(s.templatesDir, name)
	return os.MkdirAll(modulePath, 0755)
}

// DeleteModule 删除模块
func (s *TemplateService) DeleteModule(name string) error {
	modulePath := filepath.Join(s.templatesDir, name)
	return os.RemoveAll(modulePath)
}

// DeleteFile 删除文件
func (s *TemplateService) DeleteFile(relativePath string) error {
	fullPath := filepath.Join(s.templatesDir, relativePath)
	return os.Remove(fullPath)
}

// CreateFile 创建文件
func (s *TemplateService) CreateFile(module string, fileName string, content string) error {
	// 确保模块存在
	modulePath := filepath.Join(s.templatesDir, module)
	if err := os.MkdirAll(modulePath, 0755); err != nil {
		return err
	}

	// 创建文件
	filePath := filepath.Join(modulePath, fileName)
	return os.WriteFile(filePath, []byte(content), 0644)
}

// GetFilePath 获取文件的完整路径
func (s *TemplateService) GetFilePath(relativePath string) string {
	return filepath.Join(s.templatesDir, relativePath)
}

// GetModulePath 获取模块的完整路径
func (s *TemplateService) GetModulePath(moduleName string) string {
	return filepath.Join(s.templatesDir, moduleName)
}

// isValidTemplateExt 检查是否是有效的模板文件扩展名
func isValidTemplateExt(ext string) bool {
	validExts := map[string]bool{
		".r":    true,
		".rmd":  true,
		".py":   true,
		".jl":   true,
		".md":   true,
		".txt":  true,
		".tex":  true,
		".json": true,
	}
	return validExts[ext]
}

// getLanguageFromExt 根据扩展名获取语言类型
func getLanguageFromExt(ext string) string {
	switch ext {
	case ".r", ".rmd":
		return "R"
	case ".py":
		return "Python"
	case ".jl":
		return "Julia"
	case ".md":
		return "Markdown"
	case ".tex":
		return "LaTeX"
	case ".json":
		return "JSON"
	default:
		return "Text"
	}
}
