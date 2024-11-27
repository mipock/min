package piscine

func JumpOver(str string) string {
	s := ""
	for i := 2; i < len(str); i += 3 {
		s += string(str[i])
	}
	return s + "\n"
}
