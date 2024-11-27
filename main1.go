package main

import "github.com/01-edu/z01"

func main() {
	ch := '0'
	for ch <= '9' {
		z01.PrintRune(ch)
		ch++
	}
	z01.PrintRune('\n')
}
