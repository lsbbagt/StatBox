package services

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

// StartupService 处理开机自启动
type StartupService struct {
	appName string
	appPath string
}

var (
	advapi32          = syscall.NewLazyDLL("advapi32.dll")
	regOpenKeyEx      = advapi32.NewProc("RegOpenKeyExW")
	regSetValueEx     = advapi32.NewProc("RegSetValueExW")
	regDeleteValue    = advapi32.NewProc("RegDeleteValueW")
	regCloseKey       = advapi32.NewProc("RegCloseKey")
)

const (
	HKEY_CURRENT_USER        = 0x80000001
	KEY_ALL_ACCESS           = 0xF003F
	REG_SZ                   = 1
)

// NewStartupService 创建自启动服务
func NewStartupService(appName string) (*StartupService, error) {
	// 获取当前可执行文件路径
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	return &StartupService{
		appName: appName,
		appPath: exePath,
	}, nil
}

// EnableStartup 启用开机自启动
func (s *StartupService) EnableStartup() error {
	keyPath := "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
	var hKey uintptr

	// 打开注册表项
	ret, _, err := regOpenKeyEx.Call(
		HKEY_CURRENT_USER,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(keyPath))),
		0,
		KEY_ALL_ACCESS,
		uintptr(unsafe.Pointer(&hKey)),
	)

	if ret != 0 {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer regCloseKey.Call(hKey)

	// 设置值
	value := syscall.StringToUTF16(s.appPath)
	ret, _, err = regSetValueEx.Call(
		hKey,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s.appName))),
		0,
		REG_SZ,
		uintptr(unsafe.Pointer(&value[0])),
		uintptr(len(value)*2),
	)

	if ret != 0 {
		return fmt.Errorf("failed to set registry value: %v", err)
	}

	return nil
}

// DisableStartup 禁用开机自启动
func (s *StartupService) DisableStartup() error {
	keyPath := "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
	var hKey uintptr

	// 打开注册表项
	ret, _, err := regOpenKeyEx.Call(
		HKEY_CURRENT_USER,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(keyPath))),
		0,
		KEY_ALL_ACCESS,
		uintptr(unsafe.Pointer(&hKey)),
	)

	if ret != 0 {
		return fmt.Errorf("failed to open registry key: %v", err)
	}
	defer regCloseKey.Call(hKey)

	// 删除值
	ret, _, err = regDeleteValue.Call(
		hKey,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s.appName))),
	)

	if ret != 0 {
		return fmt.Errorf("failed to delete registry value: %v", err)
	}

	return nil
}

// IsStartupEnabled 检查是否已启用自启动
func (s *StartupService) IsStartupEnabled() (bool, error) {
	// 检查注册表项是否存在
	registryPath := filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup", s.appName+".lnk")
	
	_, err := os.Stat(registryPath)
	return !os.IsNotExist(err), nil
}
