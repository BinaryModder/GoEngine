package renderer

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"goengine/core/state"
	"goengine/scene"
)

type Renderer struct {
	FrameBuffer     *FrameBuffer
	ViewportTexture uint32
}

var State Renderer

var isGridInitialized bool
var isPrimitivesInitialized bool
var isTextureLocsCached bool
var texLoc, useTexLoc int32

var Scene *scene.Scene

func Render(CurState state.State) {

	if CurState == nil {
		return
	}

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
	//binding scene
	Scene = CurState.GetProjectScene()

	gl.BindFramebuffer(gl.FRAMEBUFFER, State.FrameBuffer.ID)
	gl.Viewport(0, 0, State.FrameBuffer.Width, State.FrameBuffer.Height)

	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.UseProgram(gridProgram)

	//camera
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

		if Scene != nil {

			gl.UseProgram(MeshProgram)

			projLoc := gl.GetUniformLocation(MeshProgram, gl.Str("projection\x00"))
			viewLoc := gl.GetUniformLocation(MeshProgram, gl.Str("view\x00"))
			gl.UniformMatrix4fv(projLoc, 1, false, &projection[0])
			gl.UniformMatrix4fv(viewLoc, 1, false, &view[0])

			modelLoc := gl.GetUniformLocation(MeshProgram, gl.Str("model\x00"))
			colorLoc := gl.GetUniformLocation(MeshProgram, gl.Str("objectColor\x00"))

			if !isTextureLocsCached {
				texLoc = gl.GetUniformLocation(MeshProgram, gl.Str("uTexture\x00"))
				useTexLoc = gl.GetUniformLocation(MeshProgram, gl.Str("uUseTexture\x00"))
				isTextureLocsCached = true
			}

			for _, obj := range Scene.Objects {
				if obj.Type == "Mesh" {

					switch obj.MeshType {
					case "Cube":

						model := getModelMatrix(obj.Transform)
						gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

						r, g, b := getObjectColor(&obj)
						gl.Uniform3f(colorLoc, r, g, b)

						useTexture := bindMaterialTexture(&obj, texLoc, useTexLoc)

						gl.BindVertexArray(Cube.VAO)
						gl.DrawArrays(gl.TRIANGLES, 0, Cube.VertexCount)

						if useTexture {
							gl.BindTexture(gl.TEXTURE_2D, 0)
							gl.Uniform1i(useTexLoc, 0)
						}

					case "Pyramid":
						model := getModelMatrix(obj.Transform)
						gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

						r, g, b := getObjectColor(&obj)
						gl.Uniform3f(colorLoc, r, g, b)

						useTexture := bindMaterialTexture(&obj, texLoc, useTexLoc)

						gl.BindVertexArray(Pyramid.VAO)
						gl.DrawArrays(gl.TRIANGLES, 0, Pyramid.VertexCount)

						if useTexture {
							gl.BindTexture(gl.TEXTURE_2D, 0)
							gl.Uniform1i(useTexLoc, 0)
						}

					}

				}
			}
		}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func getModelMatrix(t scene.Transform) mgl32.Mat4 {
	model := mgl32.Ident4()
	model = model.Mul4(mgl32.Translate3D(t.Position[0], t.Position[1], t.Position[2]))
	model = model.Mul4(mgl32.HomogRotate3DX(mgl32.DegToRad(t.Rotation[0])))
	model = model.Mul4(mgl32.HomogRotate3DY(mgl32.DegToRad(t.Rotation[1])))
	model = model.Mul4(mgl32.HomogRotate3DZ(mgl32.DegToRad(t.Rotation[2])))
	model = model.Mul4(mgl32.Scale3D(t.Scale[0], t.Scale[1], t.Scale[2]))
	return model
}

func getObjectColor(obj *scene.SceneObject) (float32, float32, float32) {
	if obj.Material != nil {
		return obj.Material.Color[0], obj.Material.Color[1], obj.Material.Color[2]
	}

	if colorParam, ok := obj.Parameters["Color"].([]interface{}); ok && len(colorParam) == 3 {
		r, g, b := float32(1.0), float32(1.0), float32(1.0)
		if rv, ok := colorParam[0].(float64); ok {
			r = float32(rv)
		}
		if gv, ok := colorParam[1].(float64); ok {
			g = float32(gv)
		}
		if bv, ok := colorParam[2].(float64); ok {
			b = float32(bv)
		}
		return r, g, b
	}

	return 1.0, 1.0, 1.0
}

func bindMaterialTexture(obj *scene.SceneObject, texLoc, useTexLoc int32) bool {
	if obj.Material != nil && obj.Material.Albedo != "" && obj.Material.Albedo != "null" {
		texID, err := loadTexture(obj.Material.Albedo)
		if err != nil {
			log.Printf("Failed to load texture: %v", err)
			gl.Uniform1i(useTexLoc, 0)
			return false
		}
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, texID)
		gl.Uniform1i(texLoc, 0)
		gl.Uniform1i(useTexLoc, 1)
		return true
	}
	gl.Uniform1i(useTexLoc, 0)
	return false
}
