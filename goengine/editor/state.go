package editor

import (
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

	ErrorState string

	DefaultAssetsFolder string
}

var State EditorState
