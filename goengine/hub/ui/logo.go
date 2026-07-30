package ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub/ui/assets"
	"goengine/ui/scale"
)

// Goofer logo
func Logo() giu.Widget {
	return giu.Image(
		assets.HubTextures.Icon,
	).Size(
		scale.X(logoWidth), scale.Y(logoHeight),
	)
}
