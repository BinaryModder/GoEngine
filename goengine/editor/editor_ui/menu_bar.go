package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/functions"
)

func MenuBar() giu.Widget {

	var selectedObjectIndex int32 = -1 // 0 - cube , 1 - pyramid , 2 - camera
	return giu.Row(
		//Left Part
		giu.Button("Save").OnClick(
			func() {
				if err := functions.SaveMenuBar(editor.EditState.CurrentScene, editor.EditState.ProjectPath); err != nil {
					return
				}
			},
		).Size(saveeditSizeWeight, saveeditSizeHeight),

		giu.Button("Edit").
			Size(saveeditSizeWeight, saveeditSizeHeight),

		giu.Button("Assets").OnClick(
			func() {
				if err := functions.AssetMenuBar(editor.EditState.DefaultAssetsFolder); err != nil {
					return
				}
			},
		),

		giu.Combo("", "SceneObj", []string{"Cube", "Pyramid", "Camera"}, &selectedObjectIndex).Size(120).
			OnChange(func() {
				if editor.EditState.CurrentScene == nil {
					return
				}

				if err := functions.SceneObjectMenuBar(editor.EditState.CurrentScene, &selectedObjectIndex); err != nil {
					return
				}

			}),

		giu.Button("Window"),

		giu.Button("Help"),

		giu.Dummy(toMiddleDummyWeight, toMiddleDummyHeight),

		//Middle Part
		giu.Button("Run").OnClick(
			func() {
				if err := functions.ProjectRunner(editor.EditState.ProjectPath); err != nil {
					return
				}
			},
		).Size(runSizeWeight, runSizeHeight),
	)
}
