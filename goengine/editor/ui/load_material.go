package ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/io/dialog"
	"goengine/io/loader"
	"goengine/io/saver"
)

func LoadMaterialWindow() giu.Widget {
	return giu.Custom(func() {
		if editor.State.ShowLoadMaterial {
			giu.OpenPopup("Load Material")
			editor.State.ShowLoadMaterial = false
		}

		giu.PopupModal("Load Material").
			Flags(giu.WindowFlagsAlwaysAutoResize).
			Layout(

				giu.Label("Source"),

				giu.Row(

					giu.InputText(&editor.State.LoadMaterialSourcePath),
					giu.Button("Browse").
						OnClick(func() {
							file, err := dialog.ChooseMaterialFile("Choose .material")

							if err != nil {
								editor.State.ShowLoadMaterial = false
								//reset data
								editor.State.LoadMaterialSourcePath = ""
								return
							}

							editor.State.LoadMaterialSourcePath = file // write file

						}),
				),

				giu.Row(
					giu.Button("Load").OnClick(func() {
						//Loading material to the state
						material, err := loader.LoadMaterialFile(editor.State.LoadMaterialSourcePath)

						if err != nil {
							editor.State.ErrorState = err.Error()
							giu.CloseCurrentPopup()
							editor.State.LoadMaterialSourcePath = ""
							return
						}

						giu.CloseCurrentPopup()
						editor.State.Materials = append(editor.State.Materials, *material)
						editor.State.LoadMaterialSourcePath = ""

						//Writing loaded material to the "Materials" folder

						_, err = saver.WriteMaterialFile(
							material.Name,
							material.Albedo,
							editor.State.ProjectPath,
						)

					}),

					giu.Button("Cancel").OnClick(func() {
						giu.CloseCurrentPopup()
						//reset data
						editor.State.LoadMaterialSourcePath = ""

					}),
				),
			).Build()
	})
}
