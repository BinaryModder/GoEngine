package editor

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/scene"
	"path/filepath"
)

func Loop(project string) {

	giu.SingleWindow().
		Layout(

			giu.Label(
				"GoEngine Editor",
			),

			giu.Label(
				project,
			),
		)

}

var current_scene *scene.Scene

func LoadProject(path string) {

	scenePath := filepath.Join(
		path,
		"Assets",
		"Scenes",
		"Main.scene",
	)

	s, err := scene.LoadScene(
		scenePath,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	current_scene = s

	fmt.Printf("Scene loaded\n Name: %s \n", current_scene.Name)

	for _, obj := range current_scene.Objects {
		fmt.Println(
			"Objects:",
			obj.Name,
			obj.Type,
		)
	}
}
