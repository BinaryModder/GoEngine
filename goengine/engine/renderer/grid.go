package renderer

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var (
	gridProgram uint32
	gridVAO     uint32
	gridVBO     uint32
)

const GridVertexShader = `
#version 410 core
layout (location = 0) in vec3 aPos;

uniform mat4 view;
uniform mat4 projection;

out vec3 FragPos;

void main() {
    FragPos = aPos; 
    gl_Position = projection * view * vec4(aPos, 1.0);
}
`
const GridFragmentShader = `
#version 410 core
out vec4 FragColor;
in vec3 FragPos;

void main() {
    vec2 coord = FragPos.xz;
    vec2 derivative = fwidth(coord);
    vec2 grid = abs(fract(coord - 0.5) - 0.5) / derivative;
    float line = min(grid.x, grid.y);
    float alpha = 1.0 - min(line, 1.0);

    float dist = length(FragPos);
    alpha *= max(0.0, 1.0 - (dist / 30.0)); 

    vec3 color = vec3(0.3); 

    if(abs(FragPos.x) < 0.05) color = vec3(0.0, 0.0, 1.0); // Ось Z - синяя
    if(abs(FragPos.z) < 0.05) color = vec3(1.0, 0.0, 0.0); // Ось X - красная

    FragColor = vec4(color, alpha);
}
`

func InitGrid() {
	var err error
	gridProgram, err = CreateShaderProgram(GridVertexShader, GridFragmentShader)
	if err != nil {
		log.Fatalf("Failed to compile grid shader: %v", err)
	}

	var size float32 = 50.0
	vertices := []float32{
		-size, 0.0, -size,
		size, 0.0, -size,
		size, 0.0, size,

		size, 0.0, size,
		-size, 0.0, size,
		-size, 0.0, -size,
	}

	gl.GenVertexArrays(1, &gridVAO)
	gl.BindVertexArray(gridVAO)

	gl.GenBuffers(1, &gridVBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, gridVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)
}
