package main

// package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		fmt.Println("Error Input")
		return
	}

	for i := 1; i <= x; i++ {
		if i == 1 || i == x {
			z01.PrintRune('o')
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune('\n')
	for j := 1; j <= y-2; j++ {
		z01.PrintRune('|')
		for k := 1; k <= x-1; k++ {
			if k == x-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')

	}

	if y > 1 {
		for i := 1; i <= x; i++ {
			if i == 1 || i == x {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
	}

	z01.PrintRune('\n')
	z01.PrintRune('\n')
	z01.PrintRune('\n')
}

func main() {
	QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(3, 4)
	QuadA(3, 4)
}
