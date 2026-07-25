package run_proj_ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
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

	if !isRendererInitialized {
		if err := renderer.Init(int32(ViewportWeght), int32(ViewportHeight)); err != nil {
			log.Fatalf("Failed to initialize renderer : %v", err)
		}
		isRendererInitialized = true
	}

	renderer.Render(&editor.RunProjState)

	giu.SingleWindow().Layout(
		Viewport(),
	)
}
