package hub_ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/ui/scale"
)

func Sidebar() giu.Widget {
	return giu.Child().
		Size(
			scale.X(220),
			scale.Y(700)).
		Layout(

			giu.Label(
				"GoEngine",
			),

			giu.Separator(),

			giu.Button(
				"Projects",
			).
				OnClick(func() {

					hub.State.CurrentPage = hub.PageProjects

				}),

			giu.Button(
				"⚙ Settings",
			).
				OnClick(func() {

					hub.State.CurrentPage = hub.PageSettings
				}),
			giu.Dummy(0, scale.Y(350)),
			Logo(),
		)
}
