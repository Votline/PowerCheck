package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"PowerCheck/internal/shaders"
)

func Setup() uint32 {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	
	program := gl.CreateProgram()
	err := shaders.CompileAndAttachShaders(program)
	if err != nil {return 0}
	gl.LinkProgram(program)

	gl.LineWidth(3.0)
	gl.UseProgram(program)
	gl.ClearColor(0.0, 0.0, 0.0, 0.9)

	return program
}
