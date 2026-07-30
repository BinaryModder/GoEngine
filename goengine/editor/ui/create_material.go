package ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
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

								//reset data
								editor.State.NewMaterialName = ""
								editor.State.NewMaterialSourcePath = ""

								return
							}

							editor.State.NewMaterialSourcePath = file

						}),
				),

				giu.Row(
					giu.Button("Create").OnClick(func() {

						material, err := saver.WriteMaterialFile(
							editor.State.NewMaterialName,
							editor.State.NewMaterialSourcePath,
							editor.State.ProjectPath,
						)

						if err == nil {
							giu.CloseCurrentPopup()
							editor.State.Materials = append(editor.State.Materials, *material)
						}

					}),

					giu.Button("Cancel").OnClick(func() {
						giu.CloseCurrentPopup()
					}),
				),
			).Build()
	})
}
