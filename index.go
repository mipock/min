package piscine

func Index(s string, toFind string) int {
	Lnt := len(s)
	LT := len(toFind)

	if LT > Lnt {
		return -1
	}

	for i := 0; i < Lnt-LT; i++ {
		if s[i:i+LT] == toFind {
			return i
		}
	}
	return -1
}
