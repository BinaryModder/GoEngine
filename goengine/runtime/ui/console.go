package ui

import (
	"github.com/AllenDang/giu"
	"goengine/engine/console"
)

func Console() giu.Widget {
	return giu.Child().
		Border(true).
		Size(-1, ConsoleHeight).
		Layout(

			giu.Label("Console"),

			giu.Separator(),

			giu.Custom(func() {
				for _, log := range console.State.Logs {
					giu.Label(log.Message).Build()
				}
			}),
		)
}
