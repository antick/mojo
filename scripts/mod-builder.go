package scripts

import (
	"fmt"
	configuration "mojo/configs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

var config = configuration.Config()
var modConfig = configuration.ModConfig()

func Cleanup(modPath string) error {
	err := os.RemoveAll(modPath)
	if err != nil {
		fmt.Printf("Error while removing %s: %v \n", modPath, err)
		return err
	}

	fmt.Printf("🧹 Cleaned %s folder\n", modPath)
	return nil
}

func ProcessFile(modBuildPath, srcPath, destPath string, replacements map[string]string) error {
	destFile := filepath.Base(destPath)
	formatToReplace := "_\\$%s\\$_"

	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(fmt.Sprintf(formatToReplace, placeholder))
		destFile = re.ReplaceAllString(destFile, replacement)
	}

	finalDestPath := filepath.Join(filepath.Dir(destPath), destFile)

	if _, err := os.Stat(finalDestPath); err == nil {
		err := os.RemoveAll(finalDestPath)
		if err != nil {
			fmt.Println("Error while removing the file:", err)
			return err
		}
	}

	// Skip text replacement for certain file types
	if ext := filepath.Ext(srcPath); ext == ".dds" || ext == ".png" {
		return copyFiles(srcPath, finalDestPath)
	}

	file, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return err
	}

	data := string(file)

	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(fmt.Sprintf(formatToReplace, placeholder))
		data = re.ReplaceAllString(data, replacement)
	}

	// Ensure the destination directory exists
	if err := os.MkdirAll(modBuildPath, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(finalDestPath, []byte(data), 0644)
}

func copyFiles(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		fmt.Println("Error while copying the file:", err)
		return err
	}

	return os.WriteFile(dst, data, 0644)
}

func processDirectory(modBuildPath, srcDirectory, destDirectory string, replacements map[string]string) error {
	if err := os.MkdirAll(destDirectory, os.ModePerm); err != nil {
		return err
	}

	items, err := os.ReadDir(srcDirectory)
	if err != nil {
		return err
	}

	for _, item := range items {
		srcPath := filepath.Join(srcDirectory, item.Name())
		destPath := filepath.Join(destDirectory, item.Name())

		if item.IsDir() {
			if err = processDirectory(modBuildPath, srcPath, destPath, replacements); err != nil {
				return err
			}
		} else {
			if err := ProcessFile(modBuildPath, srcPath, destPath, replacements); err != nil {
				return err
			}
		}
	}

	return nil
}

func BaseReplacements() map[string]string {
	textReplacement := map[string]string{}

	return textReplacement
}

func mergeReplacements(map1, map2 map[string]string) map[string]string {
	mergedMap := make(map[string]string)
	for k, v := range map1 {
		mergedMap[k] = v
	}

	for k, v := range map2 {
		mergedMap[k] = v
	}

	return mergedMap
}

func Build(modBuildPath, modName string, replacements map[string]string) error {
	var filesAndFolderMapping = make(map[string]string)

	for _, value := range config.ModFoldersToProcess {
		filesAndFolderMapping[filepath.Join(config.ModPath, modName, value)] = filepath.Join(modBuildPath, value)
	}

	textReplacement := mergeReplacements(replacements, BaseReplacements())

	for srcPath, destPath := range filesAndFolderMapping {
		info, err := os.Stat(srcPath)
		if err != nil {
			if os.IsNotExist(err) {
				//fmt.Printf("Skipping %s...\n", srcPath)
				continue
			} else {
				fmt.Printf("Build failed while accessing the given path: %v \n", err)
			}
		}

		if info.IsDir() {
			if err = processDirectory(modBuildPath, srcPath, destPath, textReplacement); err != nil {
				fmt.Printf("Build failed for the given directory: %v \n", err)
				return nil
			}
		} else {
			if err := ProcessFile(modBuildPath, srcPath, destPath, textReplacement); err != nil {
				fmt.Printf("Build failed: %v \n", err)
				return nil
			}
		}
	}

	return nil
}

func SortModList(subMods map[string]configuration.SubModType) []string {
	keys := make([]string, 0, len(subMods))
	for k := range subMods {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func BuildCombinedDescriptorFile(modBuildPath string) error {
	err := ProcessFile(
		modBuildPath,
		config.ModDescriptorSourcePath,
		filepath.Join(modBuildPath, "descriptor.mod"),
		modConfig.CombinedMod.Replacements,
	)

	return err
}

func BuildLooseDescriptorFiles(modBuildPath string, replacements map[string]string) error {
	err := ProcessFile(
		modBuildPath,
		config.ModDescriptorSourcePath,
		filepath.Join(modBuildPath, "descriptor.mod"),
		replacements,
	)

	return err
}

func BuildModFile(modBuildPath, modFileName string) error {
	sourceModFilePath := filepath.Join(config.ModPath, modConfig.CombinedMod.Replacements["modFolderName"]+".mod")
	err := ProcessFile(
		modBuildPath,
		sourceModFilePath,
		filepath.Join(modBuildPath, modFileName),
		modConfig.CombinedMod.Replacements,
	)

	return err
}

func BuildLooseModFiles(modBuildPath string, modFileName string, replacements map[string]string) error {
	sourceModFilePath := filepath.Join(config.ModPath, modConfig.CombinedMod.Replacements["modFolderName"]+".mod")
	err := ProcessFile(
		modBuildPath,
		sourceModFilePath,
		filepath.Join(modBuildPath, modFileName),
		replacements,
	)

	return err
}

func BuildCombinedThumbnailFile(modBuildPath string) error {
	err := ProcessFile(
		modBuildPath,
		config.ThumbnailSourcePath,
		filepath.Join(modBuildPath, "thumbnail.png"),
		map[string]string{},
	)

	return err
}

func BuildLooseThumbnailFiles(modBuildPath string, modFolderName string) error {
	err := ProcessFile(
		modBuildPath,
		filepath.Join(config.ModPath, modFolderName, "thumbnail.png"),
		filepath.Join(modBuildPath, "thumbnail.png"),
		map[string]string{},
	)

	return err
}

func BuildCombinedMod(modBuildPath string, selectedModKeys []string) error {
	fmt.Println("-----------------------------------")
	err := Cleanup(modBuildPath)
	if err != nil {
		return err
	}

	if err = BuildCombinedDescriptorFile(modBuildPath); err != nil {
		fmt.Println("Error while processing descriptor.mod file")
		return err
	}

	if err = BuildCombinedThumbnailFile(modBuildPath); err != nil {
		fmt.Println("Error while processing thumbnail.png file")
		return err
	}

	//sortedModList := SortModList(modConfig.SubMods)
	//for _, modName := range sortedModList {
	//	modDetails := modConfig.SubMods[modName]
	//
	//	if !modDetails.Enabled {
	//		fmt.Printf("❗️%s is disabled, skipping \n", modName)
	//		continue
	//	} else {
	//		fmt.Printf("📦 Building %s\n", modName)
	//	}
	//
	//	err := Build(
	//		modBuildPath,
	//		modName,
	//		modDetails.Replacements,
	//	)
	//	if err != nil {
	//		return err
	//	}
	//}

	for _, modKey := range selectedModKeys {
		modDetails, exists := modConfig.SubMods[modKey]
		if !exists {
			fmt.Printf("Mod %s not found\n", modKey)
			continue
		}

		if !modDetails.Enabled {
			fmt.Printf("❗️%s is disabled, skipping \n", modKey)
			continue
		} else {
			fmt.Printf("📦 Building %s\n", modKey)
		}

		err := Build(
			modBuildPath,
			modKey,
			modDetails.Replacements,
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Printf("✅ Build successful in %s folder \n", modBuildPath)
	fmt.Println("-----------------------------------")

	return nil
}

func BuildLooseMods(buildPath string, selectedModKeys []string) error {
	fmt.Println("-----------------------------------")
	for _, modKey := range selectedModKeys {
		modDetails, exists := modConfig.SubMods[modKey]
		if !exists {
			fmt.Printf("Mod %s not found\n", modKey)
			continue
		}

		if !modDetails.Enabled {
			fmt.Printf("❗️%s is disabled, skipping \n", modKey)
			continue
		} else {
			fmt.Printf("📦 Building %s\n", modKey)
		}

		modFolderName := modDetails.Replacements["modFolderName"]
		modBuildPath := filepath.Join(buildPath, modFolderName)
		modFileName := modFolderName + ".mod"

		fmt.Println("-----------------------------------")
		err := Cleanup(modBuildPath)
		if err != nil {
			return err
		}

		if err := BuildLooseModFiles(buildPath, modFileName, modDetails.Replacements); err != nil {
			fmt.Println("Error while processing mojo.mod file")
			return err
		}

		if err = BuildLooseDescriptorFiles(modBuildPath, modDetails.Replacements); err != nil {
			fmt.Println("Error while processing descriptor.mod file")
			return err
		}

		if err = BuildLooseThumbnailFiles(modBuildPath, modFolderName); err != nil {
			fmt.Println("Error while processing thumbnail.png file")
			return err
		}

		err = Build(
			modBuildPath,
			modKey,
			modDetails.Replacements,
		)
		if err != nil {
			return err
		}

		fmt.Printf("✅ Build successful in %s folder \n", modBuildPath)
	}
	fmt.Println("-----------------------------------")

	return nil
}
