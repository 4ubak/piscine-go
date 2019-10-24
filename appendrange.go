package piscine

func AppendRange(min, max int) []int {
	var array []int
	if max > min && min > 0 && max > 0 {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
	}
	return array
}
