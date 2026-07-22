package scene

import (
	"errors"
)

func (s *Scene) AddSceneObjectToTheScene(sceneObject *SceneObject) error {

	var flag bool

	for _, obj := range s.Objects {
		if obj.Name == sceneObject.Name {
			flag = true
		}
	}

	if flag {
		return errors.New("Object with this name already exists")
	}

	s.Objects = append(
		s.Objects,
		*sceneObject,
	)

	return nil
}
