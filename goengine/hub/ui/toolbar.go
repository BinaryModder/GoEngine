package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/hub"
	"goengine/hub/functions"
)

func Toolbar() giu.Widget {

	return giu.Row(
		giu.Button(
			"Create 3D Project",
		).
			OnClick(func() {

				hub.State.ShowCreateProject = true

			}),
		giu.Row(
			giu.InputText(&hub.State.NewCreatePath),
			giu.Button("Browse").OnClick(func() {
				err := functions.ChooseFolder()
				if err != nil {
					fmt.Println(err)
				}
			}),
		),

		giu.Button(
			"Load 3D Project",
		).OnClick(func() {
			functions.LoadProject()
		}),
	)
}
