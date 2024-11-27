package piscine

func ReverseMenuIndex(menu []string) []string {
	rev := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		rev[i] = menu[len(menu)-1-i]
	}

	return rev
}
