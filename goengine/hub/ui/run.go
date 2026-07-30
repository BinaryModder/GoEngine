package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/engine/logger"
	"goengine/hub"
	"goengine/ui/scale"
)

func Run() error {

	if err := hub.State.Init(); err != nil {
		logger.Error(fmt.Sprintf("Failed to initialize hub: %s", err))
	}
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
