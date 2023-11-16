/**
 * This script pulls original CK3 game folder files here
 * for tracking the new updates in the game
 */

package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var allowedExtensions = []string{".txt", ".info", ".gui", ".yml"}

func Pull(mapping map[string]string) error {
	for srcPath, destPath := range mapping {
		err := processCk3Directory(srcPath, destPath)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	return nil
}

func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}

	return false
}

func copyFile(srcPath, destPath string) error {
	if contains(allowedExtensions, filepath.Ext(srcPath)) {
		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}

		defer func() {
			closeErr := srcFile.Close()
			if closeErr != nil {
				fmt.Println("Failed to close the file:", closeErr)
			}
		}()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}

		defer func() {
			closeErr := destFile.Close()
			if closeErr != nil {
				fmt.Println("Failed to close the file:", closeErr)
			}
		}()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func processCk3Directory(srcDirectory, destDirectory string) error {
	// Remove the destination directory if it exists, then recreate it
	if _, err := os.Stat(destDirectory); !os.IsNotExist(err) {
		err := os.RemoveAll(destDirectory)
		if err != nil {
			return err
		}
	}

	// Ensure the destination directory exists
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
			// If the item is a directory, process it recursively
			err := processCk3Directory(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// If the item is a file, process it directly
			err := copyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
