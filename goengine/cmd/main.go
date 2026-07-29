package main

import (
	"flag"
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/editor_ui"
	"goengine/hub/hub_ui"
	"goengine/runtime"
	"goengine/ui/scale"
	"log"
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
		"run_proj",
		false,
		"Run Project",
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

		if err := editor.State.Init(); err != nil {
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

	}
	if RunMode {

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
			runtime.Loop,
		)

	}
	if !EditorMode && !RunMode {
		window := giu.NewMasterWindow(
			"GoEngine Hub",
			scale.I(1150),
			scale.I(725),
			giu.MasterWindowFlagsNotResizable,
		)

		window.Run(
			hub_ui.Loop,
		)

	}

}
