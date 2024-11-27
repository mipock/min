package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, str := range table {

		for _, ch := range str {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
