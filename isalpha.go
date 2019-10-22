package piscine

func IsAlpha(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, char := range runeArray {
		if isNumberOrAlph(char) {
			count++
		}
	}
	if count == runeCount {
		return true
	}
	return false
}
