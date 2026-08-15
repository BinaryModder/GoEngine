package ui

import (
	"github.com/AllenDang/giu"
	"goengine/editor"
	"goengine/editor/ui/loader"
	"goengine/engine/logger"
	"goengine/engine/renderer"
	"goengine/ui/layout"
	"goengine/ui/scale"
)

var (
	isTextureLoaded       bool
	isFontScaled          bool
	isRendererInitialized bool
)

func Loop() {
	processBuildResult()

	if !isTextureLoaded {
		if err := loader.LoadTextures(); err != nil {
			logger.Error(err.Error())
		}
		isTextureLoaded = true
		logger.Info("Textures are loaded")
	}

	if !isFontScaled {
		scale.SetFontScale()
		isFontScaled = true
		logger.Info("Font scale is ready")
	}

	if !isRendererInitialized {
		if err := renderer.Init(int32(layout.ViewportWidth), int32(layout.ViewportHeight)); err != nil {
			logger.Error(err.Error())
		}
		isRendererInitialized = true

		logger.Info("Renderer is initialized. Render in process...")
	}

	renderer.Render(&editor.State)
	giu.SingleWindow().Layout(
		MenuBar(),
		giu.Separator(),
		giu.Row(
			Hierarchy(),
			Viewport(),
			Inspector(),
		),
		giu.Separator(),
		Project(),
		CreateMaterialWindow(),
		LoadMaterialWindow(),
	)
	Console()
}
