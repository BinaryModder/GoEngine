package hub_ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/settings"
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
		if isSettingsFailed {
			widgets = append(widgets,
				giu.Label("Failed to create configuration file"),
			)
		}
		widgets = append(
			widgets, giu.Row(
				giu.Label("Login: "),
				giu.Label(settings.State.Login),
			),
			giu.Row(
				giu.Label("Theme: "),
				giu.Label(settings.State.Theme),
			),
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
