package configs

import (
	"path/filepath"
	"strings"
)

type ModType struct {
	CombinedMod CombinedModType
	SubMods     map[string]SubModType
}

type CombinedModType struct {
	Replacements map[string]string
}

type SubModType struct {
	Enabled      bool
	Replacements map[string]string
}

var config = Config()

func ModConfig() *ModType {
	combinedModName := "mojo"
	modBuildPath := filepath.Join(config.GameCustomModPath, combinedModName)

	return &ModType{
		CombinedMod: CombinedModType{
			Replacements: map[string]string{
				"modFolderName":        combinedModName,
				"modName":              "Mojo by Antick",
				"modVersion":           "0.1.0",
				"supportedGameVersion": "1.11.0.1",
				"modRemoteFileId":      "",
				"modBuildPath":         strings.ReplaceAll(modBuildPath, "\\", "/"),
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
			"enter-the-warlord": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_enter_the_warlord",
					"modFolderName":        "enter-the-warlord",
					"modName":              "Enter the Warlord",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.16.2.3",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "enter-the-warlord"), "\\", "/"),
					"modTags": `{
	"Gameplay"
	"Utilities"
	"Warfare"
}`,
				},
			},

			"holdings-reimagined": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_holdings_reimagined",
					"modFolderName":        "holdings-reimagined",
					"modName":              "Holdings Reimagined",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.16.2.3",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "holdings-reimagined"), "\\", "/"),
					"modTags": `{
	"Balance"
	"Fixes"
	"Historical"
	"Graphics"
	"Gameplay"
}`,
				},
			},

			"grand-council": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_grand_council",
					"modFolderName":        "grand-council",
					"modName":              "Grand Council",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "grand-council"), "\\", "/"),
					"modTags": `{
	"Gameplay"
	"Utilities"
}`,
				},
			},

			"refurbished-titles": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_refurbished_titles",
					"modFolderName":        "refurbished-titles",
					"modName":              "Refurbished Titles",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "refurbished-titles"), "\\", "/"),
					"modTags": `{
	"Culture"
	"Fixes"
}`,
				},
			},

			"tweak-n-treat": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_tweak_n_treat",
					"modFolderName":        "tweak-n-treat",
					"modName":              "Tweak n Treat",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.14.2.2",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "tweak-n-treat"), "\\", "/"),
					"modTags": `{
	"Fixes"
}`,
				},
			},

			"united-thrones": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_united_thrones",
					"modFolderName":        "united-thrones",
					"modName":              "United Thrones",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "united-thrones"), "\\", "/"),
					"modTags": `{
	"Fixes"
}`,
				},
			},

			"let-there-be-music": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_let_there_be_music",
					"modFolderName":        "let-there-be-music",
					"modName":              "Let there be music",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.10.0.1",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "let-there-be-music"), "\\", "/"),
					"modTags": `{
	"Utilities"
}`,
				},
			},

			"auto-pause-game": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_auto_pause",
					"modFolderName":        "auto-pause-game",
					"modName":              "Auto Pause Game",
					"modVersion":           "1.2.0",
					"supportedGameVersion": "1.16.2.3",
					"modRemoteFileId":      "remote_file_id=\"2906586207\"",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "auto-pause-game"), "\\", "/"),
					"modTags": `{
	"Fixes"
	"Gameplay"
	"Utilities"
}`,
				},
			},

			"epe-barbershop-idle-pose-compatch": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_idle_pose_compatch",
					"modFolderName":        "epe-barbershop-idle-pose-compatch",
					"modName":              "EPE + Barbershop Idle Pose Compatch By Antick",
					"modVersion":           "1.1.1",
					"supportedGameVersion": "1.8.1",
					"modRemoteFileId":      "remote_file_id=\"2870777363\"",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "epe-barbershop-idle-pose-compatch"), "\\", "/"),
					"modTags": `{
	"Fixes"
	"Portraits"
}`,
				},
			},

			"traits-mania": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.ModIdPrefix + "_traits_mania",
					"modFolderName":        "traits-mania",
					"modName":              "Traits Mania",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.7.2",
					"modRemoteFileId":      "",
					"modBuildPath":         strings.ReplaceAll(filepath.Join(config.GameCustomModPath, "traits-mania"), "\\", "/"),
					"modTags": `{
	"Balance"
	"Decisions"
	"Gameplay"
	"Utilities"
}`,
				},
			},
		},
	}
}
