package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
)

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

	} else {

		for _, object := range editor.State.CurrentScene.Objects {

			obj := object

			widgets = append(

				widgets,

				giu.Selectable(obj.Name),
			)
		}

	}

	return giu.Child().
		Size(HierarchyWidth, ViewportHeight).
		Layout(widgets...)
}
