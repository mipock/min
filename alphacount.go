package piscine

func AlphaCount(s string) int {
	an := 0

	for i := 0; i < len(s); i++ {
		d := s[i]
		if d >= 'A' && d <= 'z' {
			an++
		}
	}
	return an
}
