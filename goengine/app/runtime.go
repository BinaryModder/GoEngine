package app

import (
	"os"
	"os/exec"
)

func Runtime(path string) error {

	exePath, err := os.Executable()

	if err != nil {
		return err
	}

	cmd := exec.Command(
		exePath,
		"-runtime",
		"-project",
		path,
	)

	cmd.Start()

	return nil
}
