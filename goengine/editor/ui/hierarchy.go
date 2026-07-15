package ui

import "github.com/AllenDang/giu"

func Hierarchy() giu.Widget {

	return giu.Child().
		Size(HierarchyWidth, ViewportHeight).
		Layout(

			giu.Label("Hierarchy"),
		)
}
