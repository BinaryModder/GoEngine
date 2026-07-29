package loader

import (
	"encoding/json"
	"goengine/scene"
	"os"
	"path/filepath"
)

// Loading current .scene file
func LoadScene(path string) (*scene.Scene, error) {

	scenePath := filepath.Join(
		path,
		"Assets",
		"Scenes",
		"Main.scene",
	)

	s, err := ReadScene(
		scenePath,
	)

	if err != nil {
		return &scene.Scene{}, err
	}

	return s, nil
}

func ReadScene(scene_path string) (*scene.Scene, error) {

	data, err := os.ReadFile(scene_path)

	if err != nil {
		return nil, err
	}

	var result_scene scene.Scene

	err = json.Unmarshal(
		data,
		&result_scene,
	)

	if err != nil {
		return nil, err
	}

	return &scene.Scene{
		Name:    result_scene.Name,
		Objects: result_scene.Objects,
	}, nil

}
