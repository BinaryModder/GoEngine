package renderer

import "github.com/go-gl/gl/v4.1-core/gl"

const MeshVertexShader = `
#version 410 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec3 FragPos;
out vec2 TexCoord;

void main() {
    FragPos = aPos;
    TexCoord = aTexCoord;
    gl_Position = projection * view * model * vec4(aPos, 1.0);
}
`

const MeshFragmentShader = `
#version 410 core

in vec3 FragPos;
in vec2 TexCoord;
out vec4 FragColor;

uniform vec3 objectColor;
uniform sampler2D uTexture;
uniform bool uUseTexture;

void main() {
    float fakeLight = (FragPos.y + 0.5) * 0.3 + 0.7;
    vec4 color = vec4(objectColor, 1.0);
    if (uUseTexture) {
        color *= texture(uTexture, TexCoord);
    }
    FragColor = vec4(color.rgb * fakeLight, 1.0);
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

	stride := int32(5 * 4)

	gl.VertexAttribPointer(
		0,
		3,
		gl.FLOAT,
		false,
		stride,
		gl.PtrOffset(0),
	)
	gl.EnableVertexAttribArray(0)

	gl.VertexAttribPointer(
		1,
		2,
		gl.FLOAT,
		false,
		stride,
		gl.PtrOffset(3*4),
	)
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(0)

	p.VertexCount = int32(len(vertices) / 5)
}

func CreateCube() {

	// format: px, py, pz, u, v
	cubeVertices := []float32{
		// Front face (z = -0.5)
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		// Back face (z = 0.5)
		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		// Left face (x = -0.5)
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 0.0,
		-0.5, -0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,

		// Right face (x = 0.5)
		0.5, 0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0,

		// Bottom face (y = -0.5)
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		// Top face (y = 0.5)
		-0.5, 0.5, -0.5, 0.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0,
	}

	InitPrimitive(&Cube, cubeVertices)
}

func CreatePyramid() {

	// format: px, py, pz, u, v
	pyramidVertices := []float32{
		// Front-right face
		0.0, 0.5, 0.0, 0.5, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,

		// Right-back face
		0.0, 0.5, 0.0, 0.5, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,

		// Back-left face
		0.0, 0.5, 0.0, 0.5, 1.0,
		0.5, -0.5, -0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 1.0, 0.0,

		// Left-front face
		0.0, 0.5, 0.0, 0.5, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,
		-0.5, -0.5, 0.5, 1.0, 0.0,

		// Base
		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 1.0,

		0.5, -0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
	}

	InitPrimitive(&Pyramid, pyramidVertices)
}
