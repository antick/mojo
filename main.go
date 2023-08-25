package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	configuration "mojo/config"
	"mojo/mods"
	"mojo/tools"
	"os"
	"path/filepath"
	"sort"
)

//go:embed all:frontend/dist
var assets embed.FS

var config = configuration.Config()

func buildMods(modBuildPath string) error {
	fmt.Println("-----------------------------------")
	err := tools.Cleanup(modBuildPath)
	if err != nil {
		return err
	}

	modConfig := mods.ModConfig()

	// Extract and sort the mod names
	keys := make([]string, 0, len(modConfig.Mods))
	for k := range modConfig.Mods {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Processing and building descriptor.mod file
	err = tools.ProcessFile(
		modBuildPath,
		config.ModDescriptorPath,
		filepath.Join(modBuildPath, "descriptor.mod"),
		tools.BaseReplacements(),
	)

	if err != nil {
		fmt.Println("Error while processing descriptor.mod file")
		return err
	}

	for _, modName := range keys {
		modDetails := modConfig.Mods[modName]

		if !modDetails.Enabled {
			fmt.Printf("❗️%s is disabled, skipping \n", modName)
			continue
		} else {
			fmt.Printf("📦 Building %s\n", modName)
		}

		err := tools.Build(
			modBuildPath,
			modName,
			modDetails.Replacements,
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("✅ Build successful in %s folder \n", modBuildPath)
	fmt.Println("-----------------------------------")

	return nil
}

func pullCk3GameFiles() error {
	err := tools.Pull(config.Ck3PullMapping)
	if err != nil {
		return err
	}

	fmt.Printf("📦 Pulled CK3 game files to %s \n", config.ModCk3Path)
	return nil
}

func updateLauncherSettings() error {
	err := tools.UpdateGameDataPath(config.LauncherSettingsPath, config.GameDataPath)
	if err != nil {
		fmt.Println("Error updating launcher settings")
		return err
	}

	fmt.Println("✏️ Updated launcher settings")
	return nil
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "Mojo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	configuration.InitDB(configuration.GetDBPath())
	defer configuration.CloseDB()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----------------")
	fmt.Println("Welcome to Mojo!")
	fmt.Println("----------------")
	fmt.Println("Please enter your choice:")
	fmt.Println("0: Start the app")
	fmt.Println("1: Build Mod (Local)")
	fmt.Println("2: Build Mod (Game)")
	fmt.Println("3: Pull CK3 Game Files")
	fmt.Println("4: Update Game Path in Launcher Settings")
	fmt.Print("\nYour choice (1/2/3/4): ")

	answer, _ := reader.ReadString('\n')

	switch answer {
	case "0\n":
		err := buildMods(config.ModBuildPathLocal)
		if err != nil {
			fmt.Printf("Error building mods: %v \n", err)
		}

	case "1\n":
		err := buildMods(config.ModBuildPathLocal)
		if err != nil {
			fmt.Printf("Error building mods: %v \n", err)
		}

	case "2\n":
		err := buildMods(config.ModBuildPath)
		if err != nil {
			fmt.Printf("Error building mods: %v \n", err)
		}

	case "3\n":
		err := pullCk3GameFiles()
		if err != nil {
			fmt.Printf("Error pulling CK3 game files: %v \n", err)
		}

	case "4\n":
		err := updateLauncherSettings()
		if err != nil {
			fmt.Printf("Error updating launcher settings: %v \n", err)
		}

	default:
		fmt.Println("Invalid choice")
	}
}
