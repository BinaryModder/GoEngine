package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/runtime"
	"log"
)

func Run(ProjectPath string) {
	if ProjectPath == "" {
		log.Fatal(
			"Project path is empty",
		)
	}

	runtime.State.ProjectPath = ProjectPath

	if err := runtime.State.Init(); err != nil {
		log.Fatal(err)
	}

	window := giu.NewMasterWindow(
		fmt.Sprintf(
			"Version: %s ; EngineVersion: %s",
			runtime.State.ProjectConfig.Version,
			runtime.State.ProjectConfig.EngineVersion,
		),
		1920,
		1080,
		0,
	)
	window.Run(
		Loop,
	)
}
