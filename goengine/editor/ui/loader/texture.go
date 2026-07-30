package loader

import (
	"github.com/AllenDang/giu"
	"goengine/core/filesystem"
	"goengine/editor/ui/assets"
	"goengine/ui/resources"
	"os"
	"path/filepath"
)

func LoadTextures() error {
	path := filesystem.AbsolutePath("ui/resources/editor")

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
				assets.EditorTextures.FolderEmptyTexture = curr_texture

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
				assets.EditorTextures.FileTexture = curr_texture

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
				assets.EditorTextures.FolderContainingTexture = curr_texture

			}); err != nil {
				return err
			}

		}

	}

	return nil

}
