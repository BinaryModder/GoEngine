package loader

import (
	"encoding/json"
	"goengine/project"
	"os"
	"path/filepath"
)

// Loading metadata of ProjectConfig file
func LoadProjectConfig(path string) (*project.ProjectConfig, error) {
	projectPath := filepath.Join(
		path,
		"ProjectSettings",
		"project.json",
	)

	p, err := ReadProjectConfig(
		projectPath,
	)

	if err != nil {
		return &project.ProjectConfig{}, err
	}

	return p, nil
}

func ReadProjectConfig(conf_path string) (*project.ProjectConfig, error) {

	data, err := os.ReadFile(conf_path)
	if err != nil {
		return &project.ProjectConfig{}, err
	}
	var config project.ProjectConfig

	err = json.Unmarshal(
		data,
		&config,
	)

	if err != nil {

		return nil, err
	}

	return &project.ProjectConfig{
		Name:          config.Name,
		EngineVersion: config.EngineVersion,
		Version:       config.Version,
		CreatedAt:     config.CreatedAt,
	}, nil

}
