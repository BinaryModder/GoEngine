package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Renderer struct {
	FrameBuffer     *FrameBuffer
	ViewportTexture uint32
}

var State Renderer

func Render() {

	if State.FrameBuffer == nil {
		return
	}

	fbo := State.FrameBuffer

	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.ID)

	gl.Viewport(
		0,
		0,
		fbo.Width,
		fbo.Height,
	)

	gl.ClearColor(
		0.18,
		0.18,
		0.20,
		1.0,
	)

	gl.Clear(
		gl.COLOR_BUFFER_BIT |
			gl.DEPTH_BUFFER_BIT,
	)

	gl.BindFramebuffer(
		gl.FRAMEBUFFER,
		0,
	)
}
