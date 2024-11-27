// package main

// import (
// 	"fmt"
// )

// func CountAlpha(s string) int {
// 	tmp := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] <= 'z' && s[i] >= 'a' {
// 			tmp++
// 		}
// 		if s[i] <= 'Z' && s[i] >= 'A' {
// 			tmp++
// 		}
// 	}
// 	return tmp
// }

func CountAlphaRange(s string) int {
	r := 0
	for _, c := range s {
		if c <= 'z' && c >= 'a' {
			r++
		}
		if c <= 'Z' && c >= 'A' {
			r++
		}
	}
	return r
}

// func CountAlphaRangeLen(s string) int {
// 	r := 0
// 	for _, c := range s {
// 		if c != 'H' {
// 			r++
// 		}
// 	}
// 	return r
// }

// func main() {
// 	fmt.Println(CountAlphaRange("Hello world"))
// 	fmt.Println(CountAlphaRange("H e l l o"))
// 	fmt.Println(CountAlphaRange("H1e2l3l4o"))
// }