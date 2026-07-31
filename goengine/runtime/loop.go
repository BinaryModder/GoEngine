package runtime

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"goengine/engine/renderer"
	enginescript "goengine/engine/script"
	"time"
)

func Loop() {

	last := time.Now()

	for !Window.Window.ShouldClose() {

		glfw.PollEvents()

		now := time.Now()

		dt := float32(now.Sub(last).Seconds())

		last = now

		if dt > 0.1 {
			dt = 0.1
		}

		if enginescript.State.HasScripts() {
			enginescript.State.Update(dt)
		}

		renderer.Render(&State)

		Window.Window.SwapBuffers()
	}

	DestroyWindow()
}
