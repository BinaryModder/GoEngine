package hub_ui

import (
	"github.com/AllenDang/giu"
	"goengine/ui/scale"
	"log"
)

var isFontScalingInitialized bool

func Loop() {

	if !isAssetsLoaded {
		if err := LoadTextures(); err != nil {
			log.Fatal("Failed to load hub textures")
		}
		isAssetsLoaded = true

	}
	if !isFontScalingInitialized {
		scale.SetFontScale()
		isFontScalingInitialized = true
	}

	giu.SingleWindow().
		Layout(

			giu.Row(
				Sidebar(),
				MainPanel(),
			),
		)

}
