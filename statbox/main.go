package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "StatBox",
		Width:  1024,
		Height: 700,
		MinWidth: 800,
		MinHeight: 600,
		// 启用无边框窗口
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// 应用设计系统的背景色
		BackgroundColour: &options.RGBA{R: 248, G: 250, B: 251, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
