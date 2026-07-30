package loader

import (
	"github.com/AllenDang/giu"
	"goengine/hub/ui/assets"
	"goengine/ui/resources"
)

func LoadTextures(path string) error {

	if err := resources.DecodeTextureFile(path, func(texture *giu.Texture) {

		assets.HubTextures.Icon = texture

	}); err != nil {
		return err
	}

	return nil
}
