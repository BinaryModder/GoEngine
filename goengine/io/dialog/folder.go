package dialog

import (
	"github.com/sqweek/dialog"
)

// Choosing folder dialog
func ChooseFolder() (string, error) {

	folder, err := dialog.Directory().Title("Choose Project path").Browse()

	if err != nil {
		return "", err
	}

	return folder, nil

}

// Choosing project folder dialog

func ChooseProjectDialog(title string) (string, error) {
	if title != "" {
		path, err := dialog.Directory().
			Title(title).
			Browse()

		if err != nil {
			return "", err
		}

		return path, nil

	} else {
		path, err := dialog.Directory().
			Browse()

		if err != nil {
			return "", err
		}

		return path, nil

	}

}
