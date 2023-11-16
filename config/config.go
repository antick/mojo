package config

import (
	"os"
	"path/filepath"
)

type MainModType struct {
	Replacements map[string]string `json:"replacements"`
}

type SubModType struct {
	Enabled      bool              `json:"enabled"`
	Replacements map[string]string `json:"replacements"`
}

type Type struct {
	GameDataPath            string                `json:"gameDataPath"`
	ModBuildPath            string                `json:"modBuildPath"`
	ModBuildPathLocal       string                `json:"modBuildPathLocal"`
	ModPath                 string                `json:"modPath"`
	ModDescriptorSourcePath string                `json:"modDescriptorSourcePath"`
	ThumbnailSourcePath     string                `json:"thumbnailSourcePath"`
	OriginalCk3Path         string                `json:"originalCk3Path"`
	ModCk3Path              string                `json:"modCk3Path"`
	Ck3PullMapping          map[string]string     `json:"ck3PullMapping"`
	ModIdPrefix             string                `json:"modIdPrefix"`
	SyncedCk3Version        string                `json:"syncedCk3Version"`
	ModFoldersToProcess     []string              `json:"modFoldersToProcess"`
	MainMod                 MainModType           `json:"mainMod"`
	SubMods                 map[string]SubModType `json:"subMods"`
}

func Config() *Type {
	userHomeDir, _ := os.UserHomeDir()

	// Path of the folder where CK3 is installed in your system
	installedCk3GamePath := filepath.Join(`C:\Program Files (x86)\Steam\steamapps\common\Crusader Kings III\game`)

	// Path of the folder where CK3 stores all the game data including your save files and mods
	gameDataPath := filepath.Join(userHomeDir, "OneDrive", "Documents", "Paradox Interactive", "Crusader Kings III")

	modCk3Path := "game"

	// Path of the folder where source code of the mods is stored
	modPath := filepath.Join("game", "mojo-mods")

	// This name will be used as prefix for all the mods in the build folder
	ModIdPrefix := "antick"

	return &Type{
		GameDataPath: gameDataPath,
		ModBuildPath: filepath.Join(userHomeDir, "Paradox Interactive", "Crusader Kings III", "mod", "mojo-flavor"),

		// This is where all the mods will be built for local testing
		ModBuildPathLocal: "build/mods",

		ModPath:                 modPath,
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

		ModIdPrefix: ModIdPrefix,

		// Current ck3 version that we have synced in "game" folder
		SyncedCk3Version: "1.10.0.1",

		ModFoldersToProcess: []string{
			"common",
			"events",
			"gfx",
			"gui",
			"history",
			"localization",
			"music",
		},

		MainMod: MainModType{
			Replacements: map[string]string{
				"modFolderName":        "mojo",
				"modName":              "Mojo Flavor by Antick",
				"modVersion":           "0.1.0",
				"supportedGameVersion": "1.10.0.1",
				"modRemoteFileId":      "",
				"modTags": `{
	"Alternative History"
	"Character Interactions"
	"Culture"
	"Decisions"
	"Events"
	"Fixes"
	"Gameplay"
	"Historical"
	"Utilities"
	"Warfare"
}`,
			},
		},

		SubMods: map[string]SubModType{
			"age-of-invasion": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_age_of_invasion",
					"modName":              "Age of Invasion",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "remote_file_id=\"2906586207\"",
					"modTags": `{
	"Gameplay"
	"Utilities"
	"Warfare"
}`,
				},
			},

			"grand-council": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_grand_council",
					"modName":              "Grand Council",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modTags": `{
	"Gameplay"
	"Utilities"
}`,
				},
			},

			"refurbished-titles": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_refurbished_titles",
					"modName":              "Refurbished Titles",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modTags": `{
	"Culture"
	"Fixes"
}`,
				},
			},

			"tweak-it": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_tweak_it",
					"modName":              "Tweak It",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modTags": `{
	"Fixes"
}`,
				},
			},

			"united-thrones": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_united_thrones",
					"modName":              "United Thrones",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modTags": `{
	"Fixes"
}`,
				},
			},

			"let-there-be-music": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_let_there_be_music",
					"modName":              "Let there be music",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modTags": `{
	"Utilities"
}`,
				},
			},

			"auto-pause-game": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_auto_pause_game",
					"modName":              "Auto Pause Game",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "remote_file_id=\"2906586207\"",
					"modTags": `{
	"Fixes"
	"Utilities"
}`,
				},
			},
		},
	}
}
