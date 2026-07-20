package renderer

import "github.com/go-gl/gl/v4.1-core/gl"

const MeshVertexShader = `
#version 410 core
layout (location = 0) in vec3 aPos;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec3 FragPos;

void main() {
    FragPos = aPos;
    gl_Position = projection * view * model * vec4(aPos, 1.0);
}
`

const MeshFragmentShader = `
#version 410 core

in vec3 FragPos;
out vec4 FragColor;
uniform vec3 objectColor;
	
void main() {
    float fakeLight = (FragPos.y + 0.5) * 0.3 + 0.7;
    FragColor = vec4(objectColor * fakeLight, 1.0);
}
`

type Primitive struct {
	VAO uint32
	VBO uint32

	VertexCount int32
}

var (
	MeshProgram uint32

	Cube    Primitive
	Pyramid Primitive
	Sphere  Primitive
)

func InitPrimitives() {
	var err error

	MeshProgram, err = CreateShaderProgram(
		MeshVertexShader,
		MeshFragmentShader,
	)
	if err != nil {
		panic(err)
	}

	CreateCube()
	CreatePyramid()
}

func InitPrimitive(p *Primitive, vertices []float32) {
	gl.GenVertexArrays(1, &p.VAO)
	gl.GenBuffers(1, &p.VBO)

	gl.BindVertexArray(p.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, p.VBO)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(vertices)*4,
		gl.Ptr(vertices),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointer(
		0,
		3,
		gl.FLOAT,
		false,
		3*4,
		gl.PtrOffset(0),
	)

	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)

	p.VertexCount = int32(len(vertices) / 3)
}

func CreateCube() {

	cubeVertices := []float32{
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

	InitPrimitive(&Cube, cubeVertices)
}

func CreatePyramid() {

	pyramidVertices := []float32{

		0.0, 0.5, 0.0,
		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,

		0.0, 0.5, 0.0,
		0.5, -0.5, 0.5,
		0.5, -0.5, -0.5,

		0.0, 0.5, 0.0,
		0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,

		0.0, 0.5, 0.0,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, 0.5,

		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, -0.5, -0.5,

		0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, 0.5,
	}

	InitPrimitive(&Pyramid, pyramidVertices)
}
