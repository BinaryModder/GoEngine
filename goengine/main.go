package main

import (
	"flag"

	"goengine/editor"
	"goengine/hub"

	"github.com/AllenDang/giu"
)

var editorMode bool
var projectPath string

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

		window := giu.NewMasterWindow(
			"GoEngine Editor",
			1200,
			800,
			0,
		)

		window.Run(
			func() {
				editor.Loop(projectPath)
			},
		)

	} else {

		window := giu.NewMasterWindow(
			"GoEngine Hub",
			800,
			600,
			0,
		)

		window.Run(
			hub.Loop,
		)

	}

}
