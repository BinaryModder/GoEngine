package hub_ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/ui/scale"
)

func MainPanel() giu.Widget {
	widgets := []giu.Widget{}

	switch hub.State.CurrentPage {

	case hub.PageProjects:

		widgets = append(
			widgets, giu.Separator(), ProjectsView(),
		)

	case hub.PageSettings:
		widgets = append(
			widgets, giu.Label("Settings"),
		)

	}

	if hub.State.ShowCreateProject {
		widgets = append(widgets, CreateProjectView())

	}

	return giu.Child().
		Size(
			scale.X(900),
			scale.Y(700),
		).
		Layout(
			widgets...,
		)

}
