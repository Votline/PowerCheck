package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerCheck/internal/digits"
	"PowerCheck/internal/letters"
)

const windowWidth = 210
const windowHeight = 90

func CreateWindow() *glfw.Window  {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "TimeCheck", nil, nil)
	if err != nil {
		log.Println("Create window error. \nErr: ", err)
		window = nil
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	window.SetPos(vidMode.Width-220, vidMode.Height-1075)
	return window
}

func GetDigits(power string) ([]float32, []int32) {
	var (
		offset float32  = 0.0
		allV []float32
		vQn []int32
	)
	for _, ch := range power {
		digits.GetVert(ch, offset, &allV, &vQn)
		offset += 0.2
	}
	return allV, vQn
}

func GetButtons() ([]float32, []int32) {
	var (
		offset float32 = -0.02
		width float32 = 0.0
		s string = "SD SS"
		
		allV []float32
		vQn []int32
	)

	for _, ch := range s {
		letters.GetVert(ch, offset, &allV, &vQn, &width)
		offset += width
	}
	return allV, vQn
}
