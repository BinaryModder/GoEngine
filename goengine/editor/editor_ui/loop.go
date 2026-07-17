package editor_ui

import (
	"github.com/AllenDang/giu"
	"log"
)

func Loop() {

	if !EditorTextures.IsAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatal("Failed to load editor textures")
		}
	}
	giu.SingleWindow().Layout(
		MenuBar(),
		//Toolbar(),
		giu.Separator(),
		giu.Row(
			ErrorMessage(),

			Hierarchy(),
			Viewport(),
			Inspector(),
		),
		giu.Separator(),
		Project(),
	)
}
