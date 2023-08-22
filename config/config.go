package config

import "path/filepath"

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
}

func Config() *Type {
	originalCk3Path := "/Users/pankaj/Library/Application Support/Steam/steamapps/common/Crusader Kings III/game"
	modCk3Path := "game"
	modPath := "mods"

	return &Type{
		// Put your original CK3 path here
		GameDataPath:         "/Users/pankaj/Paradox Interactive/Crusader Kings III",
		LauncherSettingsPath: "/Users/pankaj/Library/Application Support/Steam/steamapps/common/Crusader Kings III/launcher/launcher-settings.json",
		ModBuildPath:         "/Users/pankaj/Paradox Interactive/Crusader Kings III/mod/mojo-flavor",

		ModBuildPathLocal: "build",
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
		ModIdPrefix: "antick",
		ModVersion:  "0.1.0",

		// Current ck3 version that we have synced in "game" folder
		CurrentCk3Version: "1.9.2.1",

		// Supported ck3 version by the mods
		SupportedGameVersion: "1.9.2.1",

		ModInternalFolders: []string{
			"common",
			"events",
			"gfx",
			"gui",
			"history",
			"localization",
		},
	}
}
