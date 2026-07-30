package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/io/dialog"
	"goengine/project"
	"goengine/ui/scale"
)

func CreateProjectView() giu.Widget {

	return giu.Child().
		Size(
			scale.X(400),
			scale.Y(300)).
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

						folder, err := dialog.ChooseFolder()

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
					err := project.ValidateNewProjectNamePath(
						hub.State.NewCreateName,
						hub.State.NewCreatePath,
					)
					if err != nil {
						hub.State.ErrorMessage = err.Error()
						return
					}

					newProject, err := project.CreateProject(
						hub.State.NewCreateName,
						hub.State.NewCreatePath,
					)

					if err == nil {

						hub.State.ShowCreateProject = false
						hub.State.Projects = append(hub.State.Projects, *newProject)

					} else {
						//resets input data
						hub.State.NewCreateName = ""
						hub.State.NewCreatePath = ""

						//hiding creating window
						hub.State.ShowCreateProject = false

						fmt.Println(err)
					}

				}),

			giu.Button("Cancel").
				OnClick(func() {
					//hiding creating window
					hub.State.ShowCreateProject = false

					//resets input data
					hub.State.NewCreateName = ""
					hub.State.NewCreatePath = ""

				}),
		)
}
