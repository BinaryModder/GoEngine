package functions

import (
	"errors"
	"fmt"
	"github.com/sqweek/dialog"
	"os/exec"
	"runtime"
)

func FileMenuBar() error {

	path, err := dialog.Directory().
		Title("Choose GoEngine Project").
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
