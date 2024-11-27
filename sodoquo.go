package main

import (
	"os"

	"github.com/01-edu/z01"
)

func putErr() {
	z01.PrintRune('E')
	z01.PrintRune('r')
	z01.PrintRune('r')
	z01.PrintRune('o')
	z01.PrintRune('r')
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 10 {
		putErr()
		return
	}
	grid := make([][]rune, 9)
	for i, arg := range os.Args[1:] {
		if len(arg) != 9 {
			putErr()
			return
		}
		grid[i] = []rune(arg)
	}
	if !isValidSudoku(grid) {
		putErr()
		return
	}
	if solveSudoku(grid) {
		printGrid(grid)
	} else {
		putErr()
	}
}

func solveSudoku(grid [][]rune) bool {
	row, col := findEmptyCell(grid)
	if row == -1 && col == -1 {
		return true
	}
	for sh := '1'; sh <= '9'; sh++ {
		if isSafe(grid, row, col, sh) {
			grid[row][col] = sh
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = '.'
		}
	}
	return false
}

func findEmptyCell(grid [][]rune) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isSafe(grid [][]rune, row, col int, sh rune) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == sh || grid[i][col] == sh || grid[3*(row/3)+i/3][3*(col/3)+i%3] == sh {
			return false
		}
	}
	return true
}

func isValidSudoku(grid [][]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			sh := grid[i][j]
			if sh != '.' && (sh < '1' || sh > '9') {
				return false
			}
		}
	}
	return true
}

func printGrid(grid [][]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			z01.PrintRune(grid[i][j])
			if j < 8 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
