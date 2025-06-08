package letters

func CreateVertexLetters(letter rune, offset float32) ([]float32, int32, float32) {
	switch letter {
	case 'S':
		return []float32{
			-0.7 + offset, -0.4, 0.0,
			-0.8 + offset, -0.5, 0.0,
			-0.7 + offset, -0.6, 0.0,
			-0.8 + offset, -0.7, 0.0,
		}, 4, 0.15
	case 'D':
		return []float32{
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.7, 0.0,

			-0.7 + offset, -0.65, 0.0,
			-0.7 + offset, -0.45, 0.0,
			
			-0.8 + offset, -0.4, 0.0,
		}, 5, 0.15
	case 'T':
		return []float32{
			-0.82 + offset, -0.4, 0.0,
			-0.75 + offset, -0.4, 0.0,

			-0.75 + offset, -0.71, 0.0,
			-0.75 + offset, -0.4, 0.0,

			-0.68 + offset, -0.4, 0.0,
		}, 5, 0.15
	case 'E':
		return []float32{
			-0.73 + offset, -0.4, 0.0,
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.55, 0.0,

			-0.73 + offset, -0.55, 0.0,
			-0.8 + offset, -0.55, 0.0,

			-0.8 + offset, -0.7, 0.0,
			-0.73 + offset, -0.7, 0.0,
		}, 7, 0.1
	case 'O':
		return []float32{
			-0.8 + offset, -0.4, 0.0,
			-0.8 + offset, -0.7, 0.0,

			-0.7 + offset, -0.7, 0.0,
			-0.7 + offset, -0.4, 0.0,

			-0.8 + offset, -0.4, 0.0,
		}, 5, 0.15
	case 'N':
		return []float32{
			-0.8 + offset, -0.7, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.7 + offset, -0.7, 0.0,
			-0.7 + offset, -0.4, 0.0,
		}, 4, 0.2
	case 'P':
		return []float32{
			-0.8 + offset, -0.73, 0.0,
			-0.8 + offset, -0.4, 0.0,

			-0.75 + offset, -0.4, 0.0,
			-0.7 + offset, -0.55, 0.0,
			-0.8 + offset, -0.62, 0.0,
		}, 5, 0.15
	default:
		return []float32{
			0.0, 0.0, 0.0,
		}, 1, 0.15
	}
}

func GetVert(ch rune, offset float32, allV *[]float32, vQn *[]int32) {
	vt, qn := CreateVertexDigits(ch, offset)
	*allV = append(*allV, vt...)
	*vQn = append(*vQn, qn)
}
