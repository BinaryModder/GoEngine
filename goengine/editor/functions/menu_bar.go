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
func SaveProject(scene *scene.Scene, projectPath string) error {
	if scene != nil {
		savePath := filepath.Join(
			projectPath,
			"Assets",
			"Scenes",
			"Main.scene",
		)
		err := scene.SaveToFile(savePath)

		if err != nil {
			return errors.New("Failed to save scene")
		}
		return nil
	} else {
		return errors.New("Could not find scene file")
	}
}
