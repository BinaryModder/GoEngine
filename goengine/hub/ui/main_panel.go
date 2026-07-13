package ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
)

func MainPanel() giu.Widget {

	switch hub.State.CurrentPage {

	case hub.PageProjects:

		return giu.Child().
			Size(900, 700).
			Layout(
				Toolbar(),
				giu.Separator(),
				ProjectsView(),
			)

	case hub.PageSettings:

		return giu.Child().
			Size(900, 700).
			Layout(

				giu.Label(
					"Settings",
				),
			)

	}

	return giu.Child().
		Size(900, 700).
		Layout(

			giu.Label(
				"Welcome to GoEngine Hub",
			),
		)

}
