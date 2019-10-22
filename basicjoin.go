package piscine

func BasicJoin(strs []string) string {
	result := ""
	for _, str := range strs {
		result += string(str)
	}
	return result
}
