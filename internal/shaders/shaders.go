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
}`

const fragmentShaderSource = `
#version 410 core
out vec4 FragColor;

void main() {
	FragColor = vec4(1.0, 1.0, 1.0, 0.7);
}`


func CompileAndAttachShaders(program uint32) error {
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
