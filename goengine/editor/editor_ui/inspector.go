package editor_ui

import "github.com/AllenDang/giu"

func Inspector() giu.Widget {

	return giu.Child().
		Size(InspectorWidth, ViewportHeight).
		Layout(

			giu.Label("Inspector"),
		)
}
