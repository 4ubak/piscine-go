package main

import "github.com/01-edu/z01"

func main() {
	for bigAlph := 'A'; bigAlph <= 'Z'; bigAlph++ {
		z01.PrintRune(rune(bigAlph))
	}
	z01.PrintRune(10)
}
