package functions

import (
	"errors"
	"fmt"
	"github.com/sqweek/dialog"
	"goengine/scene"
	"os/exec"
	"path/filepath"
	"runtime"
)

func FileMenuBar() error {

	path, err := dialog.Directory().Title("Choose GoEngine Project").
		Browse()

	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func AssetMenuBar(path string) error {

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", path)

	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)

	default:
		return errors.New("Failed to open assets")
	}
	return cmd.Start()

}
func SaveMenuBar(scene *scene.Scene, path string) error {
	if scene == nil {
		return errors.New("Could not find scene")
	}
	savePath := filepath.Join(
		path,
		"Assets",
		"Scenes",
		"Main.scene",
	)
	if err := scene.SaveToFile(savePath); err != nil {
		return errors.New("Failed to save scene")
	}

	return nil
}

func SceneObjectMenuBar(scen *scene.Scene, obj_index *int32) error {
	if scen == nil {
		return errors.New("Could not find scene")
	}
	switch *obj_index {
	case 0:
		nameObj := fmt.Sprintf("Cube_%d", len(scen.Objects))
		obj := scene.NewCube(nameObj, [3]float32{1, 1, 1})

		if err := scen.AddSceneObjectToTheScene(&obj); err != nil {
			return err
		}
	case 1:
		nameObj := fmt.Sprintf("Pyramid_%d", len(scen.Objects))
		obj := scene.NewPyramid(nameObj, [3]float32{1, 1, 1})

		if err := scen.AddSceneObjectToTheScene(&obj); err != nil {
			return err
		}
	case 2:
		nameObj := fmt.Sprintf("Camera_%d", len(scen.Objects))
		obj := scene.NewCamera(nameObj)

		if err := scen.AddSceneObjectToTheScene(&obj); err != nil {
			return err
		}

	}

	*obj_index = -1

	return nil
}
