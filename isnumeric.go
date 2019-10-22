package piscine

func IsNumeric(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, char := range runeArray {
		if isDigit(char) {
			count++
		}
	}
	if count == runeCount {
		return true
	}
	return false
}
