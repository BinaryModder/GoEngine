package functions

import (
	"encoding/json"
	"github.com/sqweek/dialog"
	"goengine/hub"
	"os"
	"path/filepath"
	"time"
)

func LoadProject() (hub.Project, error) {

	path, err := dialog.Directory().
		Title("Choose GoEngine Project").
		Browse()

	if err != nil {
		return hub.Project{}, nil
	}

	project, err := readProject(path)

	if err != nil {
		return hub.Project{}, nil
	}

	return project, nil

}

func readProject(path string) (hub.Project, error) {

	absPath := AbsolutePath(path)

	projectFile := filepath.Join(
		absPath,
		"ProjectSettings",
		"project.json",
	)

	data, err := os.ReadFile(projectFile)

	if err != nil {
		return hub.Project{}, err
	}
	var config hub.ProjectConfig

	err = json.Unmarshal(
		data,
		&config,
	)

	if err != nil {

		return hub.Project{}, err
	}

	return hub.Project{
		Name:       config.Name,
		Path:       absPath,
		CreatedAt:  config.CreatedAt,
		LastOpened: time.Now(),
	}, nil

}
