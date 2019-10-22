package piscine

func IsLower(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, i := range runeArray {
		if i >= 'a' && i <= 'z' {
			count++
		}
	}
	if count == runeCount {
		return true
	}
	return false
}
