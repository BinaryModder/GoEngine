package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Renderer struct {
	FrameBuffer     *FrameBuffer
	ViewportTexture uint32
}

var State Renderer

var isGridInitialized bool

func Render() {
	if State.FrameBuffer == nil {
		return
	}

	if !isGridInitialized {
		InitGrid()
		isGridInitialized = true
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, State.FrameBuffer.ID)
	gl.Viewport(0, 0, State.FrameBuffer.Width, State.FrameBuffer.Height)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(gridProgram)

	EditorCam.Update()

	aspectRatio := float32(State.FrameBuffer.Width) / float32(State.FrameBuffer.Height)
	projection := mgl32.Perspective(mgl32.DegToRad(45.0), aspectRatio, 0.1, 100.0)

	view := EditorCam.GetViewMatrix()

	projLoc := gl.GetUniformLocation(gridProgram, gl.Str("projection\x00"))
	viewLoc := gl.GetUniformLocation(gridProgram, gl.Str("view\x00"))

	gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])
	gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])

	gl.BindVertexArray(gridVAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
