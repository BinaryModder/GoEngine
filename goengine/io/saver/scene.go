package saver

import (
	"encoding/json"
	"goengine/scene"
	"os"
)

func SaveToFile(s *scene.Scene, path string) error {
	data, err := json.MarshalIndent(s, "", "    ")

	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
