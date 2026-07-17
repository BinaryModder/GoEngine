package editor

import (
	"goengine/editor/functions"
)

func LoadWholeProject() error {

	state, err := functions.LoadScene(State.ProjectPath)
	if err != nil {
		return err
	}

	State.CurrentScene = state

	projectConfig, err := functions.LoadProjectConfig(State.ProjectPath)

	if err != nil {
		return err
	}
	State.ProjectConfig = projectConfig

	projectFiles, assetsPath, err := functions.LoadProjectFiles(State.ProjectPath)

	if err != nil {
		return err
	}

	State.AssetsFolder = assetsPath
	State.ProjectFiles = projectFiles

	return nil

}
