package main

import (
	"flag"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/hub/ui"
	"log"
)

var (
	editorMode  bool
	projectPath string
)

func main() {
	flag.BoolVar(
		&editorMode,
		"editor",
		false,
		"Start editor",
	)

	flag.StringVar(
		&projectPath,
		"project",
		"",
		"Project path",
	)

	flag.Parse()

	if editorMode {

		if projectPath == "" {
			log.Fatal(
				"Project path is empty",
			)
		}

		editor.State.ProjectPath = projectPath

		err := editor.LoadScene()

		if err != nil {
			log.Fatal(err)
		}

		window := giu.NewMasterWindow(
			"GoEngine Editor",
			1200,
			800,
			0,
		)

		window.Run(
			editor.Loop,
		)

	} else {

		window := giu.NewMasterWindow(
			"GoEngine Hub",
			1050,
			750,
			giu.MasterWindowFlagsNotResizable,
		)

		window.Run(
			ui.Loop,
		)

	}

}
