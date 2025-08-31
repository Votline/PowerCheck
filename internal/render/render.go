package render

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

func Setup() (uint32, int32) {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	
	pg := gl.CreateProgram()
	shaders := attachShaders(pg)

	gl.LinkProgram(pg)
	gl.UseProgram(pg)

	ofl := gl.GetUniformLocation(pg, gl.Str("xOffset\x00"))
	detachShaders(pg, shaders)

	gl.ClearColor(0.0, 0.0, 0.0, 0.7)

	return pg, ofl
}

func CreateVAO(vtc []float32) uint32 {
	var vao, vbo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)

	gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(vtc)*4, gl.Ptr(vtc),gl.STATIC_DRAW)

		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
		gl.EnableVertexAttribArray(0)
	gl.BindVertexArray(0)

	return vao
}

func ElemRender(pg uint32, ofl int32, vao uint32, vtq int32, offset float32) {
	gl.UseProgram(pg)
	
	gl.Uniform1f(ofl, offset)
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.LINE_STRIP, 0, vtq)
}
