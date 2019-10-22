package piscine

func IsPrintable(str string) bool {
	runeArray := []rune(str)
	for _, i := range runeArray {
		if charIsPrintable(i) {
			return false
		}
	}
	return true
}

func charIsPrintable(r rune) bool {
	if r >= 0 && r <= 32 {
		return true
	}
	return false
}
