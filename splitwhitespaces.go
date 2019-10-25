package piscine

func SplitWhiteSpaces(str string) []string {
	var arrayString []string
	countWords := 1
	lenStr := 0
	result := ""
	for c := range str {
		if isWhiteSpace(str[c]) {
			countWords++
		}
		lenStr++
	}
	arrayString = make([]string, countWords)
	i := 0
	for j, c := range str {
		if j+1 == lenStr {
			arrayString[i] = result + string(str[j])
		}
		if isWhiteSpace(str[j]) {
			if i <= countWords {
				arrayString[i] = result
				i++
				result = ""
			}
		} else {
			result += string(c)
		}
	}
	return arrayString
}

func isWhiteSpace(r byte) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
