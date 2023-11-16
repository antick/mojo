package main

import (
	"context"
	"fmt"
	"mojo/tools"
	"path/filepath"
	"sort"
)

type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func sortModList() []string {
	keys := make([]string, 0, len(config.SubMods))
	for k := range config.SubMods {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func buildDescriptorFile(modBuildPath string) error {
	err := tools.ProcessFile(
		modBuildPath,
		config.ModDescriptorSourcePath,
		filepath.Join(modBuildPath, "descriptor.mod"),
		config.MainMod.Replacements,
	)

	return err
}

func buildThumbnailFile(modBuildPath string) error {
	err := tools.ProcessFile(
		modBuildPath,
		config.ThumbnailSourcePath,
		filepath.Join(modBuildPath, "thumbnail.png"),
		map[string]string{},
	)

	return err
}

func BuildMods(modBuildPath string) error {
	fmt.Println("-----------------------------------")
	err := tools.Cleanup(modBuildPath)
	if err != nil {
		return err
	}

	if err = buildDescriptorFile(modBuildPath); err != nil {
		fmt.Println("Error while processing descriptor.mod file")
		return err
	}

	if err = buildThumbnailFile(modBuildPath); err != nil {
		fmt.Println("Error while processing thumbnail.png file")
		return err
	}

	sortedModList := sortModList()
	for _, modName := range sortedModList {
		modDetails := config.SubMods[modName]

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

func (a *App) BuildModsInLocal() error {
	err := BuildMods(filepath.Join(config.ModBuildPathLocal, config.MainMod.Replacements["modFolderName"]))
	if err != nil {
		fmt.Printf("Error building mods in local: %v \n", err)
		return err
	}

	return err
}

func (a *App) BuildModsInGame() error {
	err := BuildMods(config.ModBuildPath)
	if err != nil {
		fmt.Printf("Error building mods in game: %v \n", err)
		return err
	}

	return err
}

func (a *App) PullCk3GameFiles() error {
	err := tools.Pull(config.Ck3PullMapping)
	if err != nil {
		fmt.Printf("Error pulling CK3 game files: %v \n", err)
		return err
	}

	fmt.Printf("📦 Pulled CK3 game files to %s \n", config.ModCk3Path)
	return nil
}

func (a *App) UpdateLauncherSettings() error {
	err := tools.UpdateGameDataPath(config.LauncherSettingsPath, config.GameDataPath)
	if err != nil {
		fmt.Println("Error updating launcher settings")
		return err
	}

	fmt.Println("✏️ Updated launcher settings")
	return nil
}
