package main

import (
	"flag"
	editor_ui "goengine/editor/ui"
	"goengine/engine/logger"
	hub_ui "goengine/hub/ui"
	runtime_ui "goengine/runtime/ui"
	"log"
)

var (
	EditorMode  bool
	RunMode     bool
	ProjectPath string
)

func main() {

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("GoEngine started")

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
