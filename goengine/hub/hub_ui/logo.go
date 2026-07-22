package hub_ui

import (
	"github.com/AllenDang/giu"
	"goengine/ui/scale"
)

func Logo() giu.Widget {
	return giu.Image(
		Icon,
	).Size(
		scale.X(170), scale.Y(170),
	)
}
