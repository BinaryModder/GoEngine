package hub

type Page int

const (
	PageProjects Page = iota
	PageSettings
)

type HubState struct {
	ShowCreateProject bool

	CurrentPage Page

	Projects []Project

	NewCreateName string

	NewCreatePath string

	ErrorMessage string
}

var (
	State HubState
)
