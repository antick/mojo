package tools

import (
	"fmt"
	configuration "mojo/config"
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

func ProcessFile(srcPath, destPath string, replacements map[string]string) error {
	destFile := filepath.Base(destPath)
	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(fmt.Sprintf("{{%s}}", placeholder))
		destFile = re.ReplaceAllString(destFile, replacement)
	}

	finalDestPath := filepath.Join(filepath.Dir(destPath), destFile)

	if _, err := os.Stat(finalDestPath); err == nil {
		err := os.RemoveAll(finalDestPath)
		if err != nil {
			return err
		}
	}

	if filepath.Ext(srcPath) == ".dds" {
		return copyFiles(srcPath, finalDestPath)
	}

	file, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	data := string(file)

	for placeholder, replacement := range replacements {
		re := regexp.MustCompile(fmt.Sprintf("{{%s}}", placeholder))
		data = re.ReplaceAllString(data, replacement)
	}

	return os.WriteFile(finalDestPath, []byte(data), 0644)
}

func copyFiles(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, 0644)
}

func processDirectory(srcDirectory, destDirectory string, replacements map[string]string) error {
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
			if err = processDirectory(srcPath, destPath, replacements); err != nil {
				return err
			}
		} else {
			if err := ProcessFile(srcPath, destPath, replacements); err != nil {
				return err
			}
		}
	}

	return nil
}

func Build(modName string, replacements map[string]string) error {
	var filesAndFolderMapping = make(map[string]string)
	filesAndFolderMapping[config.ModDescriptorPath] = filepath.Join(config.ModBuildPath, "descriptor.mod")

	for _, value := range config.ModInternalFolders {
		filesAndFolderMapping[filepath.Join(config.ModPath, modName, value)] = filepath.Join(config.ModBuildPath, value)
	}

	textReplacement := replacements
	textReplacement["modVersion"] = config.ModVersion
	textReplacement["supportedGameVersion"] = config.SupportedGameVersion

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
			if err = processDirectory(srcPath, destPath, textReplacement); err != nil {
				fmt.Printf("Build failed for the given directory: %v \n", err)
				return nil
			}
		} else {
			if err := ProcessFile(srcPath, destPath, textReplacement); err != nil {
				fmt.Printf("Build failed: %v \n", err)
				return nil
			}
		}
	}

	return nil
}
