package piscine

func PrintIfNot(str string) string {
	if len(str) > 2 {
		return "Invalid Input\n"
	} else {
		return "G\n"
	}
}
