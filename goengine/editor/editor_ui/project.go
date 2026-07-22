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

		giu.Label(editor.State.ProjectConfig.Name),
		giu.Separator(),
	}

	rootAssets := filepath.Join(
		editor.State.ProjectPath,
		"Assets",
	)

	if editor.State.CurrentAssetsFolder != rootAssets {

		layout = append(layout,

			giu.Button("← Back").
				OnClick(func() {

					parent := filepath.Dir(
						editor.State.CurrentAssetsFolder,
					)

					files, folder, err :=
						functions.LoadFolder(parent)

					if err != nil {
						return
					}

					editor.State.ProjectFiles = files
					editor.State.CurrentAssetsFolder = folder
				}),

			giu.Separator(),
		)
	}

	row := giu.Layout{}

	for i, file := range editor.State.ProjectFiles {

		f := file

		icon := EditorTextures.FileTexture

		if f.IsDir {
			if f.AmountFiles != 0 {
				icon = EditorTextures.FolderContainingTexture
			} else {
				icon = EditorTextures.FolderEmptyTexture
			}

		}

		card := giu.Child().
			Size(folbutWeight, folbutHeight).
			Layout(

				giu.ImageButton(icon).
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
						editor.State.CurrentAssetsFolder = folder
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
