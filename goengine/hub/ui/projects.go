package ui

import (
	"github.com/AllenDang/giu"
	"goengine/app"
	"goengine/hub"
	"goengine/hub/functions"
	"goengine/project"
	"goengine/ui"
	"goengine/ui/scale"
)

// Project part of interface (project cards , "New" , "Load")
func ProjectsView() giu.Widget {

	widgets := []giu.Widget{

		giu.Row(

			giu.Dummy(scale.X(20), 0),
			giu.Button("New Project").
				OnClick(func() {
					hub.State.ShowCreateProject = true
				}),
			giu.Button("Load Project").
				OnClick(func() {
					functions.LoadProject()
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
			scale.X(projectCardWeight),
			scale.Y(projectCardHeight),
		).Layout(
		giu.Row(
			giu.Column(
				giu.Label(project.Name),
				giu.Label(ui.TruncatePath(project.Path, pathSize)),
				giu.Label("Last opened: "+project.LastOpened.Format("02.01.2006")),
			),

			giu.Dummy(
				35,
				0,
			),

			giu.Button("Open").OnClick(func() {
				app.Editor(project.Path)
			}),
		),
	)

}
