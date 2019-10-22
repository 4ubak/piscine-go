package piscine

import "github.com/01-edu/z01"

func TrimAtoi(s string) int {
	var array []int
	result := 0
	minusIndex := 0
	firstDigitIndex := 0
	index := 0
	arrayCount := 0
	for _, rune := range s {
		if rune == '-' {
			minusIndex = index
		}
		if isDigit(rune) {
			if firstDigitIndex == 0 {
				firstDigitIndex = index
			}
			array = append(array, int(rune-'0'))
		}
		index++
	}

	for count := range array {
		arrayCount = count + 1
	}

	for i := 0; i < arrayCount; i++ {
		result = result*10 + array[i]
	}

	if minusIndex < firstDigitIndex && minusIndex != 0 {
		result = result * -1
	}
	return result
}

func isDigit(digit rune) bool {
	for a := '0'; a <= '9'; a++ {
		if digit == a {
			return true
		}
	}
	return false
}
