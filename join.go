package piscine

func Join(strs []string, sep string) string {
	result := ""
	count := 0
	for i := range strs {
		count = i + 1
	}
	for i, str := range strs {
		if i == count-1 {
			result += str
		} else {
			result += str + sep
		}
	}
	return result
}
