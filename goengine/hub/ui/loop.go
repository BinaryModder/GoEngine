package ui

import (
	"fmt"
	"github.com/AllenDang/giu"
	"goengine/core/filesystem"
	"goengine/engine/logger"
	"goengine/hub/ui/loader"
	"goengine/ui/scale"
)

var (
	isTextureLoaded bool
	isFontScaled    bool
)

// The centre of HUB Interface
func Loop() {

	// Loading Assets
	if !isTextureLoaded {

		icon_path := filesystem.AbsolutePath("ui/resources/hub/GoEngineIcon.png")
		if err := loader.LoadTextures(icon_path); err != nil {
			logger.Error(fmt.Sprintf("Failed to load hub textures: %v", err))
		}

		isTextureLoaded = true
		logger.Info("Textures are loaded")

	}
	// Loading Fonts

	if !isFontScaled {
		scale.SetFontScale()

		isFontScaled = true
		logger.Info("Scaling is initialized")

	}

	//Connecting all widgets
	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

	Console()
}
