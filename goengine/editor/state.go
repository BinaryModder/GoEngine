package editor

import (
	"goengine/project"
	"goengine/scene"
)

type EditorState struct {
	ProjectPath string

	AssetsFolder string

	ProjectConfig *project.ProjectConfig

	CurrentScene *scene.Scene

	ProjectFiles []project.ProjectFile

	//SelectedObject string

	//EditorCamera Camera

	ErrorState string
}

var State EditorState
