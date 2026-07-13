package scene

import (
	"encoding/json"
	"os"
)

func LoadScene(path string) (*Scene, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var scene Scene

	err = json.Unmarshal(
		data,
		&scene,
	)

	if err != nil {
		return nil, err
	}

	return &scene, nil

}
