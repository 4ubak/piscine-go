package piscine

func IsAlpha(str string) bool {
	runeArray := []rune(str)
	isAlpha := false
	for _, char := range runeArray {
		if isNumberOrAlph(char) {
			isAlpha = true
		} else {
			isAlpha = false
		}
	}
	return isAlpha
}
