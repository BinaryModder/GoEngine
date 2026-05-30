package main

import (
	"log"

	"github.com/BinaryModder/GoEngine/engine"

	"github.com/AllenDang/giu"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	myShader *engine.Shader
	vao      uint32
	vbo      uint32

	projMatrix mgl32.Mat4
	viewMatrix mgl32.Mat4
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

func loop() {
	gl.ClearColor(0.1, 0.2, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	myShader.Use()

	modelMatrix := mgl32.Ident4()
	mvp := projMatrix.Mul4(viewMatrix).Mul4(modelMatrix)

	mvpLocation := gl.GetUniformLocation(myShader.ProgramID, gl.Str("mvp\x00"))
	gl.UniformMatrix4fv(mvpLocation, 1, false, &mvp[0])

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)

	giu.Window("Inspector").Pos(20, 20).Size(300, 150).Layout(
		giu.Label("GoEngine v0.0 - Editor"),

		giu.Button("Generate Cube").OnClick(func() {
			log.Println("Кнопка работает! Куб уже на фоне.")
		}),
	)
}

func main() {
	w := giu.NewMasterWindow("GoEngine v0.0", 1024, 768, giu.MasterWindowFlagsNotResizable)

	if err := gl.Init(); err != nil {
		log.Fatalln("Не удалось инициализировать OpenGL:", err)
	}

	gl.Enable(gl.DEPTH_TEST)

	var err error
	myShader, err = engine.NewShader(`
		#version 410 core
		layout(location = 0) in vec3 position;
		uniform mat4 mvp;
		void main() {
			gl_Position = mvp * vec4(position, 1.0);
		}
	`, `
		#version 410 core
		out vec4 color;
		void main() {
			color = vec4(0.8, 0.3, 0.2, 1.0); // Ржаво-оранжевый цвет
		}
	`)
	if err != nil {
		log.Fatalln("Ошибка компиляции шейдеров:", err)
	}

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(0)

	projMatrix = mgl32.Perspective(mgl32.DegToRad(45.0), 1024.0/768.0, 0.1, 100.0)
	viewMatrix = mgl32.LookAtV(mgl32.Vec3{2, 2, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

	w.Run(loop)
}
