package runtime

import "github.com/go-gl/glfw/v3.3/glfw"

func IsKeyPressed(key glfw.Key) bool {

	return Window.Window.GetKey(key) == glfw.Press
}

func MousePosition() (float64, float64) {

	return Window.Window.GetCursorPos()
}
