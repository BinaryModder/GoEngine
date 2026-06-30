package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"time"

	"goengine/engine"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/giu"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	isEditorMode  bool
	isInitialized bool

	myShader *engine.Shader
	vao      uint32
	vbo      uint32

	fbo        uint32
	fboTexture uint32
	viewWidth  int32 = 800
	viewHeight int32 = 600

	projMatrix mgl32.Mat4
	viewMatrix mgl32.Mat4
	startTime  time.Time
)

var cubeVertices = []float32{
	-0.5, -0.5, -0.5, 0.5, -0.5, -0.5, 0.5, 0.5, -0.5,
	0.5, 0.5, -0.5, -0.5, 0.5, -0.5, -0.5, -0.5, -0.5,
	-0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, -0.5, 0.5,
	-0.5, 0.5, 0.5, -0.5, 0.5, -0.5, -0.5, -0.5, -0.5,
	-0.5, -0.5, -0.5, -0.5, -0.5, 0.5, -0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, -0.5, 0.5, -0.5, -0.5,
	0.5, -0.5, -0.5, 0.5, -0.5, 0.5, 0.5, 0.5, 0.5,
	-0.5, -0.5, -0.5, 0.5, -0.5, -0.5, 0.5, -0.5, 0.5,
	0.5, -0.5, 0.5, -0.5, -0.5, 0.5, -0.5, -0.5, -0.5,
	-0.5, 0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, 0.5, -0.5,
}

func loopLauncher() {
	giu.SingleWindow().Layout(
		giu.Dummy(0, 100),
		giu.Row(
			giu.Dummy(220, 0),
			giu.Label("GoEngine Hub"),
		),
		giu.Dummy(0, 30),
		giu.Row(
			giu.Dummy(220, 0),
			giu.Button("Create New Project").Size(150, 40).OnClick(func() {
				exePath, err := os.Executable()
				if err != nil {
					log.Fatalln("Не удалось найти исполняемый файл:", err)
				}

				cmd := exec.Command(exePath, "-editor")

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Start(); err != nil {
					log.Fatalln("Не удалось запустить редактор:", err)
				}

				os.Exit(0)
			}),
		),
	)
}

func initEngine() {
	if err := gl.Init(); err != nil {
		log.Fatalln("Не удалось инициализировать OpenGL:", err)
	}

	var err error
	myShader, err = engine.NewShader(`#version 410 core
		layout(location = 0) in vec3 position;
		uniform mat4 mvp;
		void main() {
			gl_Position = mvp * vec4(position, 1.0);
		}
	`, `#version 410 core
		out vec4 color;
		void main() {
			color = vec4(0.8, 0.3, 0.2, 1.0);
		}
	`)
	if err != nil {
		log.Fatalln("Ошибка шейдеров:", err)
	}

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	var rbo uint32
	gl.GenFramebuffers(1, &fbo)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo)

	gl.GenTextures(1, &fboTexture)
	gl.BindTexture(gl.TEXTURE_2D, fboTexture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, viewWidth, viewHeight, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, fboTexture, 0)

	gl.GenRenderbuffers(1, &rbo)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbo)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH24_STENCIL8, viewWidth, viewHeight)
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		log.Fatalln("Ошибка: FBO не готов!")
	}
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	projMatrix = mgl32.Perspective(mgl32.DegToRad(45.0), float32(viewWidth)/float32(viewHeight), 0.1, 100.0)
	viewMatrix = mgl32.LookAtV(mgl32.Vec3{2, 2, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	startTime = time.Now()
}

func loopEditor() {
	if !isInitialized {
		initEngine()
		isInitialized = true
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo)
	gl.Viewport(0, 0, viewWidth, viewHeight)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.ClearColor(0.1, 0.2, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	myShader.Use()

	elapsed := float32(time.Since(startTime).Seconds())
	modelMatrix := mgl32.HomogRotate3D(elapsed, mgl32.Vec3{0.5, 1.0, 0.2}.Normalize())
	mvp := projMatrix.Mul4(viewMatrix).Mul4(modelMatrix)

	mvpLocation := gl.GetUniformLocation(myShader.ProgramID, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvpLocation, 1, false, &mvp[0])

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	gl.BindVertexArray(0)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	giu.SingleWindow().Layout(
		giu.Label("GoEngine v0.0 - Editor Viewport"),
		giu.Custom(func() {
			imgui.Image(imgui.TextureID(fboTexture), imgui.Vec2{X: float32(viewWidth), Y: float32(viewHeight)})
		}),
	)
}

func main() {
	flag.BoolVar(&isEditorMode, "editor", false, "Запустить в режиме редактора")
	flag.Parse()

	if isEditorMode {
		w := giu.NewMasterWindow("GoEngine v0.0 - Editor", 1024, 768, 0)
		w.Run(loopEditor)
	} else {
		w := giu.NewMasterWindow("GoEngine Hub", 600, 400, 0)
		w.Run(loopLauncher)
	}
}
