/**
 * This script allows you to update the gameDataPath in the
 * launcher-settings.json file in case you moved your game folder.
 */

package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func UpdateGameDataPath(filePath, newGameDataPath string) error {
	data, err := decodeJsonFile(filePath)
	if err != nil {
		return err
	}

	data["gameDataPath"] = newGameDataPath

	return encodeJsonFile(filePath, data)
}

func decodeJsonFile(filePath string) (map[string]interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Failed to close the file:", closeErr)
		}
	}()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func encodeJsonFile(filePath string, data map[string]interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Failed to close the file:", closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	return encoder.Encode(data)
}
