package main

import "github.com/01-edu/z01"

func main() {
	for reversealAlph := 'z'; reversealAlph >= 'a'; reversealAlph-- {
		z01.PrintRune(reversealAlph)
	}
	z01.PrintRune(10)
}
