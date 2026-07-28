package script

import (
	"goengine/scene"
)

type Script interface {
	Start()
	Update(dt float32)
	Destroy()
}

type BaseScript struct {
	Object *scene.SceneObject
}

func (BaseScript) Start()   {}
func (BaseScript) Update()  {}
func (BaseScript) Destroy() {}
