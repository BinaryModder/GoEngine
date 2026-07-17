package editor_ui

import (
	"path/filepath"

	"github.com/AllenDang/giu"

	"goengine/editor"
	"goengine/editor/functions"
)

const columns = 13

func Project() giu.Widget {

	layout := giu.Layout{

		giu.Label("Project"),
		giu.Separator(),
	}

	rootAssets := filepath.Join(
		editor.State.ProjectPath,
		"Assets",
	)

	if editor.State.AssetsFolder != rootAssets {

		layout = append(layout,

			giu.Button("← Back").
				OnClick(func() {

					parent := filepath.Dir(
						editor.State.AssetsFolder,
					)

					files, folder, err :=
						functions.LoadFolder(parent)

					if err != nil {
						return
					}

					editor.State.ProjectFiles = files
					editor.State.AssetsFolder = folder
				}),

			giu.Separator(),
		)
	}

	row := giu.Layout{}

	for i, file := range editor.State.ProjectFiles {

		f := file

		icon := "📄"

		if f.IsDir {
			icon = "📁"
		}

		card := giu.Child().
			Size(125, 115).
			Layout(

				giu.Button(icon).
					Size(64, 64).
					OnClick(func() {

						if !f.IsDir {
							return
						}

						files, folder, err :=
							functions.LoadFolder(f.Path)

						if err != nil {
							return
						}

						editor.State.ProjectFiles = files
						editor.State.AssetsFolder = folder
					}),

				giu.Label(f.Name),
			)

		row = append(row, card)

		if (i+1)%columns == 0 {

			layout = append(layout,
				giu.Row(row...),
			)

			row = giu.Layout{}
		}
	}

	if len(row) > 0 {

		layout = append(layout,
			giu.Row(row...),
		)
	}

	return giu.Child().
		Size(-1, -1).
		Layout(layout...)
}
