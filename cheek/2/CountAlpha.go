package piscine

func CountAlpha(s string) int {
	count := 0
	for _, B := range s {
		if B >= 'a' && B <= 'z' || B >= 'A' && B <= 'Z' {
			count++
		}
	}
	 return count
}
 