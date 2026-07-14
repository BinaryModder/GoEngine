package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/hub/functions"
)

func CreateProjectView() giu.Widget {

	return giu.Child().
		Size(400, 300).
		Border(true).
		Layout(

			giu.Label(
				"Create 3D Project",
			),

			giu.Separator(),

			giu.Label(
				"Project Name",
			),

			giu.InputText(
				&hub.State.NewCreateName,
			),

			giu.Label(
				"Location",
			),

			giu.Row(

				giu.InputText(
					&hub.State.NewCreatePath,
				),

				giu.Button("Browse").
					OnClick(func() {

						folder, err := functions.ChooseFolder()

						if err != nil {
							fmt.Println(err)
							return
						}
						hub.State.NewCreatePath = folder

					}),
			),

			giu.Separator(),

			giu.Button("Create").
				OnClick(func() {

					err := functions.CreateProject(
						hub.State.NewCreateName,
						hub.State.NewCreatePath,
					)

					if err == nil {

						hub.State.ShowCreateProject = false

					} else {
						fmt.Println(err)
					}

				}),

			giu.Button("Cancel").
				OnClick(func() {

					hub.State.ShowCreateProject = false

				}),
		)
}
