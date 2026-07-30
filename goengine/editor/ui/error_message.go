package ui

import (
	"github.com/AllenDang/giu"

	"fmt"
	"goengine/editor"
)

func ErrorMessage() giu.Widget {
	return giu.Custom(func() {
		if editor.State.ErrorState != "" {
			giu.OpenPopup("Error message")
		}

		giu.PopupModal("Error message").
			Flags(giu.WindowFlagsAlwaysAutoResize).
			Layout(
				giu.Label(fmt.Sprintf("Error: %v", editor.State.ErrorState)),
			).Build()
	})
}
