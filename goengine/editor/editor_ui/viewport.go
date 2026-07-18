package editor_ui

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"goengine/engine/renderer"
)

func Viewport() giu.Widget {

	return giu.Child().
		Size(ViewportWidth, ViewportHeight).
		Layout(

			giu.Custom(func() {

				if renderer.State.FrameBuffer == nil {
					giu.Label("Initializing renderer").Build()
					return
				}
				textureID := imgui.TextureID(uintptr(renderer.State.FrameBuffer.ColorTexture))

				width := float32(ViewportWidth)
				height := float32(ViewportHeight)

				size := imgui.Vec2{
					X: width,
					Y: height,
				}

				uv0 := imgui.Vec2{
					X: 0,
					Y: 1,
				}
				uv1 := imgui.Vec2{
					X: 1,
					Y: 0,
				}

				imgui.ImageV(textureID, size, uv0, uv1)
			}),
		)
}
