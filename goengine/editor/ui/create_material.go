package ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/engine/logger"
	"goengine/io/dialog"
	"goengine/io/saver"
)

func CreateMaterialWindow() giu.Widget {
	return giu.Custom(func() {
		if editor.State.ShowCreateMaterial {
			giu.OpenPopup("Create Material")
			editor.State.ShowCreateMaterial = false
		}

		giu.PopupModal("Create Material").
			Flags(giu.WindowFlagsAlwaysAutoResize).
			Layout(

				giu.Label("Material name"),

				giu.InputText(&editor.State.NewMaterialName),

				giu.Label("Albedo(Path)"),

				giu.Row(

					giu.InputText(&editor.State.NewMaterialSourcePath),
					giu.Button("Browse").
						OnClick(func() {
							file, err := dialog.ChooseImageFile("Choose image")
							if err != nil {

								editor.State.ShowCreateMaterial = false

								//reset data
								editor.State.NewMaterialName = ""
								editor.State.NewMaterialSourcePath = ""

								logger.Error(err.Error())

								return
							}

							editor.State.NewMaterialSourcePath = file

							logger.Info("Image file found")

						}),
				),

				giu.Row(
					giu.Button("Create").OnClick(func() {

						material, err := saver.WriteMaterialFile(
							editor.State.NewMaterialName,
							editor.State.NewMaterialSourcePath,
							editor.State.ProjectPath,
						)

						if err != nil {
							logger.Error(err.Error())
							return

						}
						giu.CloseCurrentPopup()
						editor.State.Materials = append(editor.State.Materials, *material)

						//reset data
						editor.State.NewMaterialName = ""
						editor.State.NewMaterialSourcePath = ""

						logger.Info("New .material file written")

					}),

					giu.Button("Cancel").OnClick(func() {
						giu.CloseCurrentPopup()
						//reset data
						editor.State.NewMaterialName = ""
						editor.State.NewMaterialSourcePath = ""

						logger.Error("Cancelled")

					}),
				),
			).Build()
	})
}
