package piscine

func RectPerimeter(w, h int) int {
	p := (w + h) * 2
	if w < 0 || h < 0 {
		return -1
	}
	return p
}
