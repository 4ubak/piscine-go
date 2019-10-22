package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(0)
		return
	}
	if n > 0 {
		var array []int
		minValue := 0
		eachValue := 0
		result := 0
		arrayCount := 0
		for n != 0 {
			eachValue = n % 10
			n /= 10
			if eachValue < minValue {
				minValue = eachValue
				array = append([]int{minValue}, array...)
			} else {
				minValue = eachValue
				array = append(array, minValue)
			}
		}
		for _, count := range array {
			arrayCount = count + 1
		}
		for i := 0; i < arrayCount-1; i++ {
			result *= 10
			result = result + array[i]
		}
		z01.PrintRune(rune(result))
	}
}
