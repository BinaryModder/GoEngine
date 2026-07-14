package functions

import (
	"encoding/json"
	"github.com/sqweek/dialog"
	"goengine/hub"
	"os"
	"path/filepath"
	"time"
)

func CreateProject(
	name string,
	location string,
) error {

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

		err := os.MkdirAll(
			filepath.Join(root, folder),
			0755,
		)

		if err != nil {
			return err
		}

	}

	err := createProjFile(root, name)
	if err != nil {
		return err
	}

	err = createScene(root)
	if err != nil {
		return err
	}

	saveNewProjectToList()

	return nil

}
func createProjFile(
	root string,
	name string,
) error {

	config := hub.ProjectConfig{

		Name: name,

		Version: "1.0.0",

		EngineVersion: "0.1.0",

		CreatedAt: time.Now(),
	}

	data, err := json.MarshalIndent(
		config,
		"",
		"    ",
	)

	if err != nil {
		return err
	}

	path := filepath.Join(

		root,

		"ProjectSettings",

		"project.json",
	)

	return os.WriteFile(

		path,

		data,

		0644,
	)

}

func createScene(root string) error {

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

	return os.WriteFile(

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

func saveNewProjectToList() {

	hub.State.Projects = append(
		hub.State.Projects,
		hub.Project{
			Name:       hub.State.NewCreateName,
			Path:       filepath.Join(AbsolutePath(hub.State.NewCreatePath), hub.State.NewCreateName),
			CreatedAt:  time.Now(),
			LastOpened: time.Now(),
		},
	)

}

func ChooseFolder() (string, error) {

	folder, err := dialog.Directory().Title("Choose Project path").Browse()

	if err != nil {
		return "", err
	}

	return folder, nil

}
