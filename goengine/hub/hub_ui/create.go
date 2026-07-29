package hub_ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/hub/functions"
	"goengine/hub/validate"
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
					err := validate.ValidateNamePath()
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
						fmt.Println(err)
					}

				}),

			giu.Button("Cancel").
				OnClick(func() {

					hub.State.ShowCreateProject = false

				}),
		)
}
