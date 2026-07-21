package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func Init(width int32, height int32) error {

	fbo := &FrameBuffer{
		Width:  width,
		Height: height,
	}

	if err := gl.Init(); err != nil {
		return err
	}

	if err := fbo.NewFrameBuffer(); err != nil {
		return err
	}

	State.FrameBuffer = fbo

	State.ViewportTexture = fbo.ColorTexture

	return nil
}
