package saver

import (
	"encoding/json"
	"goengine/settings"
	"os"
	"path/filepath"
)

func CreateSettings() error {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	goengineSettingsDir := filepath.Join(
		configDir,
		"GoEngine",
	)

	os.MkdirAll(goengineSettingsDir, os.ModePerm)

	default_data := settings.GoEngineSettings{
		Theme: "blue",
		Login: "guest",
	}

	fileData, err := json.MarshalIndent(default_data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(goengineSettingsDir, "settings.json"), fileData, 0644)
}
func SaveSettings() error {

	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	goengineSettingsFile := filepath.Join(
		configDir,
		"GoEngine",
		"settings.json",
	)

	fileData, err := json.MarshalIndent(settings.State, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(goengineSettingsFile, fileData, 0644)
}
