package piscine

func Split(s, sep string) []string {
	var result []string

	if sep == "" {
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}

	for {
		index := Index(s, sep)
		if index == -1 {
			result = append(result, s)
			break
		}
		result = append(result, s[:index])
		s = s[index+len(sep):]
	}
	return result
}

func Index(s, sep string) int {
	if sep == "" {
		return -1
	}

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			return i
		}
	}
	return -1
}
