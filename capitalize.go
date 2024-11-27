package piscine

func Capitalize(s string) string {
	R := []rune(s)
	new := true

	for i, ch := range R {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if new && ch >= 'a' && ch <= 'z' {
				R[i] = ch - 'a' + 'A'
			} else if !new && ch >= 'A' && ch <= 'Z' {
				R[i] = ch - 'A' + 'a'
			}
			new = false
		} else {
			new = true
		}
	}
	return string(R)
}
