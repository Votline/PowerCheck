package main

import (
	"log"
	//"time"
	//"flag"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerCheck/internal/ui"
	//"PowerCheck/internal/power"
	"PowerCheck/internal/render"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	//smode := flag.Bool("smode", false, "Silence mode")
	//flag.Parse()

	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
	defer glfw.Terminate()
	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}

	win := ui.CreateWin()
	pg, ofl := render.Setup()
	//pm := power.NewPM()
	pc := ui.CreatePC(pg, ofl)

	for !win.ShouldClose() {
		renderFrame(win, pc)
	}

/*	rCh := make(chan struct{}, 1)
	defer close(rCh)
	
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			if *smode && power.Check() && pm.NfCnt < 4 {
				pm.NfCnt++
				rCh <- struct{}{}
				time.Sleep(1 * time.Minute)
			} else if pm.NfCnt >= 4 {
				pm.NfCnt = 0
				time.Sleep(3 *  time.Minute)
			}
		}
	}()

	if !*smode {
		win.Show()
	}

	for {
		select {
		case <-rCh:
			win.Show()
			renderFrame(win, pc)
		default:
			if !win.ShouldClose() {
				renderFrame(win, pc)
			}
		}
	}*/
}

func renderFrame(win *glfw.Window, pc *ui.PowerChecker) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	
	pc.Render(win)
	
	win.SwapBuffers()
	glfw.PollEvents()
}
