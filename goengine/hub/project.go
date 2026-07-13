package hub

import (
	"encoding/json"
	"fmt"
	"github.com/sqweek/dialog"
	"os"
	"path/filepath"
)

func CreateProject(
	name string,
	location string,
) {

	root :=
		filepath.Join(
			location,
			name,
		)

	folders := []string{

		"Assets",
		"Assets/Scenes",
		"Assets/Meshes",
		"Assets/Materials",
		"Assets/Textures",
		"Assets/Scripts",

		"ProjectSettings",
	}

	for _, folder := range folders {

		os.MkdirAll(
			filepath.Join(root, folder),
			0755,
		)

	}

	createSettings(root)

	createScene(root)

}

func createSettings(root string) {

	data, _ :=
		json.MarshalIndent(

			map[string]string{

				"name": "GoEngine Project",
				"type": "3D",
			},

			"",
			" ",
		)

	os.WriteFile(

		filepath.Join(
			root,
			"ProjectSettings",
			"project.json",
		),

		data,

		0644,
	)

}

func createScene(root string) {

	scene := `

{
 "objects":[

  {
   "name":"Camera",
   "type":"Camera"
  },

  {
   "name":"Light",
   "type":"DirectionalLight"
  },

  {
   "name":"Cube",
   "type":"Mesh"
  }

 ]

}

`

	os.WriteFile(

		filepath.Join(
			root,
			"Assets",
			"Scenes",
			"Main.scene",
		),

		[]byte(scene),

		0644,
	)

}

func chooseFolder() {

	folder, err := dialog.Directory().Title("Choose Project path").Browse()

	if err != nil {
		fmt.Println(err)

		return
	}

	projectPath = folder

}
