package editor

import (
	"fmt"
	"goengine/scene"
	"path/filepath"
)

var currentScene *scene.Scene

func LoadScene() error {

	scenePath := filepath.Join(
		State.ProjectPath,
		"Assets",
		"Scenes",
		"Main.scene",
	)

	s, err := scene.LoadScene(
		scenePath,
	)

	if err != nil {
		return err
	}

	currentScene = s

	fmt.Printf("Scene loaded\n Name: %s \n", currentScene.Name)

	for _, obj := range currentScene.Objects {
		fmt.Println(
			"Objects:",
			obj.Name,
			obj.Type,
		)
	}
	return nil
}
