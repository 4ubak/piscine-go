package piscine

func SplitWhiteSpaces(str string) []string {
	count := 0
	hasSpace := false
	for c := range str {

		if hasSpace && c != 0 && (str[c-1] == '\n' || str[c-1] == '\t' || str[c-1] == ' ') && str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			count++
		}
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			hasSpace = true
		}
	}
	count++

	i := 0
	array := make([]string, count)
	hasSpace2 := true
	for _, c := range str {
		if c == '\n' || c == '\t' || c == ' ' {
			if !hasSpace2 {
				i++
			}
			hasSpace2 = true
			continue
		}
		array[i] = array[i] + string(c)
		hasSpace2 = false
	}
	return array
}
