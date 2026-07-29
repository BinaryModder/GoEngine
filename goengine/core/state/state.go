package state

import (
	"goengine/scene"
)

type State interface {
	Init() error

	GetProjectPath() string
	GetProjectScene() *scene.Scene
}
