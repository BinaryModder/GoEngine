package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor/functions"
	"goengine/ui/resources"
	"os"
	"path/filepath"
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

func LoadTextures() error {
	path := functions.AbsolutePath("ui/resources/editor")

	files, err := os.ReadDir(path)

	if err != nil {
		return err
	}

	for _, texture := range files {

		if texture.Name() == "FolderEmptyIcon.png" {
			folder_icon_path := filepath.Join(
				path,
				"FolderEmptyIcon.png",
			)
			if err := resources.DecodeTextureFile(folder_icon_path, func(curr_texture *giu.Texture) {
				EditorTextures.FolderEmptyTexture = curr_texture

			}); err != nil {
				return err
			}

		}

		if texture.Name() == "FileIcon.png" {
			file_icon_path := filepath.Join(
				path,
				"FileIcon.png",
			)

			if err := resources.DecodeTextureFile(file_icon_path, func(curr_texture *giu.Texture) {
				EditorTextures.FileTexture = curr_texture

			}); err != nil {
				return err
			}

		}

		if texture.Name() == "FolderContainingIcon.png" {
			folder_icon_path := filepath.Join(
				path,
				"FolderContainingIcon.png",
			)
			if err := resources.DecodeTextureFile(folder_icon_path, func(curr_texture *giu.Texture) {
				EditorTextures.FolderContainingTexture = curr_texture

			}); err != nil {
				return err
			}

		}

	}

	return nil

}
