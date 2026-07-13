package editor

type EditorState struct {
	ProjectPath string

	Project ProjectConfig

	CurrentScene string

	//SelectedObject string

	//EditorCamera Camera
}

var State EditorState
