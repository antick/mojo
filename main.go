package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	configuration "mojo/configs"
)

//go:embed all:frontend/dist
var assets embed.FS

var config = configuration.Config()

func main() {
	//configuration.InitDB(configuration.GetDBPath())
	//defer configuration.CloseDB()
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "Mojo",
		Width:  1000,
		Height: 540,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        false,
		DisableResize:    true,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
