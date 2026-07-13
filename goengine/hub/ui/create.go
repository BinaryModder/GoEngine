package ui

import (
	"github.com/AllenDang/giu"

	"goengine/hub"
	"goengine/hub/functions"
)

func CreateWindow() giu.Widget {

	return giu.Child().
		Layout(

			giu.Label(
				"Create 3D Project",
			),

			giu.InputText(
				&hub.State.NewCreateName,
			),

			giu.InputText(
				&hub.State.NewCreatePath,
			),

			giu.Button(
				"Create",
			).
				OnClick(func() {

					functions.CreateProject(
						hub.State.NewCreateName,
						hub.State.NewCreatePath,
					)

				}),
		)

}
