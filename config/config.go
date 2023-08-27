package config

import (
	"os"
	"path/filepath"
)

type Mod struct {
	Enabled      bool              `json:"enabled"`
	Mapping      map[string]string `json:"mapping"`
	Replacements map[string]string `json:"replacements"`
}

type Type struct {
	GameDataPath         string            `json:"gameDataPath"`
	LauncherSettingsPath string            `json:"launcherSettingsPath"`
	ModBuildPath         string            `json:"modBuildPath"`
	ModBuildPathLocal    string            `json:"modBuildPathLocal"`
	ModPath              string            `json:"modPath"`
	ModDescriptorPath    string            `json:"modDescriptorPath"`
	OriginalCk3Path      string            `json:"originalCk3Path"`
	ModCk3Path           string            `json:"modCk3Path"`
	Ck3PullMapping       map[string]string `json:"ck3PullMapping"`
	ModIdPrefix          string            `json:"modIdPrefix"`
	ModVersion           string            `json:"modVersion"`
	CurrentCk3Version    string            `json:"currentCk3Version"`
	SupportedGameVersion string            `json:"supportedGameVersion"`
	ModInternalFolders   []string          `json:"modInternalFolders"`
	Mods                 map[string]Mod    `json:"mods"`
}

func Config() *Type {
	homeDir, _ := os.UserHomeDir()

	originalCk3Path := filepath.Join(homeDir, "Library", "Application Support", "Steam", "steamapps", "common", "Crusader Kings III", "game")
	modCk3Path := "game"
	modPath := "mods"
	ModIdPrefix := "antick"

	return &Type{
		// Put your original CK3 path here
		GameDataPath:         filepath.Join(homeDir, "Paradox Interactive", "Crusader Kings III"),
		LauncherSettingsPath: filepath.Join(homeDir, "Library", "Application Support", "Steam", "steamapps", "common", "Crusader Kings III", "launcher", "launcher-settings.json"),
		ModBuildPath:         filepath.Join(homeDir, "Paradox Interactive", "Crusader Kings III", "mod", "mojo-flavor"),

		// This is where all the mods will be built
		ModBuildPathLocal: "mod-build",

		ModPath:           modPath,
		ModDescriptorPath: filepath.Join(modPath, "descriptor.mod"),
		OriginalCk3Path:   originalCk3Path,
		ModCk3Path:        modCk3Path,
		Ck3PullMapping: map[string]string{
			filepath.Join(originalCk3Path, "common"):               filepath.Join(modCk3Path, "common"),
			filepath.Join(originalCk3Path, "events"):               filepath.Join(modCk3Path, "events"),
			filepath.Join(originalCk3Path, "gui"):                  filepath.Join(modCk3Path, "gui"),
			filepath.Join(originalCk3Path, "history"):              filepath.Join(modCk3Path, "history"),
			filepath.Join(originalCk3Path, "localization/english"): filepath.Join(modCk3Path, "localization/english"),
		},
		ModIdPrefix: ModIdPrefix,
		ModVersion:  "0.1.0",

		// Current ck3 version that we have synced in "game" folder
		CurrentCk3Version: "1.10.0.1",

		// Supported ck3 version by the mods
		SupportedGameVersion: "1.10.0.1",

		ModInternalFolders: []string{
			"common",
			"events",
			"gfx",
			"gui",
			"history",
			"localization",
			"music",
		},

		Mods: map[string]Mod{
			"age-of-invasion": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_age_of_invasion",
					"modName":              "Age of Invasion",
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "remote_file_id=\"2906586207\"",
					"tags": `{
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
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "",
					"tags": `{
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
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "",
					"tags": `{
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
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "",
					"tags": `{
	"Fixes"
}`,
				},
			},

			"united-thrones": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_united_thrones",
					"modName":              "United Thrones",
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "",
					"tags": `{
	"Fixes"
}`,
				},
			},

			"let-there-be-music": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_let_there_be_music",
					"modName":              "Let there be music",
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "",
					"tags": `{
	"Utilities"
}`,
				},
			},

			"auto-pause-game": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                ModIdPrefix + "_auto_pause_game",
					"modName":              "Auto Pause Game",
					"version":              "1.0.0",
					"supportedGameVersion": "{{supportedGameVersion}}",
					"remoteFileId":         "remote_file_id=\"2906586207\"",
					"tags": `{
	"Fixes"
	"Utilities"
}`,
				},
			},
		},
	}
}
