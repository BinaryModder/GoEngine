package app

import (
	"fmt"
	"os"
	"os/exec"
)

func Editor(path string) {

	exePath, err := os.Executable()

	if err != nil {

		fmt.Println(err)

		return

	}

	cmd := exec.Command(

		exePath,

		"-editor",

		"-project",

		path,
	)

	cmd.Start()

}
