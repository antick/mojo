package main

import (
	"context"
	"fmt"
	"mojo/mods"
	"mojo/tools"
	"path/filepath"
	"sort"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) buildMods(modBuildPath string) error {
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
