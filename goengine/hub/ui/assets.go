package ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub/functions"
	"goengine/resources"
)

var (
	Icon           *giu.Texture
	isAssetsLoaded bool
)

func LoadAssets() {

	path := functions.AbsolutePath("resources/GoEngineIcon.png")

	resources.LoadTexture(path, func(texture *giu.Texture) {

		Icon = texture

	})

	isAssetsLoaded = true
}
