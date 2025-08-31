package ui

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func (b *btn) create() element {
	return &btn{
		vtc: []float32{
			b.pos[0], b.pos[1], 0.0,
			b.pos[0], b.pos[3], 0.0,

			b.pos[2], b.pos[3], 0.0,
			b.pos[2], b.pos[1], 0.0,
			b.pos[0], b.pos[1], 0.0},
		vtq: int32(5)}
}

func (b *btn) getVtq() int32 {
	return b.vtq
}

func (b *btn) getVtc() []float32 {
	return b.vtc
}

func (b *btn) setPos(pos [4]float32) {
	b.pos = pos
}

func (b *btn) setText(t string) {
	b.Text = t
}

func (b *btn) Hover(w *glfw.Window, winW, winH int) (bool) {
	mX, mY := w.GetCursorPos()
	glX := float32(mX)/float32(winW)*2-1
	glY := 1-float32(mY)/float32(winH)*2

	if glX >= b.pos[0] && glX <= b.pos[2] && glY <= b.pos[1] && glY >= b.pos[3] {
		return true
	}

	return false
}

func (b *btn) setData(r rune) {}
