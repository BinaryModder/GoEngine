package editor_ui

import (
	"github.com/AllenDang/giu"
)

func Loop() {
	giu.SingleWindow().Layout(
		MenuBar(),
		//Toolbar(),
		giu.Separator(),
		giu.Row(
			ErrorMessage(),

			Hierarchy(),
			Viewport(),
			Inspector(),
		),
		giu.Separator(),
		Project(),
	)
}
