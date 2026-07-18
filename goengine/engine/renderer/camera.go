package renderer

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Camera struct {
	Position mgl32.Vec3
	Rotation mgl32.Vec3
	Fov      float32
	Near     float32
	Far      float32
}
