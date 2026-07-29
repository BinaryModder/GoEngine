package editor

import (
	"goengine/io/loader"
	"goengine/project"
	"goengine/scene"
)

type EditorState struct {
	ProjectPath string

	CurrentAssetsFolder string

	ProjectConfig *project.ProjectConfig

	CurrentScene *scene.Scene

	ProjectFiles []project.ProjectFile

	SelectedObject string

	Materials []scene.Material

	ErrorState string

	DefaultAssetsFolder string
}

func (s *EditorState) Init() error {
	scene, err := loader.LoadScene(State.ProjectPath)
	if err != nil {
		return err
	}

	State.CurrentScene = scene

	projectConfig, err := loader.LoadProjectConfig(State.ProjectPath)

	if err != nil {
		return err
	}
	State.ProjectConfig = projectConfig

	projectFiles, assetsPath, err := loader.LoadProjectFiles(State.ProjectPath)

	if err != nil {
		return err
	}

	State.DefaultAssetsFolder = assetsPath
	State.CurrentAssetsFolder = assetsPath
	State.ProjectFiles = projectFiles

	return nil

}

func (s *EditorState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *EditorState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}
