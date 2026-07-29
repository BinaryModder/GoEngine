package runtime

import (
	"goengine/io/loader"
	"goengine/project"
	"goengine/scene"
)

type RuntimeState struct {
	ProjectPath   string
	CurrentScene  *scene.Scene
	ProjectConfig *project.ProjectConfig
}

func (s *RuntimeState) Init() error {
	state, err := loader.LoadScene(State.ProjectPath)

	if err != nil {
		return err
	}

	State.CurrentScene = state

	projectConf, err := loader.LoadProjectConfig(State.ProjectPath)

	if err != nil {
		return nil
	}

	State.ProjectConfig = projectConf

	return nil

}

func (s *RuntimeState) GetProjectPath() string {
	return s.ProjectPath
}
func (s *RuntimeState) GetProjectScene() *scene.Scene {
	return s.CurrentScene
}
