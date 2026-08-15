package ui

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/imgui"
	"goengine/editor"
	"goengine/engine/logger"
)

func Console() {
	if !editor.State.ShowConsole {
		return
	}

	open := true

	imgui.SetNextWindowSizeV(
		imgui.Vec2{X: 800, Y: 500},
		imgui.CondFirstUseEver,
	)

	imgui.BeginV("Console", &open, 0)

	for _, e := range logger.GetEntries() {

		switch e.Level {
		case "INFO":
			imgui.TextColored(
				imgui.Vec4{X: 0.2, Y: 0.9, Z: 0.2, W: 1.0},
				fmt.Sprintf("[%s]", e.Level),
			)
			imgui.SameLine()
			imgui.Text(e.Text)

		case "WARN":
			imgui.TextColored(
				imgui.Vec4{X: 1.0, Y: 0.8, Z: 0.0, W: 1.0},
				fmt.Sprintf("[%s]", e.Level),
			)
			imgui.SameLine()
			imgui.Text(e.Text)

		case "ERROR":
			imgui.TextColored(
				imgui.Vec4{X: 1.0, Y: 0.2, Z: 0.2, W: 1.0},
				fmt.Sprintf("[%s]", e.Level),
			)
			imgui.SameLine()
			imgui.Text(e.Text)

		default:
			imgui.Text(fmt.Sprintf("[%s] %s", e.Level, e.Text))
		}
	}

	imgui.End()

	if !open {
		editor.State.ShowConsole = false
	}
}
