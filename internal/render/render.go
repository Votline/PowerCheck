package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"

	"PowerCheck/internal/ui"
	"PowerCheck/internal/power"
	"PowerCheck/internal/shaders"
)

func Setup() uint32 {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	
	program := gl.CreateProgram()
	err := shaders.CompileAndAttachShaders(program)
	if err != nil {return 0}
	gl.LinkProgram(program)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)

	gl.LineWidth(3.0)
	gl.UseProgram(program)
	gl.ClearColor(0.0, 0.0, 0.0, 0.9)

	return program
}

func Digits() {
	pw := power.Show()
	allV, vQn := ui.GetDigits(pw)
	draw(allV, vQn)
}

func Buttons() {
	allV, vQn := ui.GetButtons()
	draw(allV, vQn)
}

func draw(allV []float32, vQn []int32) {
	gl.BufferData(gl.ARRAY_BUFFER, len(allV)*4, gl.Ptr(allV), gl.STATIC_DRAW)
	start := int32(0)
	for _, vt := range vQn {
		gl.DrawArrays(gl.LINE_STRIP, start, vt)
		start += vt
	}
}
