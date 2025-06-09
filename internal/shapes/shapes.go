package shapes

func CreateQuad(x1, y1, x2, y2 float32) []float32 {
	return []float32 {
		x1, y1, 0.0,
		x1, y2, 0.0,

		x2, y2, 0.0,
		x2, y1, 0.0,
		x1, y1, 0.0,
	}
}
