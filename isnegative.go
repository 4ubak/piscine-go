package main

import "github.com/01-edu/z01"

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}

func IsNegative(nb int) {
	if nb >= 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
}
