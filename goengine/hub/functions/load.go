package functions

import (
	"encoding/json"
	"fmt"
	"github.com/sqweek/dialog"
	"goengine/hub"
	"os"
	"os/exec"
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

	OpenEditor(path)

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

func OpenEditor(path string) {

	exePath, err := os.Executable()

	if err != nil {

		fmt.Println(err)

		return

	}

	cmd := exec.Command(

		exePath,

		"-editor",

		"-project",

		path,
	)

	cmd.Start()

}
