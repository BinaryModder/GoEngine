package editor_ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/functions"
)

func MenuBar() giu.Widget {

	return giu.Row(
		//Left Part
		giu.Button("Save").OnClick(
			func() {
				if err := functions.SaveProject(editor.State.CurrentScene, editor.State.ProjectPath); err != nil {
					return
				}
			},
		).Size(saveeditSizeWeight, saveeditSizeHeight),

		giu.Button("Edit").
			Size(saveeditSizeWeight, saveeditSizeHeight),

		giu.Button("Assets").OnClick(
			func() {
				if err := functions.AssetMenuBar(editor.State.DefaultAssetsFolder); err != nil {
					return
				}
			},
		),

		giu.Button("GameObject"),

		giu.Button("Window"),

		giu.Button("Help"),

		giu.Dummy(toMiddleDummyWeight, toMiddleDummyHeight),

		//Middle Part
		giu.Button("Run").OnClick(
			func() {
				fmt.Println("Running project...")
			},
		).Size(runSizeWeight, runSizeHeight),
	)
}
