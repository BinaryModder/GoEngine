package hub

import (
	"goengine/project"
)

type Page int

const (
	PageProjects Page = iota
	PageSettings
)

type HubState struct {
	ShowCreateProject bool

	CurrentPage Page

	Projects []project.Project

	NewCreateName string

	NewCreatePath string

	ErrorMessage string
}

var (
	State HubState
)
