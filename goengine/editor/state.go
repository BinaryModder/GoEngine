package editor

import (
	"goengine/project"
	"goengine/scene"
)

type State interface {
	GetProjectPath() string
	GetProjectScene() *scene.Scene
}
type EditorState struct {
	ProjectPath string

	CurrentAssetsFolder string

	ProjectConfig *project.ProjectConfig

	CurrentScene *scene.Scene

	ProjectFiles []project.ProjectFile

	SelectedObject string

	ErrorState string

	DefaultAssetsFolder string
}

func (s *EditorState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *EditorState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}

type RunProjectState struct {
	ProjectPath   string
	CurrentScene  *scene.Scene
	ProjectConfig *project.ProjectConfig
}

func (s *RunProjectState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *RunProjectState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}

var EditState EditorState
var RunProjState RunProjectState
