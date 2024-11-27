package piscine

func NRune(s string, n int) rune {
	for i, c := range s {
		if i+1 == n {
			return c
		}
	}
	return 0
}
