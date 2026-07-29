package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// Running editor mode
func OpenEditor(path string) {

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
