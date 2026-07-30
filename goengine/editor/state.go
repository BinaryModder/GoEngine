package editor

import (
	"fmt"
	"goengine/engine/logger"
	"goengine/engine/platform"
	"goengine/io/loader"
	"goengine/project"
	"goengine/scene"
	"goengine/settings"
	"goengine/ui/layout"
)

type EditorState struct {
	ProjectPath string

	CurrentAssetsFolder string

	ProjectConfig *project.ProjectConfig

	CurrentScene *scene.Scene

	ProjectFiles []project.ProjectFile

	SelectedObject string

	Materials []scene.Material

	DefaultAssetsFolder string

	//Material creating states
	ShowCreateMaterial bool

	NewMaterialName       string
	NewMaterialSourcePath string

	//Loading material state
	ShowLoadMaterial bool

	LoadMaterialSourcePath string

	//Settings
	ShowConsole bool
}

func (s *EditorState) Init() error {
	//Platform information initializing
	platform.Init()
	logger.Info(fmt.Sprintf("Platform is initialized: %s", platform.State.OS))

	//Configuring sizes
	layout.ConfigureSize()
	logger.Info("Sizes are configured")

	// Scene initializing
	scene, err := loader.LoadScene(State.ProjectPath)
	if err != nil {
		return err
	}

	State.CurrentScene = scene

	logger.Info("Scene is loaded")

	projectConfig, err := loader.LoadProjectConfig(State.ProjectPath)

	if err != nil {
		return err
	}
	State.ProjectConfig = projectConfig

	logger.Info("Project config is loaded")

	projectFiles, assetsPath, err := loader.LoadProjectFiles(State.ProjectPath)

	if err != nil {
		return err
	}

	State.DefaultAssetsFolder = assetsPath
	State.CurrentAssetsFolder = assetsPath
	State.ProjectFiles = projectFiles

	logger.Info("Project files is loaded")

	projectMaterials, err := loader.LoadProjectMaterials(State.ProjectPath)

	if err != nil {
		return err
	}

	State.Materials = projectMaterials

	logger.Info("Materials are loaded")

	if err = loader.LoadSettings(); err != nil {
		return err
	}

	State.ShowConsole = settings.State.Console

	logger.Info("Settings are loaded")

	return nil

}

func (s *EditorState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *EditorState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}
