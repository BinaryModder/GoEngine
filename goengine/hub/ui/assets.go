package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/resources"
	"path/filepath"
)

var (
	Icon           *giu.Texture
	isAssetsLoaded bool
)

func LoadAssets() {
	path, err := filepath.Abs("resources/GoEngineIcon.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("Loading", path, "...")

	resources.LoadTexture(path, func(texture *giu.Texture) {

		Icon = texture

	})

	isAssetsLoaded = true
}
