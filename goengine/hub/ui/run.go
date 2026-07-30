package ui

import (
	"github.com/AllenDang/giu"
	"goengine/ui/scale"
)

func Run() error {
	window := giu.NewMasterWindow(
		"GoEngine Hub",
		scale.I(1150),
		scale.I(725),
		giu.MasterWindowFlagsNotResizable,
	)

	window.Run(
		Loop,
	)
	return nil

}
