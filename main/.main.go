package main

import "github.com/01-edu/z01"

var (
	N = "Mangati"
	T = "livree"
	Y = "2001"
)

func main() {
	for _, c := range N {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
	for i := 0; i < len(T); i++ {
		z01.PrintRune(rune(T[i]))
		// if i == 2 {
		// 	break
		// }
	}
	z01.PrintRune('\n')
	for _, n := range Y {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
