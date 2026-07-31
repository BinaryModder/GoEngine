package runtime

import (
	"fmt"
	"goengine/engine/logger"
	"goengine/io/loader"
	"goengine/project"
	"goengine/scene"
	"goengine/settings"
)

type RuntimeState struct {
	ProjectPath   string
	CurrentScene  *scene.Scene
	ProjectConfig *project.ProjectConfig
	Materials     []scene.Material

	//Settings
	ShowConsole bool
}

func (s *RuntimeState) Init() error {

	scene, err := loader.LoadScene(State.ProjectPath)

	if err != nil {
		return err
	}

	State.CurrentScene = scene

	logger.Info("Scene is loaded")

	projectConf, err := loader.LoadProjectConfig(State.ProjectPath)

	if err != nil {
		return err
	}

	State.ProjectConfig = projectConf

	logger.Info("Project Configuration file is loaded")

	projectMaterials, err := loader.LoadProjectMaterials(State.ProjectPath)

	if err != nil {
		return err
	}

	State.Materials = projectMaterials

	logger.Info("Materials is loaded")

	if err := loader.LoadSettings(); err != nil {
		logger.Error(fmt.Sprintf("Failed to load settings: %s", err.Error()))
	}
	State.ShowConsole = settings.State.Console

	logger.Info("Settings are loaded")
	return nil

}

func (s *RuntimeState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *RuntimeState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}
