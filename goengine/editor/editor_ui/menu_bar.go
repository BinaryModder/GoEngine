package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/functions"
	"goengine/io/saver"
)

func MenuBar() giu.Widget {

	var selectedObjectIndex int32 = -1   // 0 - cube , 1 - pyramid , 2 - camera
	var selectedMaterialIndex int32 = -1 // 0 - new material (png)
	return giu.Row(
		//Left Part
		giu.Button("Save").OnClick(
			func() {
				if err := functions.SaveMenuBar(editor.State.CurrentScene, editor.State.ProjectPath); err != nil {
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

		giu.Combo("", "SceneObj", []string{"Cube", "Pyramid", "Camera"}, &selectedObjectIndex).Size(120).
			OnChange(func() {
				if editor.State.CurrentScene == nil {
					return
				}

				if err := functions.SceneObjectMenuBar(editor.State.CurrentScene, &selectedObjectIndex); err != nil {
					return
				}

			}),

		giu.Combo("", "Material", []string{"New material"}, &selectedMaterialIndex).Size(120).
			OnChange(func() {
				if err := saver.WriteMaterialFile(editor.State.ProjectPath, "", ""); err != nil {
					return
				}
			}),

		giu.Button("Window"),

		giu.Button("Help"),

		giu.Dummy(toMiddleDummyWeight, toMiddleDummyHeight),

		//Middle Part
		giu.Button("Run").OnClick(
			func() {
				if err := functions.ProjectRunner(editor.State.ProjectPath); err != nil {
					return
				}
			},
		).Size(runSizeWeight, runSizeHeight),
	)
}
