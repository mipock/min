package piscine

func RetainFirstHalf(str string) string {
	if len(str) == 1 {
		return str
	}
	if len(str) == 0 {
		return str
	}
	return str[:(len(str))/2]
	
}
