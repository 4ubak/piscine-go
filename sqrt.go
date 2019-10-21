package piscine

func Sqrt(nb int) int {
	result := 1
	sqrt := 0
	for a := 1; result <= nb; a++ {
		result = a * a
		sqrt++
	}
	return sqrt-1
}
