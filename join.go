package piscine

func Join(strs []string, sep string) string {
	result := ""
	for _, str := range strs {
		result += str + sep
	}
	return result
}
