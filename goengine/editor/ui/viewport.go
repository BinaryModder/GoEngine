package ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"github.com/go-gl/gl/v4.1-core/gl"
	"goengine/engine/renderer"
	"goengine/ui/layout"
)

func Viewport() giu.Widget {

	return giu.Child().
		Border(false).
		Size(-layout.InspectorWidth, -layout.ProjectHeight).Flags(
		giu.WindowFlagsNoScrollbar |
			giu.WindowFlagsNoScrollWithMouse).
		Layout(
			giu.Custom(func() {
				if renderer.State.FrameBuffer == nil {
					giu.Label("Initializing renderer").Build()
					return
				}

				availW, availH := giu.GetAvailableRegion()

				if availW < 1 || availH < 1 {
					return
				}

				newW := int32(availW)
				newH := int32(availH)

				if renderer.State.FrameBuffer.Width != newW || renderer.State.FrameBuffer.Height != newH {

					renderer.State.FrameBuffer.Resize(newW, newH)
					renderer.SetWinSize(layout.ViewportWidth, layout.ViewportHeight)
					gl.Viewport(0, 0, newW, newH)
				}

				imgui.SetCursorPos(imgui.Vec2{X: 0, Y: 0})
				textureID := imgui.TextureID(uintptr(renderer.State.FrameBuffer.ColorTexture))

				size := imgui.Vec2{X: availW, Y: availH}
				uv0 := imgui.Vec2{X: 0, Y: 1}
				uv1 := imgui.Vec2{X: 1, Y: 0}

				imgui.ImageV(textureID, size, uv0, uv1)
			}),
		)
}
