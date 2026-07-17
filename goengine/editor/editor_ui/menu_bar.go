package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/functions"
)

func MenuBar() giu.Widget {

	return giu.Row(

		giu.Button("File").OnClick(
			func() {
				if err := functions.FileMenuBar(); err != nil {
					return
				}
			},
		),

		giu.Button("Edit"),

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
	)
}
