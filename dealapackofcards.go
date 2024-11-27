package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	fmt.Printf("Player 1: %d, %d, %d", deck[0], deck[1], deck[2])
	z01.PrintRune('\n')
	fmt.Printf("Player 2: %d, %d, %d", deck[3], deck[4], deck[5])
	z01.PrintRune('\n')
	fmt.Printf("Player 3: %d, %d, %d", deck[6], deck[7], deck[8])
	z01.PrintRune('\n')
	fmt.Printf("Player 4: %d, %d, %d", deck[9], deck[10], deck[11])
	z01.PrintRune('\n')
}
