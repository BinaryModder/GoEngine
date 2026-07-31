package runtime

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func CreateWindow(title string, width, height int) error {

	if err := glfw.Init(); err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		return err
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return fmt.Errorf("OpenGL init: %w", err)
	}

	glfw.SwapInterval(1)

	fbWidth, fbHeight := window.GetFramebufferSize()

	Window.Window = window
	Window.Width = fbWidth
	Window.Height = fbHeight
	Window.Running = true

	return nil
}

func DestroyWindow() {

	if Window.Window != nil {
		Window.Window.Destroy()
	}

	glfw.Terminate()
}
