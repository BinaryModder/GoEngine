package project

import (
	"encoding/json"
	"os"
)

func ReadProjectConfig(conf_path string) (*ProjectConfig, error) {

	data, err := os.ReadFile(conf_path)

	if err != nil {
		return &ProjectConfig{}, err
	}
	var project ProjectConfig

	err = json.Unmarshal(
		data,
		&project,
	)

	if err != nil {

		return &ProjectConfig{}, err
	}

	return &ProjectConfig{
		Name:          project.Name,
		EngineVersion: project.EngineVersion,
		Version:       project.Version,
		CreatedAt:     project.CreatedAt,
	}, nil

}
