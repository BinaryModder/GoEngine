package ui

import (
	"github.com/AllenDang/giu"
	"goengine/app"
	"goengine/editor"
	"goengine/editor/functions"
	"goengine/engine/logger"
	"goengine/ui/layout"
)

func MenuBar() giu.Widget {

	var selectedObjectIndex int32 = -1   // 0 - cube , 1 - pyramid , 2 - camera
	var selectedMaterialIndex int32 = -1 // 0 - new material (png)
	return giu.Row(
		//Left Part
		giu.Button("Save").OnClick(
			func() {
				if err := functions.SaveMenuBar(editor.State.CurrentScene, editor.State.ProjectPath); err != nil {
					logger.Error(err.Error())
					return
				}
			},
		).Size(layout.SaveeditSizeWeight, layout.SaveeditSizeHeight),

		giu.Button("Edit").
			Size(layout.SaveeditSizeWeight, layout.SaveeditSizeHeight),

		giu.Button("Assets").OnClick(
			func() {
				if err := functions.AssetMenuBar(editor.State.DefaultAssetsFolder); err != nil {
					logger.Error(err.Error())
					return
				}
			},
		),

		giu.Combo("", "SceneObj", []string{"Cube", "Pyramid", "Camera"}, &selectedObjectIndex).Size(120).
			OnChange(func() {

				if err := functions.SceneObjectMenuBar(editor.State.CurrentScene, &selectedObjectIndex); err != nil {
					logger.Error(err.Error())
					return
				}
				logger.Info("New object created")

			}),

		giu.Combo("", "Material", []string{"New material", "Load material"}, &selectedMaterialIndex).Size(120).
			OnChange(func() {

				switch selectedMaterialIndex {
				case 0:
					editor.State.ShowCreateMaterial = true
				case 1:

					editor.State.ShowLoadMaterial = true
				}

			}),

		giu.Button("Window"),

		giu.Button("Help"),

		giu.Dummy(layout.ToMiddleDummyWeight, layout.ToMiddleDummyHeight),

		//Middle Part
		giu.Button("Run").OnClick(
			func() {
				if err := app.Runtime(editor.State.ProjectPath); err != nil {
					logger.Error(err.Error())
					return
				}
				logger.Info("Runtime starting...")
			},
		).Size(layout.RunSizeWeight, layout.RunSizeHeight),
	)
}
