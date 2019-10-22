package piscine

func ToUpper(s string) string {
	runeArray := []rune(s)
	for i := range runeArray {
		if runeArray[i] >= 'a' && runeArray[i] <= 'z' {
			runeArray[i] = runeArray[i] - 32
		}
	}
	return string(runeArray)
}
