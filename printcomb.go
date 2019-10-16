package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for startNumber := 0; startNumber <= 999; startNumber++ {
		var fisrtDigit = startNumber / 100
		var secondDigit = (startNumber / 10) % 10
		var thirdDigit = startNumber % 10
		if fisrtDigit < secondDigit {
			if secondDigit < thirdDigit {
				z01.PrintRune(startNumber)
			}
		}
	}
}
