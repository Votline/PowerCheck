package ui

const (
	minX = -0.52
	maxX = -0.42

	scale = 1.0

	maxY = 0.4 * scale
	minY = -0.2 * scale
)

func (d *digit) create() element {
	switch d.rn {
	case '1':
		return &digit{
			vtc: []float32{
				minX, maxY/3, 0.0,
				maxX, maxY,   0.0,
				maxX, minY,   0.0},
			vtq: int32(3)}
	case '2':
		return &digit{
			vtc: []float32{
				minX         , maxY/3, 0.0,
				(maxX+minX)/2, maxY, 0.0,
				maxX         , maxY/3, 0.0,

				minX, minY, 0.0,
				maxX, minY, 0.0},
			vtq: int32(5)}
	case '3':
		return &digit{
			vtc: []float32{
				minX, maxY,   0.0,
				maxX, maxY,   0.0,

				maxX, maxY/3, 0.0,
				minX, maxY/3, 0.0,
				maxX, maxY/3, 0.0,
				
				maxX, minY,   0.0,
				minX, minY,   0.0,},
			vtq: int32(7)}
	case '4':
		return &digit{
			vtc: []float32{
				minX, maxY,   0.0,
				minX, maxY/3, 0.0,
				
				maxX, maxY/3, 0.0,

				maxX, maxY,   0.0,
				maxX, minY,   0.0},
			vtq: int32(5)}
	case '5':
		return &digit{
			vtc: []float32{
				maxX, maxY,   0.0,
				minX, maxY,   0.0,

				minX, maxY/3, 0.0,
				maxX, maxY/3, 0.0,
				
				maxX, minY,   0.0,
				minX, minY,   0.0},
			vtq: int32(6)}
	case '6':
		return &digit{
			vtc: []float32{
				maxX, maxY,   0.0,
				minX, maxY,   0.0,
				
				minX, minY,   0.0,
				maxX, minY,   0.0,

				maxX, maxY/3, 0.0,
				minX, maxY/3, 0.0},
			vtq: int32(6)}
	case '7':
		return &digit{
			vtc: []float32{
				minX-0.03, maxY, 0.0,
				maxX     , maxY, 0.0,
				minX     , minY, 0.0},
			vtq: int32(3)}
	case '8':
		return &digit{
			vtc: []float32{
				minX, maxY,   0.0,
				maxX, maxY,   0.0,
				
				maxX, minY,   0.0,
				minX, minY,   0.0,
				
				minX, maxY,   0.0,
				minX, maxY/3, 0.0,
				maxX, maxY/3, 0.0},
			vtq: int32(7)}
	case '9':
		return &digit{
			vtc: []float32{
				maxX, maxY,   0.0,
				minX, maxY,   0.0,
				
				minX, maxY/3, 0.0,
				maxX, maxY/3, 0.0,

				maxX, maxY,   0.0,
				maxX, minY,   0.0},
			vtq: int32(6)}
	case '0':
		return &digit{
			vtc: []float32{
				minX, maxY, 0.0,
				maxX, maxY, 0.0,
				
				maxX, minY, 0.0,
				minX, minY, 0.0,

				minX, maxY, 0.0},
			vtq: int32(5)}
	case '.':
		dotYmax := float32(minY+0.05)
		dotXmax := float32(minX+0.05)
		return &digit{
			vtc: []float32{
				dotXmax, dotYmax, 0.0,
				minX, dotYmax, 0.0,
				minX, minY, 0.0,
				dotXmax, minY, 0.0},
			vtq: int32(4)}
	case ':':
		dotYmax := float32(minY+0.05)
		dotXmax := float32(minX+0.05)
		return &digit{
			vtc: []float32{
				dotXmax, -dotYmax*1.5, 0.0,
				minX, -dotYmax*1.5, 0.0,
				minX, -minY*1.5, 0.0,
				dotXmax, -minY*1.5, 0.0},
			vtq: int32(4)}

	}
	return &digit{}
}

func (d *digit) getVtc() []float32 {
	return d.vtc
}

func (d *digit) getVtq() int32 {
	return d.vtq
}

func (d *digit) setData(r rune) {
	d.rn = r
}
