package editor

import (
	"goengine/project"
	"goengine/scene"
	"path/filepath"
)

func LoadScene() error {

	scenePath := filepath.Join(
		State.ProjectPath,
		"Assets",
		"Scenes",
		"Main.scene",
	)

	s, err := scene.ReadScene(
		scenePath,
	)

	if err != nil {
		return err
	}

	State.CurrentScene = s

	return nil
}
func LoadProjectFile() error {
	projectPath := filepath.Join(
		State.ProjectPath,
		"ProjectSettings",
		"project.json",
	)

	p, err := project.ReadProjectConfig(
		projectPath,
	)

	if err != nil {
		return err
	}

	State.ProjectConfig = p
	return nil
}
