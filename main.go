package main

import (
	"log"
	"flag"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerCheck/internal/ui"
	"PowerCheck/internal/power"
	"PowerCheck/internal/render"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	smode := flag.Bool("smode", false, "Silence mode")
	flag.Parse()

	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()
	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}

	pm := power.NewPM()
	go pm.Timer(smode)
	for {
		win := ui.CreateWin(200, 90, 220, 1075)
		pg, ofl := render.Setup()
		pc := ui.CreatePC(pg, ofl)
		win.SetMouseButtonCallback(ui.BtnCallback(pc, 200, 90))

		if !*smode {
			win.Show()
		}

		for !win.ShouldClose() {
			select {
			case <-pm.Pch:
				win.Show()
				renderFrame(win, pc)
			default:
			}
			renderFrame(win, pc)
		}
		win.Destroy()
		if !*smode {
			return
		}
	}
}

func renderFrame(win *glfw.Window, pc *ui.PowerChecker) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	
	pc.Render(win)
	
	win.SwapBuffers()
	glfw.PollEvents()
}
