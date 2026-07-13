package hub

import (
	"github.com/AllenDang/giu"
)

var (
	showCreate bool

	projectName = "NewProject"
	projectPath = "./Projects"

	projectToLoad string
)

func Loop() {

	layout := []giu.Widget{

		giu.Dummy(0, 100),

		giu.Label(
			"GoEngine Hub",
		),

		giu.Dummy(0, 30),

		giu.Button(
			"Create 3D Project",
		).
			OnClick(func() {

				showCreate = true

			}),
		giu.Row(
			giu.InputText(&projectPath),
			giu.Button("Browse").OnClick(func() {
				chooseFolder()
			}),
		),

		giu.Button(
			"Load 3D Project",
		).OnClick(func() {
			LoadProject()
		}),
	}

	if showCreate {

		layout = append(
			layout,

			giu.Separator(),

			giu.Label(
				"Create 3D Project",
			),

			giu.InputText(
				&projectName,
			),

			giu.InputText(
				&projectPath,
			),

			giu.Button(
				"Create",
			).
				OnClick(func() {

					CreateProject(
						projectName,
						projectPath,
					)

				}),
		)

	}

	giu.SingleWindow().
		Layout(layout...)

}
