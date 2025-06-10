package main

import (
	"log"
	"time"
	"sync"
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

	var window *glfw.Window
	var program uint32
	pm := power.NewPM()

	rCh := make(chan struct{}, 1)
	var rMu sync.Mutex
	defer close(rCh)

	if !*smode {
		initWin(&window, &program)
	}

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

	for {
		select {
		case <-rCh:
			rMu.Lock()
			if window == nil {
				initWin(&window, &program)
			}
			rMu.Unlock()
		default:
			if window != nil && !window.ShouldClose() {
				rMu.Lock()
				renderFrame(window)
				rMu.Unlock()
				glfw.PollEvents()
			} else if window != nil {
				rMu.Lock()
				window.Destroy()
				window = nil
				rMu.Unlock()
			} else {
				time.Sleep(16 * time.Millisecond)
			}
			if window != nil && window.ShouldClose() && !*smode{
				return
			}
		}
	}

}

func renderFrame(window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	render.Digits()
	render.Buttons()
	window.SwapBuffers()
}

func initWin(w **glfw.Window, pg *uint32) bool {
	*w = ui.CreateWindow()
	if *w == nil {
		return false
	}
	*pg = render.Setup()
	if *pg == 0 {
		(*w).Destroy()
		return false
	}

	return true
}
