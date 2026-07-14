package functions

import (
	"path/filepath"
)

func AbsolutePath(path string) string {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return ""
	}

	return abs_path

}
