package hub

import (
	"goengine/project"
)

type Page int

const (
	PageProjects Page = iota //0
	PageSettings             //1
)

type HubState struct {
	ShowCreateProject bool

	CurrentPage Page // Current Page on Interface (for switching)

	Projects []project.Project // List of loaded project (for project cards)

	NewCreateName string // Creating project part. Just storing name of new project (from InputText)

	NewCreatePath string // Creating project part. Just storing  of new project(from InputText)

	SaveSettingsShowButton bool // State for "Save" button

	ErrorMessage string //Should show error information
}

var (
	State HubState
)
