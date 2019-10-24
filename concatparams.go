package piscine

func ConcatParams(args []string) string {
	result := ""
	count := 0
	for range args {
		count++
	}
	for i := range args {
		if i == count-1 {
			result += args[i]
		} else {
			result += args[i] + "\n"
		}
	}
	return result
}
