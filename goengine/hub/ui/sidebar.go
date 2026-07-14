package ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub"
)

func Sidebar() giu.Widget {
	return giu.Child().
		Size(220, 700).
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
			giu.Dummy(0, 400),
			Logo(),
		)
}
