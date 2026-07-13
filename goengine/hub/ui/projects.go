package ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/hub/functions"
)

func ProjectsView() giu.Widget {

	widgets := []giu.Widget{

		giu.Row(

			giu.Label("Projects"),

			giu.Dummy(20, 0),

			giu.Button("New Project").
				OnClick(func() {
					hub.State.ShowCreateProject = true
				}),
			giu.Button("Load Project").
				OnClick(func() {
					functions.ChooseFolder()
				}),
		),

		giu.Separator(),
	}

	for _, project := range hub.State.Projects {

		p := project

		widgets = append(widgets,

			giu.Row(

				giu.Label(p.Name),

				giu.Dummy(30, 0),

				giu.Label(p.Path),

				giu.Dummy(30, 0),

				giu.Label(p.LastOpened),

				giu.Dummy(20, 0),

				giu.Button("Open").
					OnClick(func() {

						functions.OpenEditor(p.Path)

					}),
			),
		)
	}

	return giu.Column(
		widgets...,
	)
}
