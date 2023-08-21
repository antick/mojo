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
					"modId": config.Config().ModIdPrefix + "_age_of_invasion",
				},
			},

			"grand-council": {
				Enabled: true,
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_grand_council",
				},
			},

			"refurbished-titles": {
				Enabled: true,
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_refurbished_titles",
				},
			},

			"tweak-it": {
				Enabled: true,
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_tweak_it",
				},
			},

			"united-thrones": {
				Enabled: true,
				Replacements: map[string]string{
					"modId": config.Config().ModIdPrefix + "_united_thrones",
				},
			},
		},
	}
}
