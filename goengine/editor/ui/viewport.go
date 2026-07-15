package ui

import "github.com/AllenDang/giu"

func Viewport() giu.Widget {

	return giu.Child().
		Size(ViewportWidth, ViewportHeight).
		Layout(

			giu.Label("Viewport"),
		)
}
