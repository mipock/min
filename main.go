package main

import "github.com/01-edu/z01"

func main() {
	ch := 'a'
	for ch <= 'z' {
		z01.PrintRune(ch)
		ch++
	}
	z01.PrintRune('\n')
}
