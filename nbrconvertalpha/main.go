package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	arg := os.Args
	count := 0
	index := 96
	for i, s := range arg {
		count = i + 1
		if s == "--upper" {
			index = 64
		}
	}
	for j := 1; j <= count-1; j++ {
		for _, element := range arg[j] {
			z01.PrintRune(element + rune(index))
		}
		z01.PrintRune(10)
	}
}
