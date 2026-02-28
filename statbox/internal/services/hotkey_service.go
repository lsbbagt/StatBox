package services

import (
	"context"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// HotkeyService 处理全局快捷键
type HotkeyService struct {
	ctx       context.Context
	hotkeyID  int
	keyCode   uint
	modifiers uint
	callback  func()
}

// Windows API 常量
const (
	WM_HOTKEY     = 0x0312
	MOD_ALT       = 0x0001
	MOD_CONTROL   = 0x0002
	MOD_SHIFT     = 0x0004
	MOD_WIN       = 0x0008
	MOD_NOREPEAT  = 0x4000
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	registerHotkey       = user32.NewProc("RegisterHotKey")
	unregisterHotkey     = user32.NewProc("UnregisterHotKey")
	getMessage          = user32.NewProc("GetMessageW")
	translateMessage    = user32.NewProc("TranslateMessage")
	dispatchMessage     = user32.NewProc("DispatchMessageW")
)

// NewHotkeyService 创建快捷键服务
func NewHotkeyService() *HotkeyService {
	return &HotkeyService{
		hotkeyID: 1,
	}
}

// SetContext 设置上下文
func (h *HotkeyService) SetContext(ctx context.Context) {
	h.ctx = ctx
}

// RegisterHotkey 注册全局快捷键
func (h *HotkeyService) RegisterHotkey(key string, callback func()) error {
	// 解析快捷键
	keyCode, modifiers, err := parseHotkey(key)
	if err != nil {
		return err
	}

	h.keyCode = keyCode
	h.modifiers = modifiers
	h.callback = callback

	// 注册快捷键
	ret, _, err := registerHotkey.Call(
		0,
		uintptr(h.hotkeyID),
		uintptr(modifiers),
		uintptr(keyCode),
	)

	if ret == 0 {
		return fmt.Errorf("failed to register hotkey: %v", err)
	}

	// 启动消息监听
	go h.listenForHotkey()

	return nil
}

// UnregisterHotkey 注销快捷键
func (h *HotkeyService) UnregisterHotkey() error {
	ret, _, err := unregisterHotkey.Call(0, uintptr(h.hotkeyID))
	if ret == 0 {
		return fmt.Errorf("failed to unregister hotkey: %v", err)
	}
	return nil
}

// listenForHotkey 监听快捷键消息
func (h *HotkeyService) listenForHotkey() {
	var msg struct {
		hwnd    uintptr
		message uint32
		wParam  uintptr
		lParam  uintptr
		time    uint32
		pt      struct{ x, y int32 }
	}

	for {
		ret, _, _ := getMessage.Call(
			uintptr(unsafe.Pointer(&msg)),
			0,
			0,
			0,
		)

		if ret == 0 {
			break
		}

		if msg.message == WM_HOTKEY {
			// 触发回调
			if h.callback != nil {
				runtime.EventsEmit(h.ctx, "hotkey-triggered", "command-palette")
				h.callback()
			}
		}

		translateMessage.Call(uintptr(unsafe.Pointer(&msg)))
		dispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	}
}

// parseHotkey 解析快捷键字符串
// 支持格式: "~", "ctrl+space", "alt+shift+a" 等
func parseHotkey(key string) (uint, uint, error) {
	var keyCode uint
	var modifiers uint = MOD_NOREPEAT

	// 特殊键处理
	switch key {
	case "~":
		// 波浪号键 (VK_OEM_3)
		keyCode = 0xC0
	case "ctrl+space":
		modifiers |= MOD_CONTROL
		keyCode = 0x20 // VK_SPACE
	case "alt+space":
		modifiers |= MOD_ALT
		keyCode = 0x20
	default:
		// 单个字符键
		if len(key) == 1 {
			keyCode = uint(toupper(key[0]))
		} else {
			return 0, 0, fmt.Errorf("unsupported hotkey: %s", key)
		}
	}

	return keyCode, modifiers, nil
}

func toupper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}
