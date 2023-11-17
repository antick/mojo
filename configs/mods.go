package configs

import (
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
	return &ModType{
		CombinedMod: CombinedModType{
			Replacements: map[string]string{
				"modFolderName":        "mojo",
				"modName":              "Mojo by Antick",
				"modVersion":           "0.1.0",
				"supportedGameVersion": "1.11.0.1",
				"modRemoteFileId":      "",
				"modBuildPath":         strings.ReplaceAll(config.ModBuildPath, "\\", "/"),
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
					"modId":                config.ModIdPrefix + "_age_of_invasion",
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
					"modId":                config.ModIdPrefix + "_grand_council",
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
					"modId":                config.ModIdPrefix + "_refurbished_titles",
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
					"modId":                config.ModIdPrefix + "_tweak_it",
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
					"modId":                config.ModIdPrefix + "_united_thrones",
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
					"modId":                config.ModIdPrefix + "_let_there_be_music",
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
					"modId":                config.ModIdPrefix + "_auto_pause",
					"modName":              "Auto Pause Game",
					"modVersion":           "1.0.0",
					"supportedGameVersion": "1.11.0.1",
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
