package piscine

func MakeRange(min, max int) []int {
	var array []int
	if max > min {
		array = make([]int, max-min)
		for i := range array {
			array[i] = i + min
		}
	}
	return array
}
