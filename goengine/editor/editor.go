package editor

import (
	"goengine/editor/functions"
)

func LoadEditorModeProject() error {

	state, err := functions.LoadScene(EditState.ProjectPath)
	if err != nil {
		return err
	}

	EditState.CurrentScene = state

	projectConfig, err := functions.LoadProjectConfig(EditState.ProjectPath)

	if err != nil {
		return err
	}
	EditState.ProjectConfig = projectConfig

	projectFiles, assetsPath, err := functions.LoadProjectFiles(EditState.ProjectPath)

	if err != nil {
		return err
	}

	EditState.DefaultAssetsFolder = assetsPath
	EditState.CurrentAssetsFolder = assetsPath
	EditState.ProjectFiles = projectFiles

	return nil

}

func LoadRunModeProject() error {

	state, err := functions.LoadScene(RunProjState.ProjectPath)

	if err != nil {
		return err
	}

	RunProjState.CurrentScene = state

	projectConf, err := functions.LoadProjectConfig(RunProjState.ProjectPath)

	if err != nil {
		return nil
	}

	RunProjState.ProjectConfig = projectConf

	return nil

}
