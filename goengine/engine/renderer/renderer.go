package renderer

import (
	"math"
	"os"
	"path/filepath"

	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"goengine/core/state"
	"goengine/engine/logger"
	"goengine/scene"
)

type Renderer struct {
	FrameBuffer     *FrameBuffer
	ViewportTexture uint32
	ProjectPath     string
	UseSceneCamera  bool
}

var State Renderer

var isGridInitialized bool
var isPrimitivesInitialized bool
var isTextureLocsCached bool
var texLoc, useTexLoc int32

var Scene *scene.Scene

func SetUseSceneCamera(use bool) {
	State.UseSceneCamera = use
}

func findSceneCamera() *scene.SceneObject {
	if Scene == nil {
		return nil
	}
	for i := range Scene.Objects {
		if Scene.Objects[i].Type == "Camera" {
			return &Scene.Objects[i]

		}
	}
	logger.Warning("Camera not found")

	return nil
}

func sceneCameraViewMatrix(cam *scene.SceneObject) mgl32.Mat4 {
	pos := mgl32.Vec3{cam.Transform.Position[0], cam.Transform.Position[1], cam.Transform.Position[2]}

	yaw := float64(cam.Transform.Rotation[1])
	pitch := float64(cam.Transform.Rotation[0])

	radYaw := yaw * math.Pi / 180.0
	radPitch := pitch * math.Pi / 180.0

	front := mgl32.Vec3{
		float32(math.Cos(radYaw) * math.Cos(radPitch)),
		float32(math.Sin(radPitch)),
		float32(math.Sin(radYaw) * math.Cos(radPitch)),
	}.Normalize()

	up := mgl32.Vec3{0, 1, 0}
	return mgl32.LookAtV(pos, pos.Add(front), up)
}

func getFieldOfView(cam *scene.SceneObject) float32 {
	if cam == nil {
		return 60.0
	}
	if fov, ok := cam.Parameters["FOV"].(float64); ok {
		return float32(fov)
	}
	return 60.0
}

func Render(CurState state.State) {

	if CurState == nil {
		return
	}

	if State.ProjectPath == "" {
		State.ProjectPath = CurState.GetProjectPath()
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
	var view mgl32.Mat4
	var fov float32 = 45.0

	if State.UseSceneCamera {
		cam := findSceneCamera()
		if cam != nil {
			view = sceneCameraViewMatrix(cam)
			fov = getFieldOfView(cam)
		} else {
			view = EditorCam.GetViewMatrix()
		}
	} else {
		EditorCam.Update()
		view = EditorCam.GetViewMatrix()
	}

	aspectRatio := float32(State.FrameBuffer.Width) / float32(State.FrameBuffer.Height)
	projection := mgl32.Perspective(mgl32.DegToRad(fov), aspectRatio, 0.1, 100.0)

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

		for i := range Scene.Objects {
			obj := &Scene.Objects[i]
			if obj.Type == "Mesh" {

				switch obj.MeshType {
				case "Cube":

					model := getModelMatrix(obj.Transform)
					gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

					r, g, b := getObjectColor(obj)
					gl.Uniform3f(colorLoc, r, g, b)

					useTexture := bindMaterialTexture(obj, texLoc, useTexLoc)

					gl.BindVertexArray(Cube.VAO)
					gl.DrawArrays(gl.TRIANGLES, 0, Cube.VertexCount)

					if useTexture {
						gl.BindTexture(gl.TEXTURE_2D, 0)
						gl.Uniform1i(useTexLoc, 0)
					}

				case "Pyramid":
					model := getModelMatrix(obj.Transform)
					gl.UniformMatrix4fv(modelLoc, 1, false, &model[0])

					r, g, b := getObjectColor(obj)
					gl.Uniform3f(colorLoc, r, g, b)

					useTexture := bindMaterialTexture(obj, texLoc, useTexLoc)

					gl.BindVertexArray(Pyramid.VAO)
					gl.DrawArrays(gl.TRIANGLES, 0, Pyramid.VertexCount)

					if useTexture {
						gl.BindTexture(gl.TEXTURE_2D, 0)
						gl.Uniform1i(useTexLoc, 0)
					}

				}

			}
		}

		if !State.UseSceneCamera {
			for i := range Scene.Objects {
				obj := &Scene.Objects[i]
				if obj.Type == "Camera" {
					camModel := getCameraModelMatrix(obj)
					gl.UniformMatrix4fv(modelLoc, 1, false, &camModel[0])
					gl.Uniform3f(colorLoc, 0.2, 0.8, 0.2)
					gl.Uniform1i(useTexLoc, 0)
					gl.BindVertexArray(Pyramid.VAO)
					gl.DrawArrays(gl.TRIANGLES, 0, Pyramid.VertexCount)
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

func getCameraModelMatrix(cam *scene.SceneObject) mgl32.Mat4 {
	yawRad := float64(cam.Transform.Rotation[1]) * math.Pi / 180.0
	pitchRad := float64(cam.Transform.Rotation[0]) * math.Pi / 180.0

	front := mgl32.Vec3{
		float32(math.Cos(yawRad) * math.Cos(pitchRad)),
		float32(math.Sin(pitchRad)),
		float32(math.Sin(yawRad) * math.Cos(pitchRad)),
	}.Normalize()

	worldUp := mgl32.Vec3{0, 1, 0}
	right := worldUp.Cross(front).Normalize()
	camUp := front.Cross(right)

	rotation := mgl32.Mat4FromCols(
		mgl32.Vec4{right[0], right[1], right[2], 0},
		mgl32.Vec4{front[0], front[1], front[2], 0},
		mgl32.Vec4{camUp[0], camUp[1], camUp[2], 0},
		mgl32.Vec4{0, 0, 0, 1},
	)

	model := mgl32.Ident4()
	model = model.Mul4(mgl32.Translate3D(
		cam.Transform.Position[0],
		cam.Transform.Position[1],
		cam.Transform.Position[2],
	))
	model = model.Mul4(rotation)
	model = model.Mul4(mgl32.Scale3D(0.3, 0.3, 0.3))
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

func resolveTexturePath(albedo string) string {
	if _, err := os.Stat(albedo); err == nil {
		return albedo
	}

	base := filepath.Base(albedo)
	fallback := filepath.Join(State.ProjectPath, "Assets", "Textures", base)
	if _, err := os.Stat(fallback); err == nil {
		return fallback
	}

	return albedo
}

func bindMaterialTexture(obj *scene.SceneObject, texLoc, useTexLoc int32) bool {
	if obj.Material != nil && obj.Material.Albedo != "" && obj.Material.Albedo != "null" {
		texID, err := loadTexture(resolveTexturePath(obj.Material.Albedo))
		if err != nil {
			logger.Info(fmt.Sprintf("Failed to load texture: %v", err))
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
