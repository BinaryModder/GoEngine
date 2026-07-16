package scene

import (
	"encoding/json"
	"os"
)

func ReadScene(scene_path string) (*Scene, error) {

	data, err := os.ReadFile(scene_path)

	if err != nil {
		return nil, err
	}

	var result_scene Scene

	err = json.Unmarshal(
		data,
		&result_scene,
	)

	if err != nil {
		return nil, err
	}

	return &Scene{
		Name:    result_scene.Name,
		Objects: result_scene.Objects,
	}, nil

}
