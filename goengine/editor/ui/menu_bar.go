package ui

import (
	"fmt"

	"github.com/AllenDang/giu"
	"goengine/app"
	projectbuild "goengine/build"
	"goengine/editor"
	"goengine/editor/functions"
	"goengine/engine/logger"
	enginedialog "goengine/io/dialog"
	"goengine/ui/layout"
)

type buildResult struct {
	path string
	err  error
}

var (
	isBuilding   bool
	buildResults = make(chan buildResult, 1)
)

func MenuBar() giu.Widget {

	var selectedObjectIndex int32 = -1   // 0 - cube , 1 - pyramid , 2 - camera
	var selectedMaterialIndex int32 = -1 // 0 - new material (png)
	return giu.Row(
		//Left Part
		giu.Button("Save").OnClick(
			func() {
				if err := functions.SaveMenuBar(editor.State.CurrentScene, editor.State.ProjectPath); err != nil {
					logger.Error(err.Error())
					return
				}
			},
		).Size(layout.SaveeditSizeWeight, layout.SaveeditSizeHeight),

		giu.Button("Edit").
			Size(layout.SaveeditSizeWeight, layout.SaveeditSizeHeight),

		giu.Button("Assets").OnClick(
			func() {
				if err := functions.AssetMenuBar(editor.State.DefaultAssetsFolder); err != nil {
					logger.Error(err.Error())
					return
				}
			},
		),

		giu.Combo("", "SceneObj", []string{"Cube", "Pyramid", "Camera"}, &selectedObjectIndex).Size(120).
			OnChange(func() {

				if err := functions.SceneObjectMenuBar(editor.State.CurrentScene, &selectedObjectIndex); err != nil {
					logger.Error(err.Error())
					return
				}
				logger.Info("New object created")

			}),

		giu.Combo("", "Material", []string{"New material", "Load material"}, &selectedMaterialIndex).Size(120).
			OnChange(func() {

				switch selectedMaterialIndex {
				case 0:
					editor.State.ShowCreateMaterial = true
				case 1:

					editor.State.ShowLoadMaterial = true
				}

			}),

		giu.Button("Window"),

		giu.Button("Help"),

		giu.Button(buildButtonLabel()).OnClick(startBuild),

		//Middle Part
		giu.Button("Run").OnClick(
			func() {
				if err := app.Runtime(editor.State.ProjectPath); err != nil {
					logger.Error(err.Error())
					return
				}
				logger.Info("Runtime starting...")
			},
		).Size(layout.RunSizeWeight, layout.RunSizeHeight),
	)
}

func buildButtonLabel() string {
	if isBuilding {
		return "Building..."
	}
	return "Build"
}

func startBuild() {
	if isBuilding {
		logger.Warning("Build is already in progress")
		return
	}
	if editor.State.ProjectConfig == nil {
		logger.Error("Cannot build project: project configuration is not loaded")
		return
	}
	if err := functions.SaveMenuBar(editor.State.CurrentScene, editor.State.ProjectPath); err != nil {
		logger.Error("Cannot save project before build: " + err.Error())
		return
	}

	outputPath, err := enginedialog.ChooseProjectDialog("Choose build output folder")
	if err != nil {
		logger.Warning("Build cancelled: " + err.Error())
		return
	}
	logger.Info("Build output selected: " + outputPath)

	isBuilding = true
	editor.State.ShowConsole = true
	options := projectbuild.Options{
		ProjectPath: editor.State.ProjectPath,
		OutputPath:  outputPath,
		Name:        editor.State.ProjectConfig.Name,
	}

	go func() {
		defer func() {
			if recovered := recover(); recovered != nil {
				err := fmt.Errorf("build panic: %v", recovered)
				logger.Error("Build failed: " + err.Error())
				buildResults <- buildResult{err: err}
			}
		}()

		result, err := projectbuild.Project(options)
		if err != nil {
			logger.Error("Build failed: " + err.Error())
			buildResults <- buildResult{err: err}
			return
		}
		buildResults <- buildResult{path: result.ExecutablePath}
	}()
}

func processBuildResult() {
	select {
	case result := <-buildResults:
		isBuilding = false
		if result.err == nil {
			logger.Info("Executable is ready: " + result.path)
		}
	default:
	}
}
