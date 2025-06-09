package shaders

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const vertexShaderSource = `
#version 410 core
layout (location = 0) in vec3 aPos;

void main() {
	gl_Position = vec4(aPos, 1.0);
}` + "\x00"

const fragmentShaderSource = `
#version 410 core
out vec4 FragColor;

void main() {
	FragColor = vec4(1.0, 1.0, 1.0, 0.7);
}` + "\x00"

func CompileAndAttachShaders(program uint32) error {
	deleteShaders(program)

	vertexShader, vertErr := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	fragmentShader, fragErr := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)

	if vertErr != nil {
		return vertErr
	}
	if fragErr != nil {
		return fragErr
	}
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	return nil
}

func checkAttachedShaders(program uint32) (error, int32) {
	var cnt int32
	gl.GetProgramiv(program, gl.ATTACHED_SHADERS, &cnt)
	if cnt < 2 {
		return fmt.Errorf("В программе недостаточно шейдеров: %v", cnt), cnt
	}
	return nil, 0
}

func deleteShaders(program uint32) error {
	err, cnt := checkAttachedShaders(program)
	if cnt > 0 {
		shaders := make([]uint32, cnt)
		gl.GetAttachedShaders(program, cnt, nil, &shaders[0])

		for _, shader := range shaders {
			gl.DetachShader(program, shader)
			gl.DeleteShader(shader)
		}
	}
	return err
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	cSource, freeMemory := gl.Strs(source)
	gl.ShaderSource(shader, 1, cSource, nil)
	freeMemory()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &logMsg[0])
		log.Printf("\nShader compile error. \nType: %d\n%s\n",
			logLength, logMsg)
		err := fmt.Errorf("Shader compile error")
		return 0, err
	}
	return shader, nil
}
