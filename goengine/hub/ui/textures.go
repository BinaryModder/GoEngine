package ui

import (
	"github.com/AllenDang/giu"
	"goengine/core/filesystem"
	"goengine/ui/resources"
)

var (
	Icon           *giu.Texture
	isAssetsLoaded bool
)

func LoadTextures() error {

	path := filesystem.AbsolutePath("ui/resources/hub/GoEngineIcon.png")

	if err := resources.DecodeTextureFile(path, func(texture *giu.Texture) {

		Icon = texture

	}); err != nil {
		return err
	}

	return nil
}
