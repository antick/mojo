package main

import (
	"bufio"
	"fmt"
	configuration "mojo/config"
	"mojo/mods"
	"mojo/tools"
	"os"
	"sort"
)

var config = configuration.Config()

func buildMods() error {
	fmt.Println("-----------------------------------")
	err := tools.Cleanup(config.ModBuildPath)
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

	for _, modName := range keys {
		modDetails := modConfig.Mods[modName]

		if !modDetails.Enabled {
			fmt.Printf("❗️%s is disabled, skipping \n", modName)
			continue
		} else {
			fmt.Printf("📦 Building %s\n", modName)
		}

		err := tools.Build(modName, modDetails.Replacements)
		if err != nil {
			return err
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("✅ Build successful in %s folder \n", config.ModBuildPath)
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
		return err
	}

	fmt.Println("📦 Updated launcher settings")
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your choice:")
	fmt.Println("1: Build Mod")
	fmt.Println("2: Pull CK3 Game Files")
	fmt.Println("3: Update Game Path in Launcher Settings")
	fmt.Print("\nYour choice (1/2/3): ")

	answer, _ := reader.ReadString('\n')

	switch answer {
	case "1\n":
		err := buildMods()
		if err != nil {
			fmt.Printf("Error building mods: %v \n", err)
		}
	case "2\n":
		err := pullCk3GameFiles()
		if err != nil {
			fmt.Printf("Error pulling CK3 game files: %v \n", err)
		}
	case "3\n":
		err := updateLauncherSettings()
		if err != nil {
			fmt.Printf("Error updating launcher settings: %v \n", err)
		}
	default:
		fmt.Println("Invalid choice")
	}
}
