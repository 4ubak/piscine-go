package piscine

func Capitalize(s string) string {
	runeArray := []rune(s)
	for i, char := range runeArray {
		if isNumberOrAlph(char) {
			if i == 0 || isNumberOrAlph(runeArray[i-1]) == false {
				if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
					runeArray[i] = char - 32
				}
			} else {
				if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
					runeArray[i] = char + 32
				}
			}
		}
	}
	return string(runeArray)
}

func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
