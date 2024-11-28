package main

import "fmt"

func CountAlpha(s string) int {
	r := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			r++
		}
		if c <= 'Z' && c >= 'A' {
				r++
			}
	}
	return r
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
