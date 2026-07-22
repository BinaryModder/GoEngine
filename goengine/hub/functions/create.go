package functions

import (
	"encoding/json"
	"goengine/hub"
	"goengine/project"
	"goengine/scene"
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

	config := project.ProjectConfig{

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

	sceneData := CreateDefaultScene()

	data, err := json.MarshalIndent(
		sceneData,
		"",
		"    ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(root, "Assets", "Scenes", "Main.scene"),
		data,
		0644,
	)
}

func saveNewProjectToList() {

	hub.State.Projects = append(
		hub.State.Projects,
		project.Project{
			Name:       hub.State.NewCreateName,
			Path:       filepath.Join(AbsolutePath(hub.State.NewCreatePath), hub.State.NewCreateName),
			CreatedAt:  time.Now(),
			LastOpened: time.Now(),
		},
	)

}
func CreateDefaultScene() *scene.Scene {
	newScene := &scene.Scene{
		Name:    "Main Scene",
		Objects: []scene.SceneObject{},
	}

	defaultCube := scene.NewCube("defaultCube", [3]float32{1, 1, 1})

	newScene.Objects = append(newScene.Objects, defaultCube)

	return newScene
}
