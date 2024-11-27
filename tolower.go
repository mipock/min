package piscine

func ToLower(s string) string {
	nim := []rune(s)
	for i := range nim {
		if nim[i] >= 'A' && nim[i] <= 'Z' {
			nim[i] = nim[i] + 32
		}
	}
	return string(nim)
}
