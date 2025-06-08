package digits

const (
	minX = -0.2
	maxX = -0.1
	scale = 1.5
	
	maxY = 0.2 * scale
	minY = -0.39 * scale
)

func CreateVertexDigits(number int, offset float32) ([]float32, int32) {
	midY := maxY * float32(0.37)
	midY2 := maxY * float32(0.5)
	midY3 := maxY * float32(0.6)
	midY4 := maxY * float32(0.75)
	
	bottomY := minY * float32(0.6)
	bottomY2 := minY * float32(0.7)
	bottomY3 := minY * float32(0.3)

	switch number {
	case 1:
		return []float32{
			minX + offset, float32(0.0) * scale, 0.0,
			maxX + offset, maxY * float32(1.025), 0.0,
			maxX + offset, bottomY2, 0.0,
		}, 3
	case 2:
		return []float32{
			minX + offset, midY4, 0.0,
			(minX + maxX)/2 + offset, maxY, 0.0,
			maxX + offset, midY4, 0.0,

			minX + offset, bottomY, 0.0,
			maxX + float32(0.02) + offset, bottomY, 0.0,
		}, 5
	case 3:
		return []float32{
			minX + offset, maxY, 0.0,
			maxX + offset, maxY, 0.0,

			maxX + offset, midY, 0.0,
			minX + offset, midY, 0.0,
			maxX + offset, midY, 0.0,

			maxX + offset, bottomY, 0.0,
			minX + offset, bottomY, 0.0,
		}, 7
	case 4:
		return []float32{
			minX + offset, maxY, 0.0,
			minX + offset, midY2, 0.0,

			maxX + offset, midY2, 0.0,
			maxX + offset, maxY, 0.0,

			maxX + offset, bottomY2, 0.0,
		}, 5
	case 5:
		return []float32{
			maxX + offset, maxY, 0.0,
			minX + offset, maxY, 0.0,

			minX + offset, midY, 0.0,
			maxX + offset, midY, 0.0,

			maxX + offset, bottomY, 0.0,
			minX + offset, bottomY, 0.0,
		}, 6
	case 6:
		return []float32{
			maxX + offset, maxY, 0.0,
			minX + offset, maxY, 0.0,

			minX + offset, bottomY, 0.0,
			maxX + offset, bottomY, 0.0,

			maxX + offset, midY, 0.0,
			minX + offset, midY, 0.0,
		}, 6
	case 7:
		return []float32{
			minX + offset, maxY, 0.0,
			maxX + offset, maxY, 0.0,

			(minX + maxX)/2 + offset, bottomY2, 0.0,
		}, 3
	case 8:
		return []float32{
			minX + offset, maxY, 0.0,
			maxX + offset, maxY, 0.0,

			maxX + offset, bottomY, 0.0,
			minX + offset, bottomY, 0.0,

			minX + offset, maxY, 0.0,
			minX + offset, midY, 0.0,
			maxX + offset, midY, 0.0,
		}, 7
	case 9:
		return []float32{
			maxX + offset, maxY, 0.0,
			minX + offset, maxY, 0.0,

			minX + offset, midY, 0.0,
			maxX + offset, midY, 0.0,
			maxX + offset, maxY, 0.0,

			maxX + offset, minY, 0.0,
		}, 6
	case 0:
		return []float32{
			minX + offset, maxY, 0.0,
			maxX + offset, maxY, 0.0,

			maxX + offset, bottomY, 0.0,
			minX + offset, bottomY, 0.0,

			minX + offset, maxY, 0.0,
		}, 5
	case 10:
		return []float32{
			minX + offset, midY4, 0.0,
			minX + offset, midY3, 0.0,
		}, 2
	case 11:
		return []float32{
			minX + offset, float32(0.0) * scale, 0.0,
			minX + offset, bottomY3, 0.0,
		}, 2
	default:
		return []float32{
			float32(0.0), float32(0.0), float32(0.0),
		}, 1
	}
}

func GetVert(ch rune, offset float32, allV *[]float32, vQn *[]int32) {
	num := int(ch - '0')
	vt, qn := CreateVertexDigits(num, offset)
	*allV = append(*allV, vt...)
	*vQn = append(*vQn, qn)
}
