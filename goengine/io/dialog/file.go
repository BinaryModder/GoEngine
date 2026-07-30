package dialog

import (
	"github.com/sqweek/dialog"
)

func ChooseImageFile(title string) (string, error) {
	file, err := dialog.File().
		Title(title).
		Filter("Image files", "png", "jpg", "jpeg", "bmp", "tga").
		Load()

	if err != nil {
		return "", err
	}

	return file, nil
}
func ChooseMaterialFile(title string) (string, error) {
	file, err := dialog.File().
		Title(title).
		Filter("Material files", "material").
		Load()

	if err != nil {
		return "", err
	}

	return file, nil
}
