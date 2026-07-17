package hub_ui

import (
	"github.com/AllenDang/giu"
	"goengine/hub/functions"
	"goengine/resources"
)

var (
	Icon           *giu.Texture
	isAssetsLoaded bool
)

func LoadTextures() error {

	path := functions.AbsolutePath("resources/hub/GoEngineIcon.png")

	if err := resources.DecodeTextureFile(path, func(texture *giu.Texture) {

		Icon = texture

	}); err != nil {
		return err
	}

	isAssetsLoaded = true

	return nil
}
