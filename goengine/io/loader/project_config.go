package loader

import (
	"encoding/json"
	"goengine/project"
	"os"
	"path/filepath"
	"time"
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

// Reading the information of choosed project folder (project card format(HUB))
func LoadProjectHUB(path string) (project.Project, error) {

	projectFile := filepath.Join(
		path,
		"ProjectSettings",
		"project.json",
	)

	if _, err := os.Stat(projectFile); err != nil {
		return project.Project{}, nil
	}

	data, err := os.ReadFile(projectFile)

	if err != nil {
		return project.Project{}, err
	}
	var config project.ProjectConfig

	err = json.Unmarshal(
		data,
		&config,
	)

	if err != nil {

		return project.Project{}, err
	}

	return project.Project{
		Name:       config.Name,
		Path:       path,
		CreatedAt:  config.CreatedAt,
		LastOpened: time.Now(),
	}, nil

}
