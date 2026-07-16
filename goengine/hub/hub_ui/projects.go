package hub_ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/hub/functions"
	"goengine/project"
)

func ProjectsView() giu.Widget {

	widgets := []giu.Widget{

		giu.Row(

			giu.Dummy(20, 0),
			giu.Button("New Project").
				OnClick(func() {
					hub.State.ShowCreateProject = true
				}),
			giu.Button("Load Project").
				OnClick(func() {

					loaded_project, err := functions.LoadProject()

					if err != nil {
						fmt.Println(err)
						return
					}

					hub.State.Projects = append(hub.State.Projects, loaded_project)
				}),
		),

		giu.Separator(),
	}

	for _, project := range hub.State.Projects {

		p := project

		widgets = append(
			widgets,

			giu.Separator(),

			projectCard(p),
		)
	}

	return giu.Column(
		widgets...,
	)
}

func projectCard(project project.Project) giu.Widget {

	return giu.Child().
		Size(
			775,
			120,
		).Layout(
		giu.Row(
			giu.Column(
				giu.Label(project.Name),
				giu.Label(functions.ConfigureLabelPath(project.Path)),
				giu.Label("Last opened: "+project.LastOpened.Format("02.01.2006")),
			),

			giu.Dummy(
				35,
				0,
			),

			giu.Button("Open").OnClick(func() {
				functions.OpenEditor(project.Path)
			}),
		),
	)

}
