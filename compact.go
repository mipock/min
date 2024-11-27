package piscine

func Compact(ptr *[]string) int {
	cont := 0
	var tb []string
	for _, i := range *ptr {
		if i != "" {
			tb = append(tb, i)
			cont++

		}
		*ptr = tb
	}
	return cont
}
