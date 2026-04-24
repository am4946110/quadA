package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			switch {
			case (row == 1 || row == y) && (col == 1 || col == x):
				z01.PrintRune('o')
			case row == 1 || row == y:
				z01.PrintRune('-')
			case col == 1 || col == x:
				z01.PrintRune('|')
			default:
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
