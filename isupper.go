package piscine

func IsUpper(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, i := range runeArray {
		if i >= 'A' && i <= 'Z' {
			count++
		}
	}
	if count == runeCount {
		return true
	}
	return false
}
