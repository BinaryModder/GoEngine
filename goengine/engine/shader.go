package engine

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	ProgramID uint32
}

func NewShader(vertexSource, fragmentSource string) (*Shader, error) {
	vertexShader, err := compile(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compile(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	programID := gl.CreateProgram()
	gl.AttachShader(programID, vertexShader)
	gl.AttachShader(programID, fragmentShader)
	gl.LinkProgram(programID)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{ProgramID: programID}, nil
}

func (s *Shader) Use() {
	gl.UseProgram(s.ProgramID)
}

func compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile shader: %v", log)
	}
	return shader, nil
}
