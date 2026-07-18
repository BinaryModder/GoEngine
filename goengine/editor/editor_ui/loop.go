package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/engine/renderer"
	"log"
)

var IsRendererInitialized bool

func Loop() {

	if !EditorTextures.IsAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatalf("Failed to load editor textures : %v", err)
		}

		EditorTextures.IsAssetsLoaded = true
	}
	if !IsRendererInitialized {
		if err := renderer.Init(int32(ViewportWidth), int32(ViewportHeight)); err != nil {
			log.Fatalf("Failed to initialize renderer : %v", err)
		}
		IsRendererInitialized = true
	}

	renderer.Render()
	giu.SingleWindow().Layout(
		MenuBar(),
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
