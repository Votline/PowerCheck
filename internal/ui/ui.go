package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
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
	return window
}
