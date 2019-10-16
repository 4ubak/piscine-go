package main

import "github.com/01-edu/z01"

func main() {
	for smallAlph := 'a'; smallAlph <= 'z'; smallAlph++ {
		z01.PrintRune(rune(smallAlph))
	}
	z01.PrintRune(10)
}
