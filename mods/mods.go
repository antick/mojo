package mods

import "mojo/config"

type Mod struct {
	Enabled      bool              `json:"enabled"`
	Mapping      map[string]string `json:"mapping"`
	Replacements map[string]string `json:"replacements"`
}

type Type struct {
	Mods map[string]Mod `json:"mods"`
}

func ModConfig() *Type {
	return &Type{
		Mods: map[string]Mod{
			"age-of-invasion": {
				Enabled: true,
				Replacements: map[string]string{
					"modId":                config.Config().ModIdPrefix + "_age_of_invasion",
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
					"modId":                config.Config().ModIdPrefix + "_grand_council",
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
					"modId":                config.Config().ModIdPrefix + "_refurbished_titles",
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
					"modId":                config.Config().ModIdPrefix + "_tweak_it",
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
					"modId":                config.Config().ModIdPrefix + "_united_thrones",
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
					"modId":                config.Config().ModIdPrefix + "_let_there_be_music",
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
					"modId":                config.Config().ModIdPrefix + "_auto_pause_game",
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
