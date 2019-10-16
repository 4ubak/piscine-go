package main

import "github.com/01-edu/z01"

func main() {
	for a := '0'; a <= '9'; a++ {
		z01.PrintRune(a)
	}
	z01.PrintRune(10)
}
