package config

import (
	"os"
	"path/filepath"
)

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
	homeDir, _ := os.UserHomeDir()

	originalCk3Path := filepath.Join(homeDir, "Library", "Application Support", "Steam", "steamapps", "common", "Crusader Kings III", "game")
	modCk3Path := "game"
	modPath := "mods"

	return &Type{
		// Put your original CK3 path here
		GameDataPath:         filepath.Join(homeDir, "Paradox Interactive", "Crusader Kings III"),
		LauncherSettingsPath: filepath.Join(homeDir, "Library", "Application Support", "Steam", "steamapps", "common", "Crusader Kings III", "launcher", "launcher-settings.json"),
		ModBuildPath:         filepath.Join(homeDir, "Paradox Interactive", "Crusader Kings III", "mod", "mojo-flavor"),

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
	}
}
