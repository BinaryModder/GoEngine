package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/io/dialog"
	"goengine/io/saver"
	"goengine/scene"
	"goengine/ui/scale"
)

func CreateMaterialWindow() giu.Widget {

	return giu.Child().
		Size(
			scale.X(400),
			scale.Y(300)).
		Border(true).
		Layout(

			giu.Label(
				"Create new material",
			),

			giu.Separator(),

			giu.Label(
				"Material name",
			),

			giu.InputText(
				&editor.State.NewMaterialName,
			),

			giu.Label(
				"Location",
			),

			giu.Row(

				giu.InputText(
					&editor.State.NewMaterialSourcePath,
				),

				giu.Button("Browse").
					OnClick(func() {

						folder, err := dialog.ChooseFolder()

						if err != nil {
							fmt.Println(err)
							return
						}
						editor.State.NewMaterialSourcePath = folder

					}),
			),

			giu.Separator(),

			giu.Button("Create").
				OnClick(func() {
					err := scene.ValidateMaterialNamePath(editor.State.NewMaterialName, editor.State.NewMaterialSourcePath)
					if err != nil {
						return
					}

					newMaterial, err := saver.WriteMaterialFile(
						editor.State.NewMaterialName,
						editor.State.NewMaterialSourcePath,
						editor.State.ProjectPath,
					)

					if err == nil {

						editor.State.ShowCreateMaterial = false
						editor.State.Materials = append(editor.State.Materials, *newMaterial)

						//resets input data
						editor.State.NewMaterialName = ""
						editor.State.NewMaterialSourcePath = ""

					} else {
						//resets input data
						editor.State.NewMaterialName = ""
						editor.State.NewMaterialSourcePath = ""

						//hiding creating window
						editor.State.ShowCreateMaterial = false

						fmt.Println(err)
					}

				}),

			giu.Button("Cancel").
				OnClick(func() {
					//hiding creating window
					editor.State.ShowCreateMaterial = false

					//resets input data
					editor.State.NewMaterialName = ""
					editor.State.NewMaterialSourcePath = ""

				}),
		)
}
