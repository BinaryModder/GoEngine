package main

import (
	"flag"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/editor_ui"
	"goengine/hub/hub_ui"
	"log"
)

var (
	EditorMode  bool
	ProjectPath string
)

func main() {
	flag.BoolVar(
		&EditorMode,
		"editor",
		false,
		"Start editor",
	)

	flag.StringVar(
		&ProjectPath,
		"project",
		"",
		"Project path",
	)

	flag.Parse()

	if EditorMode {

		if ProjectPath == "" {
			log.Fatal(
				"Project path is empty",
			)
		}

		editor.State.ProjectPath = ProjectPath

		if err := editor.LoadWholeProject(); err != nil {
			log.Fatal(err)
		}

		window := giu.NewMasterWindow(
			"GoEngine Editor",
			1920,
			1080,
			0,
		)

		window.Run(
			editor_ui.Loop,
		)

	} else {

		window := giu.NewMasterWindow(
			"GoEngine Hub",
			1050,
			750,
			giu.MasterWindowFlagsNotResizable,
		)

		window.Run(
			hub_ui.Loop,
		)

	}

}
