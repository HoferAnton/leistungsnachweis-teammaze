package display

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.2-core/gl"
)

func CreateProgram(vertexShaderFile, fragmentShaderFile string) (uint32, error) {
	vertexShader, err := shaderFromFile(vertexShaderFile, gl.VERTEX_SHADER)

	FatalIfError("Could not create vertex shader: ", err)

	fragmentShader, err := shaderFromFile(fragmentShaderFile, gl.FRAGMENT_SHADER)

	FatalIfError("Could not create fragment shader: ", err)

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func shaderFromFile(file string, shaderType uint32) (uint32, error) {
	source, err := ioutil.ReadFile(file)

	FatalIfError("Could not read shader source: ", err)

	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(string(source))
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

		return 0, fmt.Errorf("failed to compile %v: %v", file, log)
	}

	return shader, nil
}
