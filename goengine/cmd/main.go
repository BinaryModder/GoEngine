package main

import (
	"flag"
	editor_ui "goengine/editor/ui"
	hub_ui "goengine/hub/ui"
	runtime_ui "goengine/runtime/ui"
)

var (
	EditorMode  bool
	RunMode     bool
	ProjectPath string
)

func main() {
	flag.BoolVar(
		&EditorMode,
		"editor",
		false,
		"Start editor",
	)

	flag.BoolVar(
		&RunMode,
		"runtime",
		false,
		"Start runtime",
	)

	flag.StringVar(
		&ProjectPath,
		"project",
		"",
		"Project path",
	)
	flag.Parse()
	if EditorMode {
		editor_ui.Run(ProjectPath)
	}
	if RunMode {
		runtime_ui.Run(ProjectPath)
	}
	if !EditorMode && !RunMode {
		hub_ui.Run()
	}

}
