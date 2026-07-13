package hub

import (
	"fmt"
	"github.com/sqweek/dialog"
	"os"
	"os/exec"
)

func LoadProject() {

	path, err := dialog.Directory().
		Title("Choose GoEngine Project").
		Browse()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(
		"Loading project:",
		path,
	)

	openEditor(path)

}

func openEditor(path string) {

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
