package main

import (
	"log"
	"flag"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerChecker/internal/ui"
	"PowerChecker/internal/power"
	"PowerChecker/internal/render"
	"PowerChecker/internal/config"
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
	log.SetFlags(log.Lshortfile)

	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("Config getting error: %v", err)
	}

	pm := power.NewPM()
	go pm.Timer(smode)
	for {
		win := ui.CreateWin(cfg.WinW, cfg.WinH, cfg.AlX, cfg.AlY)
		pg, ofL, ofS := render.Setup(cfg.TextC, cfg.BackC)
		pc := ui.CreatePC(pg, ofL, ofS)
		win.SetMouseButtonCallback(ui.BtnCallback(pc, cfg.WinW, cfg.WinH))

		if !*smode {
			win.Show()
		}

		for !win.ShouldClose() {
			select {
			case <-pm.Pch:
				win.Show()
			default:
				glfw.WaitEvents()
			}
			if win.GetAttrib(glfw.Visible) == glfw.True {
				renderFrame(win, pc, cfg.WinW, cfg.WinH)
			}
		}
		win.Destroy()
		if !*smode {
			return
		}
	}
}

func renderFrame(win *glfw.Window, pc *ui.PowerChecker, winW, winH int) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	pc.Render(win, winW, winH)
	win.SwapBuffers()
}
