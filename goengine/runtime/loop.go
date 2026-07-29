package runtime

import (
	"fmt"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"goengine/engine/console"
	"goengine/engine/platform"
	"goengine/engine/renderer"
	"goengine/ui/scale"
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
			console.State.Error(fmt.Sprint("Failed to initialize renderer : %v", err))
		}
		isRendererInitialized = true
		console.State.Info("Renderer Succesfuly Initialized")

	}
	if !isPlatformInitialized {
		platform.Init()
		isPlatformInitialized = true
		console.State.Info("Platform Succesfuly Initialized")

	}

	renderer.Render(&State)
	imgui.PushStyleVarVec2(
		imgui.StyleVarWindowPadding,
		imgui.Vec2{X: 0, Y: 0},
	)
	defer imgui.PopStyleVar()

	giu.SingleWindow().
		Layout(
			Viewport(),
			giu.Separator(),
			Console(),
		)

}
