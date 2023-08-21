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
			"age_of_invasion": {
				Enabled: true,
				// TODO: Scan these folders automatically as these are common among all mods so specifying them manually should not be required.
				Mapping: map[string]string{
					"mods/age-of-invasion/common":       config.Config().ModBuildPath + "/common",
					"mods/age-of-invasion/events":       config.Config().ModBuildPath + "/events",
					"mods/age-of-invasion/gfx":          config.Config().ModBuildPath + "/gfx",
					"mods/age-of-invasion/localization": config.Config().ModBuildPath + "/localization",
				},
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_age_of_invasion",
				},
			},

			"grand_council": {
				Enabled: true,
				Mapping: map[string]string{
					"mods/grand-council/common":       config.Config().ModBuildPath + "/common",
					"mods/grand-council/gfx":          config.Config().ModBuildPath + "/gfx",
					"mods/grand-council/gui":          config.Config().ModBuildPath + "/gui",
					"mods/grand-council/localization": config.Config().ModBuildPath + "/localization",
				},
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_grand_council",
				},
			},

			"refurbished_titles": {
				Enabled: true,
				Mapping: map[string]string{
					"mods/refurbished-titles/common":       config.Config().ModBuildPath + "/common",
					"mods/refurbished-titles/localization": config.Config().ModBuildPath + "/localization",
				},
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_refurbished_titles",
				},
			},

			"tweak_it": {
				Enabled: true,
				Mapping: map[string]string{
					"mods/tweak-it/common": config.Config().ModBuildPath + "/common",
				},
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_tweak_it",
				},
			},

			"united_thrones": {
				Enabled: true,
				Mapping: map[string]string{
					"mods/united-thrones/common":       config.Config().ModBuildPath + "/common",
					"mods/united-thrones/localization": config.Config().ModBuildPath + "/localization",
				},
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_united_thrones",
				},
			},
		},
	}
}
