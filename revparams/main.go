package main

import (
	"github.com/01-edu/z01"

	"os"
)

func main() {
	arg := os.Args
	count := 0
	for s := range arg {
		count = s + 1
	}
	for j := count - 1; j >= 1; j-- {
		z01.PrintRune('\n')
	}
}
