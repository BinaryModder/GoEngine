package ui

import (
	"github.com/AllenDang/giu"

	"fmt"
	"goengine/engine/logger"
	"goengine/hub"
	"goengine/io/saver"
	"goengine/settings"
	"goengine/ui/scale"
)

var (
	CurrentLogin string // var for user login changing
)

func MainPanel() giu.Widget {
	widgets := []giu.Widget{}

	switch hub.State.CurrentPage {

	case hub.PageProjects: // Page of Projects

		widgets = append(
			widgets, giu.Separator(), ProjectsView(),
		)

	case hub.PageSettings: // Page of Settings

		CurrentLogin = settings.State.Login

		widgets = append(
			widgets, giu.Row(
				giu.Row(
					giu.Label("Login: "),
					giu.InputText(&CurrentLogin).Size(parameterInputNameObjectSize).Flags(giu.InputTextFlagsEnterReturnsTrue).
						OnChange(func() {
							settings.State.Login = CurrentLogin
							hub.State.SaveSettingsShowButton = true
						}),
				),
			),
			giu.Row(
				giu.Label("Theme: "),
				giu.Label(settings.State.Theme), // It does not work but could soon
			),
			giu.Checkbox(
				"GoEngine with console",
				&settings.State.Console,
			).OnChange(func() {
				hub.State.SaveSettingsShowButton = true
			}),
		)

		if hub.State.SaveSettingsShowButton {
			widgets = append(widgets,
				giu.Button("Save Settings").
					OnClick(func() {
						if err := saver.SaveSettings(); err != nil {
							logger.Error(fmt.Sprintf("Failed to save settins: %s", err.Error()))
							return
						}
						hub.State.SaveSettingsShowButton = false
						logger.Info("Settings saved")
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
