package editor_ui

import (
	"github.com/AllenDang/giu"

	"goengine/editor"
)

func ErrorMessage() giu.Widget {
	show := editor.State.ErrorState != ""

	return giu.Condition(
		show,
		giu.Child().Size(-1, 60).Border(true).Layout(
			giu.Style().To(
				giu.Row(
					giu.Label(editor.State.ErrorState),
				),
			),
		),
		giu.Layout{},
	)
}
