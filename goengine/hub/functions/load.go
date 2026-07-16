package functions

import (
	"encoding/json"
	"github.com/sqweek/dialog"
	"goengine/project"
	"os"
	"path/filepath"
	"time"
)

func LoadProject() (project.Project, error) {

	path, err := dialog.Directory().
		Title("Choose GoEngine Project").
		Browse()

	if err != nil {
		return project.Project{}, nil
	}

	proj, err := readProject(path)

	if err != nil {
		return project.Project{}, nil
	}

	return proj, nil

}

func readProject(path string) (project.Project, error) {

	absPath := AbsolutePath(path)

	projectFile := filepath.Join(
		absPath,
		"ProjectSettings",
		"project.json",
	)

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
		Path:       absPath,
		CreatedAt:  config.CreatedAt,
		LastOpened: time.Now(),
	}, nil

}
