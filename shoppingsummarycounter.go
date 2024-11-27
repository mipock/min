package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	itemCount := make(map[string]int)

	word := ""
	for _, ch := range str {
		if ch == ' ' {
			if word == "" {
				itemCount[""]++
			} else {
				itemCount[word]++
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if word != "" {
		itemCount[word]++
	} else {
		itemCount[word]++
	}
	return itemCount
}
