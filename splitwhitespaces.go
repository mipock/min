package piscine

func SplitWhiteSpaces(str string) []string {
	var result []string
	current := ""

	for _, c := range str {
		if c == ' ' || c == '\n' || c == '\t' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(c)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
