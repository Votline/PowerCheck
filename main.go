package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"


	"PowerCheck/internal/ui"
	"PowerCheck/internal/render"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()
	
	window := ui.CreateWindow()
	if window == nil {return}

	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}
	program := render.Setup()
	if program == 0 {return}

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		
		render.Digits()
		render.Buttons()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
