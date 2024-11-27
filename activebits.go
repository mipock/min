package piscine

func ActiveBits(n int) int {
	nb := 0
	for n > 0 {
		if n%2 != 0 {
			nb++
		}
		n = n / 2
	}
	return nb
}
