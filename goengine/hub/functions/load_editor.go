package functions

import (
	"fmt"
	"os"
	"os/exec"
)

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
