package scene

import (
	"errors"
)

type SceneObject struct {
	Name      string    `json:"name"`      // name of scene object
	Type      string    `json:"type"`      // type(mesh , camera , light) of scene object
	Transform Transform `json:"transform"` // transform(position , rotation , scale)
	MeshType  string    `json:"meshtype"`  // type(cube , pyramid , sphere) of mesh scene object

	Script   string    `json:"script"`   // connected to scene object script(in progress...)
	Material *Material `json:"material"` // material for scene object(in progress...)

	Parameters map[string]any `json:"parameters,omitempty"` // other parameters of scene object(color and etc)
}
type Scene struct {
	Name    string        `json:"name"`
	Objects []SceneObject `json:"objects"`
}
type Transform struct {
	Position [3]float32

	Rotation [3]float32

	Scale [3]float32
}

func (s *Scene) AddSceneObjectToTheScene(sceneObject *SceneObject) error {

	var flag bool

	flag = s.HasObject(sceneObject.Name)

	if flag {
		return errors.New("Object with this name already exists")
	}

	s.Objects = append(
		s.Objects,
		*sceneObject,
	)

	return nil
}
func (s *Scene) DeleteSceneObject(sceneObjectName string) error {

	if s == nil {
		return errors.New("Scene is nil")
	}

	for index, object := range s.Objects {
		if object.Name == sceneObjectName {
			s.Objects = append(s.Objects[:index], s.Objects[index+1:]...,
			)
			return nil

		}
	}

	return nil

}

func (s *Scene) HasObject(name string) bool {
	for _, obj := range s.Objects {
		if obj.Name == name {
			return true
		}
	}

	return false

}
