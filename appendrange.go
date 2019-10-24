package piscine

func AppendRange(min, max int) []int {
	var array []int
	if max > min {
		for i := min; i < max; i++ {
			array = append(array, i)
		}
	}
	return array
}
