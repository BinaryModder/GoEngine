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
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
