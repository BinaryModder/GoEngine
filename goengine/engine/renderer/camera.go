package renderer

import (
	"math"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position    mgl32.Vec3
	Front       mgl32.Vec3
	Up          mgl32.Vec3
	Right       mgl32.Vec3
	Yaw         float32
	Pitch       float32
	Speed       float32
	Sensitivity float32

	isDragging     bool
	lockedMousePos imgui.Vec2
}

var EditorCam = Camera{
	Position:    mgl32.Vec3{0.0, 5.0, 10.0},
	Front:       mgl32.Vec3{0.0, 0.0, -1.0},
	Up:          mgl32.Vec3{0.0, 1.0, 0.0},
	Yaw:         -90.0,
	Pitch:       -25.0,
	Speed:       0.1,
	Sensitivity: 0.3,
}

func (c *Camera) Update() {
	if !giu.IsMouseDown(giu.MouseButtonRight) {
		EditorCam.isDragging = false
		return
	}

	io := imgui.CurrentIO()

	if !c.isDragging {
		c.isDragging = true
		c.lockedMousePos = io.MousePos()
	}

	imgui.SetMouseCursor(imgui.MouseCursorNone)

	delta := io.MouseDelta()
	c.Yaw += delta.X * c.Sensitivity
	c.Pitch -= delta.Y * c.Sensitivity

	if c.Pitch > 89.0 {
		c.Pitch = 89.0
	}
	if c.Pitch < -89.0 {
		c.Pitch = -89.0
	}

	radYaw := float64(c.Yaw) * math.Pi / 180.0
	radPitch := float64(c.Pitch) * math.Pi / 180.0

	front := mgl32.Vec3{
		float32(math.Cos(radYaw) * math.Cos(radPitch)),
		float32(math.Sin(radPitch)),
		float32(math.Sin(radYaw) * math.Cos(radPitch)),
	}
	c.Front = front.Normalize()
	c.Right = c.Front.Cross(mgl32.Vec3{0.0, 1.0, 0.0}).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()

	if giu.IsKeyDown(giu.KeyW) {
		c.Position = c.Position.Add(c.Front.Mul(c.Speed))
	}
	if giu.IsKeyDown(giu.KeyS) {
		c.Position = c.Position.Sub(c.Front.Mul(c.Speed))
	}
	if giu.IsKeyDown(giu.KeyA) {
		c.Position = c.Position.Sub(c.Right.Mul(c.Speed))
	}
	if giu.IsKeyDown(giu.KeyD) {
		c.Position = c.Position.Add(c.Right.Mul(c.Speed))
	}
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.Position, c.Position.Add(c.Front), c.Up)
}
