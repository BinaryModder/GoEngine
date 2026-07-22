package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/engine/renderer"
	"goengine/ui/scale"
	"log"
)

var (
	isFontScalingInitialized bool
	isRendererInitialized    bool
	isSizesConfigured        bool
)

func Loop() {

	if !isFontScalingInitialized {
		scale.SetFontScale()

		isFontScalingInitialized = true

	}

	if !isSizesConfigured {
		ConfigureSize()

		isSizesConfigured = true
	}
	if !EditorTextures.IsAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatalf("Failed to load editor textures : %v", err)
		}

		EditorTextures.IsAssetsLoaded = true
	}
	if !isRendererInitialized {
		if err := renderer.Init(int32(ViewportWidth), int32(ViewportHeight)); err != nil {
			log.Fatalf("Failed to initialize renderer : %v", err)
		}
		isRendererInitialized = true
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
