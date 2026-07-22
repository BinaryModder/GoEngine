package settings

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type GoEngineSettings struct {
	Login string `json:"login"`
	Theme string `json:"theme"`
}

var State GoEngineSettings

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

	default_data := GoEngineSettings{
		Theme: "blue",
		Login: "guest",
	}

	fileData, err := json.MarshalIndent(default_data, "", "    ")
	return os.WriteFile(filepath.Join(goengineSettingsDir, "settings.json"), fileData, 0644)
}
func LoadSettings() error {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	goengineSettingsFile := filepath.Join(
		configDir,
		"GoEngine",
		"settings.json",
	)

	if _, err := os.Stat(goengineSettingsFile); os.IsNotExist(err) {
		return errors.New("Settings file does not exists")
	}

	fileData, err := os.ReadFile(goengineSettingsFile)

	var pure_data GoEngineSettings

	if err := json.Unmarshal(fileData, &pure_data); err != nil {
		return err
	}

	State = GoEngineSettings{
		Login: pure_data.Login,
		Theme: pure_data.Theme,
	}

	return nil
}

func SaveSettings(setting *GoEngineSettings) error {

	configDir, err := os.UserConfigDir()

	if err != nil {
		return err
	}

	goengineSettingsFile := filepath.Join(
		configDir,
		"GoEngine",
		"settings.json",
	)

	fileData, err := json.MarshalIndent(State, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(goengineSettingsFile, fileData, 0644)
}
