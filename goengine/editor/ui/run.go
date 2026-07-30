package ui

import (
	"errors"

	"github.com/AllenDang/giu"
	"goengine/editor"
)

func Run(ProjectPath string) error {
	if ProjectPath == "" {
		errors.New(
			"Project path is empty",
		)
	}

	editor.State.ProjectPath = ProjectPath

	if err := editor.State.Init(); err != nil {
		return err
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
