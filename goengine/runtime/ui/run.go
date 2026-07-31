package ui

//
// import (
// 	"fmt"
// 	"github.com/AllenDang/giu"
// 	"goengine/engine/logger"
// 	"goengine/runtime"
// )
//
// func Run(ProjectPath string) {
// 	if ProjectPath == "" {
// 		logger.Error(
// 			"Project path is empty",
// 		)
// 	}
//
// 	runtime.State.ProjectPath = ProjectPath
//
// 	if err := runtime.State.Init(); err != nil {
// 		logger.Error(err.Error())
// 	}
//
// 	window := giu.NewMasterWindow(
// 		fmt.Sprintf(
// 			"Version: %s ; EngineVersion: %s",
// 			runtime.State.ProjectConfig.Version,
// 			runtime.State.ProjectConfig.EngineVersion,
// 		),
// 		1920,
// 		1080,
// 		0,
// 	)
// 	logger.Info("Runtime started")
// 	window.Run(
// 		Loop,
// 	)
// }
