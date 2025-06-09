package letters

const (
	minX = -0.6
	maxX = -0.5
	topY = -0.4
	bottomY = -0.7
)

func CreateVertexLetters(letter rune, offset float32) ([]float32, int32, float32) {
	switch letter {
	case 'S':
		return []float32{
			maxX + offset, topY, 0.0,
			minX + offset, topY - 0.1, 0.0,
			maxX + offset, topY - 0.2, 0.0,
			minX + offset, bottomY, 0.0,
		}, 4, 0.15
	case 'D':
		return []float32{
			minX + offset, topY, 0.0,
			minX + offset, bottomY, 0.0,

			maxX + offset, bottomY + 0.05, 0.0,
			maxX + offset, topY - 0.05, 0.0,
			
			minX + offset, topY, 0.0,
		}, 5, 0.15
	case 'T':
		return []float32{
			minX + 0.02 + offset, topY, 0.0,
			minX + 0.05 + offset, topY, 0.0,

			minX + 0.05 + offset, bottomY - 0.01, 0.0,
			minX + 0.05 + offset, topY, 0.0,

			minX + 0.12 + offset, topY, 0.0,
		}, 5, 0.15
	case 'E':
		return []float32{
			maxX + 0.03 + offset, topY, 0.0,
			minX + offset, topY, 0.0,
			minX + offset, topY - 0.15, 0.0,

			maxX + 0.03 + offset, topY - 0.15, 0.0,
			minX + offset, topY - 0.15, 0.0,

			minX + offset, bottomY, 0.0,
			maxX + 0.03 + offset, bottomY, 0.0,
		}, 7, 0.1
	case 'O':
		return []float32{
			minX + offset, topY, 0.0,
			minX + offset, bottomY, 0.0,

			maxX + offset, bottomY, 0.0,
			maxX + offset, topY, 0.0,

			minX + offset, topY, 0.0,
		}, 5, 0.15
	case 'N':
		return []float32{
			minX + offset, bottomY, 0.0,
			minX + offset, topY, 0.0,

			maxX + offset, bottomY, 0.0,
			maxX + offset, topY, 0.0,
		}, 4, 0.2
	case 'P':
		return []float32{
			minX + offset, bottomY - 0.03, 0.0,
			minX + offset, topY, 0.0,

			minX + 0.05 + offset, topY, 0.0,
			maxX + offset, topY - 0.15, 0.0,
			minX + offset, topY - 0.22, 0.0,
		}, 5, 0.15
	default:
		return []float32{
			0.0, 0.0, 0.0,
		}, 1, 0.0
	}
}

func GetVert(ch rune, offset float32, allV *[]float32, vQn *[]int32, width *float32) {
	vt, qn, wd := CreateVertexLetters(ch, offset)
	*allV = append(*allV, vt...)
	*vQn = append(*vQn, qn)
	*width = wd
}
