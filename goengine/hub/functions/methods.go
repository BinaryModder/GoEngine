package functions

import (
	"path/filepath"
)

const pathSize int = 38

func AbsolutePath(path string) string {

	abs_path, err := filepath.Abs(path)

	if err != nil {
		return ""
	}

	return abs_path

}

func ConfigureLabelPath(path string) string {
	if len(path) <= pathSize {
		return path
	}

	return "..." + path[len(path)-pathSize:]

}
