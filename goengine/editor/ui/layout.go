package ui

import (
	"github.com/AllenDang/giu"
)

func Layout() {
	giu.SingleWindow().Layout(

		MenuBar(),
		//Toolbar(),
		giu.Separator(),
		giu.Row(
			Hierarchy(),
			Viewport(),
			Inspector(),
		),
		giu.Separator(),
		Project(),
		StatusBar(),
	)
}
