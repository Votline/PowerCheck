package ui

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerCheck/internal/cmd"
	"PowerCheck/internal/digits"
	"PowerCheck/internal/shapes"
	"PowerCheck/internal/letters"
)

type Btn struct {
	X1, Y1 float32
	X2, Y2 float32
	Text string
}

const windowWidth = 200
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

	window.SetMouseButtonCallback(func(w *glfw.Window, btn glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey){
		if btn == glfw.MouseButtonLeft && action == glfw.Press {
			_, _, btns := box4btns()
			if crBtn := HoverOnBtn(w, btns); crBtn != nil {
				switch crBtn.Text{
				case "shutdown":
					cmd.Shutdown()
				case "suspend":
					cmd.Suspend()
				}
			}
		}
	})

	return window
}

func GetDigits(power string) ([]float32, []int32) {
	var (
		offset float32  = 0.0
		allV []float32
		vQn []int32
	)
	if len(power) == 2 {offset += 0.15
	} else if len(power) < 2 {offset += 0.25}
	for _, ch := range power {
		digits.GetVert(ch, offset, &allV, &vQn)
		offset += 0.2
	}
	return allV, vQn
}

func GetButtons() ([]float32, []int32) {
	allV, vQn := text4btns()
	sdBtnVt, ssBtnVt, _ := box4btns()

	allV = append(allV, sdBtnVt...)
	allV = append(allV, ssBtnVt...)
	vQn = append(vQn, int32(5))
	vQn = append(vQn, int32(5))
	return allV, vQn
}

func text4btns() ([]float32, []int32) {
	var (
		offset float32 = -0.02
		width float32 = 0.0
		s string = "SD SS"
		
		allV []float32
		vQn []int32
	)

	for _, ch := range s {
		if ch == ' ' {offset += 0.7}
		letters.GetVert(ch, offset, &allV, &vQn, &width)
		offset += width
	}
	return allV, vQn
}
func box4btns() ([]float32, []float32, []Btn) {
	allBtns := make([]Btn, 2)
	sdBtn := Btn{
		X1: -0.85, Y1: -0.3,
		X2: -0.18, Y2: -0.8,
		Text: "shutdown"}
	ssBtn := Btn{
		X1: 0.17, Y1: -0.3,
		X2: 0.85, Y2: -0.8, 
		Text: "suspend"}

	sdBtnVt := shapes.CreateQuad(
		sdBtn.X1, sdBtn.Y1, 
		sdBtn.X2, sdBtn.Y2,)
	ssBtnVt := shapes.CreateQuad(
		ssBtn.X1, ssBtn.Y1, 
		ssBtn.X2, ssBtn.Y2)
	allBtns = append(allBtns, sdBtn)
	allBtns = append(allBtns, ssBtn)
	return sdBtnVt, ssBtnVt, allBtns
}
func HoverOnBtn(w *glfw.Window, Btns []Btn) *Btn {
	mX, mY := w.GetCursorPos()
	glX := float32(mX)/float32(windowWidth)*2-1
	glY := 1 - float32(mY)/float32(windowHeight)*2

	for _, btn := range Btns {
		if glX >= btn.X1 && glX <= btn.X2 &&
		glY <= btn.Y1 && glY >= btn.Y2 {
			return &btn
		}
	}
	return nil
}
