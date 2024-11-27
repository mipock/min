package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1

	for _, taha := range s {
		if taha == '-' && result == 0 {
			sign = -1
		} else if taha >= '0' && taha <= '9' {
			result = result*10 + int(taha-'0')
		}
	}

	return result * sign
}
