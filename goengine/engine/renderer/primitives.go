package renderer

import "github.com/go-gl/gl/v4.1-core/gl"

const MeshVertexShader = `
#version 410 core
layout (location = 0) in vec3 aPos;

uniform mat4 model;      // Матрица трансформации (Позиция, Поворот, Масштаб)
uniform mat4 view;       // Камера
uniform mat4 projection; // Перспектива

out vec3 FragPos; // Передаем позицию во фрагментный шейдер для расчета цвета

void main() {
    FragPos = aPos;
    // Умножаем справа налево: Вершина -> Мир -> Камера -> Экран
    gl_Position = projection * view * model * vec4(aPos, 1.0);
}
`

const MeshFragmentShader = `
#version 410 core
in vec3 FragPos;
out vec4 FragColor;

void main() {
    vec3 baseColor = vec3(0.8, 0.5, 0.2); // Оранжевый цвет
    float fakeLight = (FragPos.y + 0.5);  // Градиент от -0.5 до +0.5
    FragColor = vec4(baseColor * fakeLight, 1.0);
}
`

var (
	MeshProgram uint32
	CubeVAO     uint32
	CubeVBO     uint32
)

func InitPrimitives() {
	var err error
	MeshProgram, err = CreateShaderProgram(MeshVertexShader, MeshFragmentShader)
	if err != nil {
		panic(err)
	}

	cubeVertices := []float32{
		// Задняя грань
		-0.5, -0.5, -0.5, 0.5, -0.5, -0.5, 0.5, 0.5, -0.5,
		0.5, 0.5, -0.5, -0.5, 0.5, -0.5, -0.5, -0.5, -0.5,
		// Передняя грань
		-0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, 0.5, 0.5,
		0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, -0.5, 0.5,
		// Левая грань
		-0.5, 0.5, 0.5, -0.5, 0.5, -0.5, -0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5, -0.5, -0.5, 0.5, -0.5, 0.5, 0.5,
		// Правая грань
		0.5, 0.5, 0.5, 0.5, 0.5, -0.5, 0.5, -0.5, -0.5,
		0.5, -0.5, -0.5, 0.5, -0.5, 0.5, 0.5, 0.5, 0.5,
		// Нижняя грань
		-0.5, -0.5, -0.5, 0.5, -0.5, -0.5, 0.5, -0.5, 0.5,
		0.5, -0.5, 0.5, -0.5, -0.5, 0.5, -0.5, -0.5, -0.5,
		// Верхняя грань
		-0.5, 0.5, -0.5, 0.5, 0.5, -0.5, 0.5, 0.5, 0.5,
		0.5, 0.5, 0.5, -0.5, 0.5, 0.5, -0.5, 0.5, -0.5,
	}

	gl.GenVertexArrays(1, &CubeVAO)
	gl.GenBuffers(1, &CubeVBO)

	gl.BindVertexArray(CubeVAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, CubeVBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)
}
