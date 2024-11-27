package piscine

func CountChar(str string, c rune) int {
    hasan := 0
	for _, c := range str {
		if c == 0 {
			hasan++
		}
	}
	return hasan
}