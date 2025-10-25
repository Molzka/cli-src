package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"src/internal/utils"
	"time"
)

type TokenData struct {
	Token     string `json:"token"`
	Timestamp string `json:"timestamp"`
}

func SaveToken(token string) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	tokenData := TokenData{
		Token:     token,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	data, err := json.MarshalIndent(tokenData, "", "  ")
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "token.json")
	return os.WriteFile(configFile, data, 0600)
}

func getConfigDir() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("APPDATA"), "src-cli")
	case "darwin":
		configDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "src-cli")
	default:
		xdgConfig := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfig == "" {
			xdgConfig = filepath.Join(os.Getenv("HOME"), ".config")
		}
		configDir = filepath.Join(xdgConfig, "src-cli")
	}

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}

	return configDir, nil
}

func LoadToken() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}

	configFile := filepath.Join(configDir, "token.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	var tokenData TokenData
	if err := json.Unmarshal(data, &tokenData); err != nil {
		return "", err
	}

	return tokenData.Token, nil
}

func DeleteToken() error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "token.json")
	err = os.Remove(configFile)
	if err != nil {
		return err
	}

	isEmpty, err := utils.IsEmpty(configDir)

	if err != nil {
		return err
	}

	if isEmpty {
		err = os.Remove(configDir)
		if err != nil {
			return err
		}
	}

	return nil
}
