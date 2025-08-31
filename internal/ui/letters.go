package ui

const (
	letterMinX  = -0.8
	letterMaxX  = -0.7

	letterMaxY = -0.4
	letterMinY = -0.7
)

func (l *letter) create() element {
	switch l.rn {
	case 'S':
		return &letter{
			vtc: []float32{
				letterMaxX, letterMaxY, 0.0,
				letterMinX, (letterMaxY + letterMinY) / 2, 0.0,
				letterMaxX, letterMinY, 0.0,
				letterMinX, letterMinY, 0.0},
			vtq: int32(4)}
	case 'D':
		return &letter{
			vtc: []float32{
				letterMinX, letterMinY, 0.0,
				letterMinX, letterMaxY, 0.0,

				letterMaxX, letterMaxY/0.8, 0.0,
				letterMaxX, (letterMaxY)*1.6, 0.0,
				letterMinX, letterMinY, 0.0},
			vtq: int32(5)}
	}
	return &letter{}
}

func (l *letter) getVtc() []float32 {
	return l.vtc
}

func (l *letter) getVtq() int32 {
	return l.vtq
}

func (l *letter) setData(r rune) {
	l.rn = r
}
