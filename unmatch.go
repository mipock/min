package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		flag := 0
		for k := 0; k < len(a); k++ {
			if a[i] == a[k] {
				flag++
			}
		}
		if flag%2 != 0 {
			return a[i]
		}
	}
	return -1
}

// a := []int{1, 2, 3, 1, 2, 3, 4}
