package app

import (
	"fmt"
	"os"
	"os/exec"
)

func Hub() {

	exePath, err := os.Executable()

	if err != nil {

		fmt.Println(err)

		return

	}

	cmd := exec.Command(

		exePath,
	)

	cmd.Start()

}
