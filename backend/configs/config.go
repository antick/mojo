package configs

import (
	"os"
	"path/filepath"
)

type Type struct {
	GameDataPath            string
	ModBuildPathLocal       string
	ModPath                 string
	GameCustomModPath       string
	ModDescriptorSourcePath string
	ThumbnailSourcePath     string
	OriginalCk3Path         string
	ModCk3Path              string
	Ck3PullMapping          map[string]string
	ModIdPrefix             string
	SyncedCk3Version        string
	ModFoldersToProcess     []string
}

func Config() *Type {
	userHomeDir, _ := os.UserHomeDir()

	// Path of the folder where CK3 is installed in your system
	installedCk3GamePath := filepath.Join(`C:\Program Files (x86)\Steam\steamapps\common\Crusader Kings III\game`)

	// Path of the folder where CK3 installs the custom user mods
	gameCustomModPath := filepath.Join(userHomeDir, "OneDrive", "Documents", "Paradox Interactive", "Crusader Kings III", "mod")

	// Path of the folder where we are storing a copy of game files to track the changes that comes in new releases
	modCk3Path := filepath.Join("build", "ck3-game")

	// Path of the folder where source code of the mods is stored
	modPath := filepath.Join("backend", "mods")

	return &Type{
		// Current ck3 version that we have synced in "game" folder
		SyncedCk3Version: "1.13.0.3",

		// Path of the folder where CK3 stores all the game data including your save files and mods
		GameDataPath: filepath.Join(userHomeDir, "OneDrive", "Documents", "Paradox Interactive", "Crusader Kings III"),

		// This is where all the mods will be built for local testing
		ModBuildPathLocal: "build/mods",

		ModPath:                 modPath,
		GameCustomModPath:       gameCustomModPath,
		ModDescriptorSourcePath: filepath.Join(modPath, "descriptor.mod"),
		ThumbnailSourcePath:     filepath.Join(modPath, "thumbnail.png"),
		OriginalCk3Path:         installedCk3GamePath,
		ModCk3Path:              modCk3Path,

		Ck3PullMapping: map[string]string{
			filepath.Join(installedCk3GamePath, "common"):               filepath.Join(modCk3Path, "common"),
			filepath.Join(installedCk3GamePath, "events"):               filepath.Join(modCk3Path, "events"),
			filepath.Join(installedCk3GamePath, "gui"):                  filepath.Join(modCk3Path, "gui"),
			filepath.Join(installedCk3GamePath, "history"):              filepath.Join(modCk3Path, "history"),
			filepath.Join(installedCk3GamePath, "localization/english"): filepath.Join(modCk3Path, "localization/english"),
		},

		// This name will be used as prefix for all the mods in the build folder
		ModIdPrefix: "antick",

		ModFoldersToProcess: []string{
			"common",
			"events",
			"gfx",
			"gui",
			"history",
			"localization",
			"music",
		},
	}
}
