package piscine

import "github.com/01-edu/z01"

func StrLen(str string) int {
	var c = 0
	for _, s := range str {
		c++
	}
	return c
}
