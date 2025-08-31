package ui

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"PowerChecker/internal/power"
	"PowerChecker/internal/render"
)

type elemMesh struct {
	vao uint32
	vtq int32
}

type PowerChecker struct {
	pg uint32
	ofL int32

	digs []elemMesh
	btns map[*btn]elemMesh
	lets map[rune]elemMesh
}

func createElem[T any, PT interface{*T; element}](chars []rune) []elemMesh {
	result := make([]elemMesh, len(chars))
	for i, ch := range chars {
		el := PT(new(T))
		el.setData(ch)
		created := el.create()
		vao := render.CreateVAO(created.getVtc())
		result[i] = elemMesh{vao: vao, vtq: created.getVtq()}
	}
	return result
}

func createBtn(pos [][4]float32, txts []string) map[*btn]elemMesh {
	result := make(map[*btn]elemMesh)
	for i, p := range pos {
		b := &btn{}
		b.setPos(p)
		b.setText(txts[i])
		el := b.create()
		vao := render.CreateVAO(el.getVtc())
		result[b] = elemMesh{vao: vao, vtq: el.getVtq()}
	}
	return result
}

func CreatePC(pg uint32, ofl int32) *PowerChecker {
	lets := make(map[rune]elemMesh)
	letsRunes := []rune{'S', 'D'}
	tempLets := createElem[letter](letsRunes)

	for i, ch := range letsRunes {
		lets[ch] = tempLets[i]
	}
	pos1 := [4]float32{-0.85, -0.3, -0.18, -0.8}
	pos2 := [4]float32{0.17, -0.3, 0.85, -0.8}
	txts := []string{"shutdown", "suspend"}
	pos := [][4]float32{pos1, pos2}
	btns := createBtn(pos, txts)

	digs := createElem[digit]([]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ':'})

	return &PowerChecker{pg: pg, ofL: ofl, digs: digs, btns: btns, lets: lets}
}

func (pc *PowerChecker) Render(win *glfw.Window, winW, winH int) {
	pc.renderPower()
	pc.renderBtns()

	for btn, _ := range pc.btns {
		btn.hover(win, winW, winH)
	}
}

func (pc *PowerChecker) renderPower() {
	offset := float32(0.0)
	nums := []rune(power.Show())
	if len(nums) == 3 {
		offset = float32(0.4)
	} else if len(nums) == 2 {
		offset = float32(0.35)
	} else {
		offset = float32(0.48)
	}
	for _, char := range nums {
		digit := int(char - '0')
		render.ElemRender(pc.pg, pc.ofL,
			pc.digs[digit].vao, pc.digs[digit].vtq, offset)
		offset += 0.23
	}
}

func (pc *PowerChecker) renderBtns() {
	spaceIdx := int32(0)
	offset := float32(0.15)
	
	for _, ch := range "SD SS" {
		if ch == ' ' {
			offset += 0.75
			spaceIdx++
			continue
		}
		if mesh, exists := pc.lets[ch]; exists {
			render.ElemRender(pc.pg, pc.ofL,
				mesh.vao, mesh.vtq, offset)
			offset += 0.15
		}
		for _, v := range pc.btns {
			render.ElemRender(pc.pg, pc.ofL,
				v.vao, v.vtq, 0.0)
		}
	}
}
