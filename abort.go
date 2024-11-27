package piscine

func Abort(a, b, c, d, e int) int {
	numr := []int{a, b, c, d, e}
	for i := 0; i <= len(numr); i++ {
		for j := i + 1; j < len(numr); j++ {
			if numr[i] > numr[j] {
				a := numr[i]
				numr[i] = numr[j]
				numr[j] = a
			}
		}
	}
	return numr[2]
}
