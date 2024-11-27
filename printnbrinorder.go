package piscine

func PrintNbrInOrder(n int) {
	if n == 0 {
		print("0")
		return
	}

	var amine []int
	for n > 0 {
		digit := n % 10
		amine = append(amine, digit)
		n /= 10
	}

	for i := 0; i < len(amine); i++ {
		min := i
		for j := i + 1; j < len(amine); j++ {
			if amine[j] < amine[min] {
				min = j
			}
		}
		amine[i], amine[min] = amine[min], amine[i]
	}

	for _, digit := range amine {
		print(digit)
	}
}
