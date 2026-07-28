package hub_ui

import (
	"github.com/AllenDang/giu"
	"goengine/ui/scale"
)

// Goofer logo
func Logo() giu.Widget {
	return giu.Image(
		Icon,
	).Size(
		scale.X(logoWidth), scale.Y(logoHeight),
	)
}
