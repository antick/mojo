package scripts

import (
	"fmt"
	configuration "mojo/configs"
	"os"
	"path/filepath"
	"regexp"
)

var config = configuration.Config()

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

	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(fmt.Sprintf("{{%s}}", placeholder))
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
		re := regexp.MustCompile(fmt.Sprintf("{{%s}}", placeholder))
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