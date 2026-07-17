package functions

import (
	"goengine/project"
	"goengine/scene"
	"path/filepath"
)

func LoadScene(path string) (*scene.Scene, error) {

	scenePath := filepath.Join(
		path,
		"Assets",
		"Scenes",
		"Main.scene",
	)

	s, err := scene.ReadScene(
		scenePath,
	)

	if err != nil {
		return &scene.Scene{}, err
	}

	return s, nil
}
func LoadProjectConfig(path string) (*project.ProjectConfig, error) {
	projectPath := filepath.Join(
		path,
		"ProjectSettings",
		"project.json",
	)

	p, err := project.ReadProjectConfig(
		projectPath,
	)

	if err != nil {
		return &project.ProjectConfig{}, err
	}

	return p, nil
}
func LoadProjectFiles(path string) ([]project.ProjectFile, string, error) {

	assetsPath := filepath.Join(
		path,
		"Assets",
	)

	files, _, err := LoadFolder(assetsPath)
	if err != nil {
		return []project.ProjectFile{}, "", err
	}

	return files, assetsPath, nil
}
