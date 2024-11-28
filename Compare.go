package main

import "fmt"


func Compare(a, b string) int {
	var (
	minlen = len(b)
	
)
	for i := 0; i < minlen; i++ {
		if len(a) < len(b) {
			return 1
		} else if len(b) > len(a) {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
