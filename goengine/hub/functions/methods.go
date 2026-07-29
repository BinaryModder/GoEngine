package functions

import (
	"github.com/sqweek/dialog"
)

const pathSize int = 38

// If absolute path is too long , this function should make it like (.../ProjectFolder/ProjectName)
func ConfigureLabelPath(path string) string {
	if len(path) <= pathSize {
		return path
	}

	return "..." + path[len(path)-pathSize:]

}

// Choosing project folder dialog
func ChooseFolder() (string, error) {

	folder, err := dialog.Directory().Title("Choose Project path").Browse()

	if err != nil {
		return "", err
	}

	return folder, nil

}
