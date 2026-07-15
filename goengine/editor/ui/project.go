package ui

import "github.com/AllenDang/giu"

func Project() giu.Widget {

	return giu.Child().
		Size(-1, ProjectHeight).
		Layout(

			giu.Label("Project"),
		)
}
