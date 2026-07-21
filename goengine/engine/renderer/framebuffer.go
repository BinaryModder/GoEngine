package renderer

import (
	"errors"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type FrameBuffer struct {
	ID           uint32
	ColorTexture uint32
	DepthBuffer  uint32

	Width  int32
	Height int32
}

func (fbo *FrameBuffer) NewFrameBuffer() error {

	gl.GenFramebuffers(1, &fbo.ID)

	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.ID)

	gl.GenTextures(1, &fbo.ColorTexture)

	gl.BindTexture(gl.TEXTURE_2D, fbo.ColorTexture)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA8,
		fbo.Width,
		fbo.Height,
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		nil,
	)

	gl.TexParameteri(
		gl.TEXTURE_2D,
		gl.TEXTURE_MIN_FILTER,
		gl.LINEAR,
	)

	gl.TexParameteri(
		gl.TEXTURE_2D,
		gl.TEXTURE_MAG_FILTER,
		gl.LINEAR,
	)

	gl.FramebufferTexture2D(
		gl.FRAMEBUFFER,
		gl.COLOR_ATTACHMENT0,
		gl.TEXTURE_2D,
		fbo.ColorTexture,
		0,
	)

	gl.GenRenderbuffers(1, &fbo.DepthBuffer)

	gl.BindRenderbuffer(
		gl.RENDERBUFFER,
		fbo.DepthBuffer,
	)

	gl.RenderbufferStorage(
		gl.RENDERBUFFER,
		gl.DEPTH24_STENCIL8,
		fbo.Width,
		fbo.Height,
	)

	gl.FramebufferRenderbuffer(
		gl.FRAMEBUFFER,
		gl.DEPTH_STENCIL_ATTACHMENT,
		gl.RENDERBUFFER,
		fbo.DepthBuffer,
	)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {

		return errors.New("framebuffer is not complete")
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	return nil
}

func (fbo *FrameBuffer) Resize(newWidth int32, newHeight int32) {
	if fbo.Width == newWidth && fbo.Height == newHeight {
		return
	}

	gl.DeleteFramebuffers(1, &fbo.ID)
	gl.DeleteTextures(1, &fbo.ColorTexture)
	gl.DeleteRenderbuffers(1, &fbo.DepthBuffer)

	Init(newWidth, newHeight)
}
