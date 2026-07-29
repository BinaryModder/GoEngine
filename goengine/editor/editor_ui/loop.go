package editor_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/engine/platform"
	"goengine/engine/renderer"
	"goengine/ui/scale"
	"log"
)

// Some flags for Initializing
var (
	isFontScalingInitialized bool
	isRendererInitialized    bool
	isSizesConfigured        bool
	isPlatformInitialized    bool
)

// The centre of Editor Interface
func Loop() {
	//Platform information initializing
	if !isPlatformInitialized {
		platform.Init()
		isPlatformInitialized = true
	}

	//Loading Font
	if !isFontScalingInitialized {
		scale.SetFontScale()

		isFontScalingInitialized = true

	}

	//Configuring sizes
	if !isSizesConfigured {
		ConfigureSize()

		isSizesConfigured = true
	}

	//All Textures loading
	if !EditorTextures.IsAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatalf("Failed to load editor textures : %v", err)
		}

		EditorTextures.IsAssetsLoaded = true
	}

	//Render Initilizing
	if !isRendererInitialized {
		if err := renderer.Init(int32(ViewportWidth), int32(ViewportHeight)); err != nil {
			log.Fatalf("Failed to initialize renderer : %v", err)
		}
		isRendererInitialized = true
	}

	renderer.Render(&editor.State) // Render with Editor mode

	//Connecting all widgets
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
