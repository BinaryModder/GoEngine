package editor

import (
	"goengine/project"
	"goengine/scene"
)

type EditorState struct {
	ProjectPath string

	ProjectConfig *project.ProjectConfig

	CurrentScene *scene.Scene

	//SelectedObject string

	//EditorCamera Camera

	ErrorState string
}

var State EditorState
