package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	return os.UserHomeDir()

}

func Read() (Config, error) {
	dir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	filePath := filepath.Join(dir, configFileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil

}
