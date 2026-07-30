package functions

import (
	"goengine/hub"
	"goengine/io/dialog"
	"goengine/io/loader"
	"goengine/project"
)

// Loading exists project
func LoadProject() {
	var title string = "Choose GoEngine Project"
	path, err := dialog.ChooseProjectDialog(title)

	if err != nil {
		return
	}

	proj, err := loader.LoadProjectHUB(path)

	if err != nil {
		return
	}
	if proj != (project.Project{}) {
		hub.State.Projects = append(hub.State.Projects, proj)

	}

}
