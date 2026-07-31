package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/engine/logger"
	"goengine/ui/layout"
)

// Left side of editor (List of scene objects)
func Hierarchy() giu.Widget {

	widgets := []giu.Widget{

		giu.Label("Hierarchy"),

		giu.Separator(),
	}

	if editor.State.CurrentScene == nil {

		widgets = append(

			widgets,

			giu.Label("No Scene Loaded"),
		)

		logger.Error("Scene not found")

	} else {

		for _, object := range editor.State.CurrentScene.Objects {

			obj := object

			isSelected := editor.State.SelectedObject == obj.Name

			widgets = append(

				widgets,

				giu.Selectable(obj.Name).
					Selected(isSelected).
					OnClick(func() {
						editor.State.SelectedObject = obj.Name
						logger.Info(fmt.Sprintf("Selected object: %s", obj.Name))
					}),
			)
		}

	}

	return giu.Child().
		Size(layout.HierarchyWidth, -layout.ProjectHeight).
		Layout(widgets...)
}
