package hub_ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/settings"
	"goengine/ui/scale"
)

var (
	CurrentLogin           string
	SaveSettingsShowButton bool
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

		CurrentLogin = settings.State.Login

		widgets = append(
			widgets, giu.Row(
				giu.Row(
					giu.Label("Login: "),
					giu.InputText(&CurrentLogin).Size(parameterInputNameObjectSize).Flags(giu.InputTextFlagsEnterReturnsTrue).
						OnChange(func() {
							settings.State.Login = CurrentLogin
							SaveSettingsShowButton = true
						}),
				),
			),
			giu.Row(
				giu.Label("Theme: "),
				giu.Label(settings.State.Theme),
			),
		)

		if SaveSettingsShowButton {
			widgets = append(widgets,
				giu.Button("Save Setting").
					OnClick(func() {
						_ = settings.SaveSettings()
						SaveSettingsShowButton = false
					}),
			)
		}

	}

	if hub.State.ShowCreateProject {
		widgets = append(widgets, CreateProjectView())

	}

	return giu.Child().
		Size(
			scale.X(mainpanelWidth),
			scale.Y(mainpanelHeight),
		).
		Layout(
			widgets...,
		)

}
