package runtime

import (
	"fmt"

	"goengine/engine/logger"
	"goengine/engine/platform"
	"goengine/engine/renderer"
	enginescript "goengine/engine/script"
)

func Run(projectPath string) {

	State.ProjectPath = projectPath

	if err := State.Init(); err != nil {
		logger.Error(err.Error())
		return
	}

	platform.Init()

	if err := CreateWindow(
		fmt.Sprintf(
			"GoEngine Runtime %s",
			State.ProjectConfig.Version,
		),
		1920,
		1080,
	); err != nil {

		logger.Error(err.Error())
		return
	}

	if err := renderer.Init(int32(Window.Width), int32(Window.Height)); err != nil {
		logger.Error(fmt.Sprintf("Failed to initialize renderer : %v", err))
		return
	}

	renderer.SetRuntimeRendererMode(true)
	renderer.SetWinSize(Window.Width, Window.Height)
	renderer.SetUseSceneCamera(true)

	logger.Info("Renderer succesfuly initialized")

	enginescript.State.Init(State.ProjectPath)
	enginescript.State.LoadAndStartAll(State.CurrentScene)

	logger.Info("Runtime started")

	Loop()
}
