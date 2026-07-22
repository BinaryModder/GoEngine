package resources

import (
	"github.com/AllenDang/giu"
	"image"
	_ "image/png"
	"os"
)

func DecodeTextureFile(path string, callback func(*giu.Texture)) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	giu.NewTextureFromRgba(
		img,
		callback,
	)
	return nil

}
