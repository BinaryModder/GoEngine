package main

import (
	"log"
	"os"
	"path/filepath"

	"goengine/engine/logger"
	"goengine/runtime"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	executablePath, err := os.Executable()
	if err != nil {
		logger.Error("Failed to locate executable: " + err.Error())
		return
	}

	runtime.Run(filepath.Dir(executablePath))
}
