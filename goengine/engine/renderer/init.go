package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func Init(width int32, height int32) error {

	if err := gl.Init(); err != nil {
		return err
	}

	fbo, err := NewFrameBuffer(width, height)

	if err != nil {
		return err
	}

	State.FrameBuffer = fbo

	State.ViewportTexture = fbo.ColorTexture

	return nil
}
