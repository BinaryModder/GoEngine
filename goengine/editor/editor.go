package editor

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/scene"
	"path/filepath"
)

func Loop() {

	giu.SingleWindow().
		Layout(

			giu.Label(
				"GoEngine Editor",
			),

			giu.Label(
				State.ProjectPath,
			),
		)

}

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
