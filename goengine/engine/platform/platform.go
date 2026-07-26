package platform

import (
	"runtime"
)

type PlatformState struct {
	OS string

	isLinux   bool
	isWindows bool
	isMacOS   bool
}

var State PlatformState

func Init() {
	State.OS = runtime.GOOS

	switch runtime.GOOS {
	case "linux":
		State.isLinux = true
	case "windows":
		State.isWindows = true
	case "macos":
		State.isMacOS = true
	}
}
