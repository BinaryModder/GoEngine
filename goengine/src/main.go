package main

import (
	"github.com/AllenDang/cimgui-go/imgui"
	gi "github.com/AllenDang/giu"
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
	"runtime"
)

var (
	fbo        uint32
	fboTexture uint32
	viewWidth  int32 = 800
	viewHeight int32 = 600
)

func init() {
	runtime.LockOSThread()
}

func createFBO(width, height int32) (uint32, uint32) {

	var fbo, texture, rbo uint32

	gl.GenFramebuffers(1, &fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo) // биндим opengGL в кастомный банк памяти

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, width, height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texture, 0)

	gl.GenRenderbuffers(1, &rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, width, height)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		log.Fatalln("Ошибка: FBO не готов к работе!")
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0) // возвращаем рендер на главный экран

	return fbo, texture
}

func loop() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo)
	gl.Viewport(0, 0, viewWidth, viewHeight)

	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	gi.SingleWindow().Layout(
		gi.Label("GoEngine v0.1 (editor)"),
		gi.Button("Add Cube").OnClick(
			func() {
				log.Println("Adding cube...")
			}),
		gi.Custom(
			func() {

				imgui.Image(imgui.TextureRef(*imgui.NewEmptyTextureRef()), imgui.Vec2{X: float32(viewWidth), Y: float32(viewHeight)})

			}),
	)

}
func main() {

	main_window := gi.NewMasterWindow("GoEngine v0.1", 1024, 768, gi.MasterWindowFlagsNotResizable)

	if err := gl.Init(); err != nil {
		log.Fatalln("Can't open OpenGL : ", err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE)

	main_window.Run(loop)
}
