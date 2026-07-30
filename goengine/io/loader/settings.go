package loader

import (
	"encoding/json"
	"errors"
	"goengine/settings"
	"os"
	"path/filepath"
)

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

	var pure_data settings.GoEngineSettings

	if err := json.Unmarshal(fileData, &pure_data); err != nil {
		return err
	}

	settings.State = settings.GoEngineSettings{
		Login:   pure_data.Login,
		Theme:   pure_data.Theme,
		Console: pure_data.Console,
	}

	return nil
}
