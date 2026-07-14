package ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
)

func MainPanel() giu.Widget {
	widgets := []giu.Widget{}

	switch hub.State.CurrentPage {

	case hub.PageProjects:

		widgets = append(
			widgets, giu.Separator(), ProjectsView(),
		)
		//Toolbar(),

	case hub.PageSettings:
		widgets = append(
			widgets, giu.Label("Settings"),
		)

	}

	if hub.State.ShowCreateProject {
		widgets = append(widgets, CreateProjectView())

	}

	return giu.Child().
		Size(900, 700).
		Layout(
			widgets...,
		)

}
