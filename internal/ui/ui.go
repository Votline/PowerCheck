package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type element interface{
	create() element

	getVtc() []float32
	getVtq() int32
	
	setData(rune)
}
type digit struct{
	rn rune
	vtq int32
	vtc []float32
}
type letter struct{
	rn rune
	vtq int32
	vtc []float32
}
type btn struct {
	vtq int32
	vtc []float32
	pos [4]float32
	Text string
}

const winW = 200
const winH = 90

func CreateWin() *glfw.Window  {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	win, err := glfw.CreateWindow(winW, winH, "PowerChecker", nil, nil)
	if err != nil {
		log.Println("Create win error. \nErr: ", err)
		win = nil
	}
	win.MakeContextCurrent()
	win.SetAttrib(glfw.Floating, 1)

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	win.SetPos(vidMode.Width-220, vidMode.Height-1075)

	glfw.SwapInterval(1)
	return win
}
