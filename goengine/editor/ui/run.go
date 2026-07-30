package ui

import (
	"errors"
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/engine/logger"
)

func Run(ProjectPath string) error {
	if ProjectPath == "" {
		return errors.New(
			"Project path is empty",
		)
	}

	editor.State.ProjectPath = ProjectPath

	if err := editor.State.Init(); err != nil {
		logger.Error(fmt.Sprintf("Failed to initialize editor: %s", err))
	}

	window := giu.NewMasterWindow(
		"GoEngine Editor",
		1920,
		1080,
		0,
	)
	window.Run(
		Loop,
	)
	return nil

}
