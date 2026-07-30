package functions

import (
	"goengine/engine/logger"
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
		logger.Error(err.Error())
		return
	}
	logger.Info("Folder was found")

	proj, err := loader.LoadProjectHUB(path)

	if err != nil {
		logger.Error(err.Error())
		return
	}
	if proj != (project.Project{}) {
		hub.State.Projects = append(hub.State.Projects, proj)
		logger.Info("Project is loaded")

	}

}
