package renderer

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"goengine/editor"
)

type Renderer struct {
	FrameBuffer     *FrameBuffer
	ViewportTexture uint32
}

var State Renderer

var isGridInitialized bool
var isPrimitivesInitialized bool

func Render() {
	if State.FrameBuffer == nil {
		return
	}

	if !isGridInitialized {
		InitGrid()
		isGridInitialized = true
	}

	if !isPrimitivesInitialized {
		InitPrimitives()
		isPrimitivesInitialized = true
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

	if editor.State.CurrentScene != nil {

		gl.UseProgram(MeshProgram)

		projLoc := gl.GetUniformLocation(MeshProgram, gl.Str("projection\x00"))
		viewLoc := gl.GetUniformLocation(MeshProgram, gl.Str("view\x00"))
		gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])
		gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])

		modelLoc := gl.GetUniformLocation(MeshProgram, gl.Str("model\x00"))

		for _, obj := range editor.State.CurrentScene.Objects {
			if obj.Type == "Mesh" {

				switch obj.MeshType {
				case "Cube":

					t := obj.Transform
					model := mgl32.Ident4()
					model = model.Mul4(mgl32.Translate3D(t.Position[0], t.Position[1], t.Position[2]))
					model = model.Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(t.Rotation[0])))
					model = model.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(t.Rotation[1])))
					model = model.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(t.Rotation[2])))
					model = model.Mul4(mgl32.Scale3D(t.Scale[0], t.Scale[1], t.Scale[2]))

					gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

					colorLoc := gl.GetUniformLocation(MeshProgram, gl.Str("objectColor\x00"))

					r, g, b := float32(1.0), float32(1.0), float32(1.0)

					if colorParam, ok := obj.Parameters["Color"].([]interface{}); ok && len(colorParam) == 3 {
						if rv, ok := colorParam[0].(float64); ok {
							r = float32(rv)
						}
						if gv, ok := colorParam[1].(float64); ok {
							g = float32(gv)
						}
						if bv, ok := colorParam[2].(float64); ok {
							b = float32(bv)
						}
					}

					gl.Uniform3f(colorLoc, r, g, b)

					gl.BindVertexArray(Cube.VAO)
					gl.DrawArrays(gl.TRIANGLES, 0, 36)
				}

			}
		}
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
