package launcher

import (
	"fmt"
	"os"
	"os/exec"
)

func ProjectRunner(path string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	command := fmt.Sprintf("%q -run_proj -project %q", exePath, path)

	cmd := exec.Command(
		"osascript",
		"-e",
		fmt.Sprintf(
			`tell application "Terminal"
				activate
				do script %q
			end tell`,
			command,
		),
	)

	return cmd.Start()
}
