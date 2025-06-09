package main

import (
	"flag"
	"log"
	"runtime"
	"time"

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

	if !*smode {
		window = ui.CreateWindow()
		if window == nil {
			return
		}

		program = render.Setup()
		if program == 0 {
			return
		}
	}
	for {	
		if window != nil && !window.ShouldClose() {
			gl.Clear(gl.COLOR_BUFFER_BIT)

			render.Digits()
			render.Buttons()

			window.SwapBuffers()
			glfw.PollEvents()
		} else if window != nil && window.ShouldClose() {
			window.Destroy()
			window = nil
		} else if *smode {
			if power.Check() && pm.NfCnt < 4 {
				time.Sleep(pm.Delay)
				if window == nil {
					glfw.Terminate()
					pm.NfCnt++
					pm.Delay = (5*time.Second)
					setupWin(&window, &program)
					go func(){
						if pm.NfCnt >= 3 {
							time.Sleep(5*time.Second)
							pm.NfCnt = 0
						}
					}()
					continue
				}
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}
}

func setupWin(w **glfw.Window, pg *uint32) bool {
	if err := gl.Init(); err != nil {
		log.Fatalln("OpenGL init error. \nErr: ", err)
	}
	if err := glfw.Init(); err != nil {
		log.Fatalln("GLFW init error. \nErr: ", err)
	}
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
