package ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/ui/scale"
)

// Left sidebar
func Sidebar() giu.Widget {
	return giu.Child().
		Size(
			scale.X(sidebarWeight),
			scale.Y(sidebarHeight)).
		Layout(

			giu.Label(
				"GoEngine",
			),

			giu.Separator(),

			giu.Button(
				"Projects",
			).
				OnClick(func() {

					hub.State.CurrentPage = hub.PageProjects //Changing current page

				}),

			giu.Button(
				"⚙ Settings",
			).
				OnClick(func() {

					hub.State.CurrentPage = hub.PageSettings //Changing current page
				}),
			giu.Dummy(0, scale.Y(350)),
			Logo(),
		)
}
