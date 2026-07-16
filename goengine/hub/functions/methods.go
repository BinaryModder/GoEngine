package functions

import (
	"github.com/sqweek/dialog"
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
func ChooseFolder() (string, error) {

	folder, err := dialog.Directory().Title("Choose Project path").Browse()

	if err != nil {
		return "", err
	}

	return folder, nil

}
