package hub_ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/io/saver"
	"goengine/settings"
	"goengine/ui/scale"
)

var (
	CurrentLogin           string // var for user login changing
	SaveSettingsShowButton bool   // State for "Save" button
)

func MainPanel() giu.Widget {
	widgets := []giu.Widget{}

	switch hub.State.CurrentPage {

	case hub.PageProjects: // Page of Projects

		widgets = append(
			widgets, giu.Separator(), ProjectsView(),
		)

	case hub.PageSettings: // Page of Settings

		if isSettingsFailed { // If loading settings data is failed
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
				giu.Label(settings.State.Theme), // It does not work but could soon
			),
		)

		if SaveSettingsShowButton {
			widgets = append(widgets,
				giu.Button("Save Setting").
					OnClick(func() {
						_ = saver.SaveSettings()
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
