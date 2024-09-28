package main

import (
	"context"
	"fmt"
	configuration "mojo/backend/configs"
	"mojo/backend/scripts"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var modConfig = configuration.ModConfig()

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) BuildModsInLocal(selectedModKeys []string, buildCombinedMod bool) error {
	if len(selectedModKeys) == 0 {
		return fmt.Errorf("No mods selected for building")
	}

	if buildCombinedMod {
		modFileName := modConfig.CombinedMod.Replacements["modFolderName"] + ".mod"
		if err := scripts.BuildModFile(config.ModBuildPathLocal, modFileName); err != nil {
			fmt.Println("Error while processing mojo.mod file")
			return err
		}

		modBuildPath := filepath.Join(config.ModBuildPathLocal, modConfig.CombinedMod.Replacements["modFolderName"])
		err := scripts.BuildCombinedMod(modBuildPath, selectedModKeys)
		if err != nil {
			fmt.Printf("Error building combined mod in local: %v \n", err)
			return err
		}
	} else {
		err := scripts.BuildLooseMods(config.ModBuildPathLocal, selectedModKeys)
		if err != nil {
			fmt.Printf("Error building selected mod(s) in local: %v \n", err)
			return err
		}
	}

	return nil
}

func (a *App) BuildModsInGame(selectedModKeys []string, buildCombinedMod bool) error {
	if len(selectedModKeys) == 0 {
		return fmt.Errorf("No mods selected for building")
	}

	if buildCombinedMod {
		modFileName := modConfig.CombinedMod.Replacements["modFolderName"] + ".mod"
		if err := scripts.BuildModFile(config.GameCustomModPath, modFileName); err != nil {
			fmt.Println("Error while processing mojo.mod file")
			return err
		}

		modBuildPath := filepath.Join(config.GameCustomModPath, modConfig.CombinedMod.Replacements["modFolderName"])
		err := scripts.BuildCombinedMod(modBuildPath, selectedModKeys)
		if err != nil {
			fmt.Printf("Error building combined mod in local: %v \n", err)
			return err
		}
	} else {
		err := scripts.BuildLooseMods(config.GameCustomModPath, selectedModKeys)
		if err != nil {
			fmt.Printf("Error building selected mod(s) in local: %v \n", err)
			return err
		}
	}

	return nil
}

// TODO: change "pull" to "sync"
func (a *App) PullCk3GameFiles() error {
	err := scripts.Pull(config.Ck3PullMapping)
	if err != nil {
		fmt.Printf("Error pulling CK3 game files: %v \n", err)
		return err
	}

	fmt.Printf("📦 Pulled CK3 game files to %s \n", config.ModCk3Path)
	return nil
}

func (a *App) GetModList() (*configuration.ModType, error) {
	return modConfig, nil
}
