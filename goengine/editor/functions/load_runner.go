package functions

import (
	"os"
	"os/exec"
)

func ProjectRunner(path string) error {

	exePath, err := os.Executable()

	if err != nil {
		return err
	}

	cmd := exec.Command(
		exePath,
		"-run_proj",
		"-project",
		path,
	)

	cmd.Start()

	return nil
}
