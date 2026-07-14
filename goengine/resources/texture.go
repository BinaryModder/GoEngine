package resources

import (
	"fmt"
	"github.com/AllenDang/giu"
	"image"
	_ "image/png"
	"os"
)

func LoadTexture(path string, callback func(*giu.Texture)) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file not found")
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("file cannot be decoded")
		fmt.Println(err)
		return
	}

	giu.NewTextureFromRgba(
		img,
		callback,
	)

}

