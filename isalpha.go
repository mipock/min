package piscine

func IsAlpha(s string) bool {
	for _, Ch := range s {
		if !((Ch >= 'A' && Ch <= 'Z') || (Ch >= 'a' && Ch <= 'z') || (Ch >= '0' && Ch <= '9')) {
			return false
		}
	}
	return true
}
