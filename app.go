package main

import (
	"context"
	"fmt"
	configuration "mojo/configs"
	"mojo/scripts"
	"path/filepath"
)

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

var modConfig = configuration.ModConfig()

// startup is called when the app starts. The context is saved, so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) BuildModsInLocal(selectedModKeys []string, buildCombinedMod bool) error {
	if len(selectedModKeys) == 0 {
		return fmt.Errorf("No mods selected for building")
	}

	if err := scripts.BuildModFile(config.ModBuildPathLocal); err != nil {
		fmt.Println("Error while processing mojo.mod file")
		return err
	}

	modBuildPath := filepath.Join(config.ModBuildPathLocal, modConfig.CombinedMod.Replacements["modFolderName"])
	err := scripts.BuildMods(modBuildPath, selectedModKeys)
	if err != nil {
		fmt.Printf("Error building mods in local: %v \n", err)
		return err
	}

	return err
}

func (a *App) BuildModsInGame(selectedModKeys []string, buildCombinedMod bool) error {
	if len(selectedModKeys) == 0 {
		return fmt.Errorf("No mods selected for building")
	}

	if err := scripts.BuildModFile(config.GameCustomModPath); err != nil {
		fmt.Println("Error while processing mojo.mod file")
		return err
	}

	err := scripts.BuildMods(config.ModBuildPath, selectedModKeys)
	if err != nil {
		fmt.Printf("Error building mods in game: %v \n", err)
		return err
	}

	return err
}

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
