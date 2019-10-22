package piscine

func Join(strs []string, sep string) string {
	result := ""
	count := 0
	for i := range strs {
		count++
	}
	for i, str := range strs {
		if i == count {
			result += str
		}
		result += str + sep
	}
	return result
}
