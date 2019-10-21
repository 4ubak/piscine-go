package piscine

func Sqrt(nb int) int {
	if nb%2 == 0 && nb > 0 {
		result := 1
		sqrt := 0
		for a := 1; result <= nb; a++ {
			result = a * a
			sqrt++
		}
		return sqrt - 1
	} else {
		return 0
	}
}
