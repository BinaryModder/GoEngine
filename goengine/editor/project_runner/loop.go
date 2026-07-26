package project_runner

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/engine/platform"
	"goengine/engine/renderer"
	"goengine/ui/scale"
	"log"
)

var (
	isFontScalingInitialized bool
	isRendererInitialized    bool
	isSizesConfigured        bool
	isPlatformInitialized    bool
)

func Loop() {

	if !isFontScalingInitialized {
		scale.SetFontScale()

		isFontScalingInitialized = true

	}

	if !isRendererInitialized {
		if err := renderer.Init(int32(ViewportWeight), int32(ViewportHeight)); err != nil {
			log.Fatalf("Failed to initialize renderer : %v", err)
		}
		isRendererInitialized = true
	}
	if !isPlatformInitialized {
		platform.Init()
		isPlatformInitialized = true
	}

	renderer.Render(&editor.RunProjState)
	imgui.PushStyleVarVec2(
		imgui.StyleVarWindowPadding,
		imgui.Vec2{X: 0, Y: 0},
	)
	defer imgui.PopStyleVar()

	giu.SingleWindow().
		Layout(
			Viewport(),
		)

}
