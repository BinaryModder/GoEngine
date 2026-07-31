package ui

import (
	"fmt"
	"time"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"goengine/engine/console"
	"goengine/engine/logger"
	"goengine/engine/platform"
	"goengine/engine/renderer"
	enginescript "goengine/engine/script"
	"goengine/runtime"
	"goengine/ui/scale"
)

var (
	isFontScalingInitialized   bool
	isRendererInitialized      bool
	isPlatformInitialized      bool
	isScriptManagerInitialized bool
	lastFrameTime              time.Time
)

func Loop() {

	if !isFontScalingInitialized {
		scale.SetFontScale()
		isFontScalingInitialized = true
	}

	if !isRendererInitialized {
		if err := renderer.Init(1920, 1080); err != nil {
			console.State.Error(fmt.Sprintf("Failed to initialize renderer : %v", err))
		}
		renderer.SetUseSceneCamera(true)
		isRendererInitialized = true
		console.State.Info("Renderer Succesfuly Initialized")
	}

	if !isPlatformInitialized {
		platform.Init()
		isPlatformInitialized = true
		console.State.Info("Platform Succesfuly Initialized")
	}

	if !isScriptManagerInitialized {
		enginescript.State.Init(runtime.State.ProjectPath)
		enginescript.State.LoadAndStartAll(runtime.State.CurrentScene)
		isScriptManagerInitialized = true
		lastFrameTime = time.Now()
		logger.Info("Script Manager Succesfuly Initialized")
	}

	now := time.Now()
	dt := float32(now.Sub(lastFrameTime).Seconds())
	lastFrameTime = now

	if dt > 0.1 {
		dt = 0.1
	}

	if dt > 0 && enginescript.State.HasScripts() {
		enginescript.State.Update(dt)
	}

	renderer.Render(&runtime.State)
	imgui.PushStyleVarVec2(
		imgui.StyleVarWindowPadding,
		imgui.Vec2{X: 0, Y: 0},
	)
	defer imgui.PopStyleVar()

	giu.SingleWindow().
		Layout(
			Viewport(),
			giu.Separator(),
		)
	Console()
}
