package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerChecker/internal/cmd"
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

func CreateWin(winW, winH, alX, alY int) *glfw.Window  {
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
	win.Hide()
	win.MakeContextCurrent()
	win.SetAttrib(glfw.Floating, 1)

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	win.SetPos(vidMode.Width-alX, vidMode.Height-alY)

	glfw.SwapInterval(1)
	return win
}

func BtnCallback(pc *PowerChecker, winW, winH int) func(w *glfw.Window, mBtn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
	return func(w *glfw.Window, mBtn glfw.MouseButton, act glfw.Action, mod glfw.ModifierKey) {
		if mBtn == glfw.MouseButtonLeft && act == glfw.Press {
			for btn, _ := range pc.btns {
				if btn.hover(w, winW, winH) {
					switch btn.Text {
					case "shutdown":
						cmd.Shutdown()
					case "suspend":
						cmd.Suspend()
					}
				}
			}
		}
	}
}
