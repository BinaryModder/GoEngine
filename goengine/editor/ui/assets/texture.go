package assets

import (
	"github.com/AllenDang/giu"
)

type EditorTexture struct {
	IsAssetsLoaded bool

	FolderContainingTexture *giu.Texture
	FolderEmptyTexture      *giu.Texture
	FileTexture             *giu.Texture
}

var (
	EditorTextures EditorTexture
)
