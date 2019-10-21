package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		result := 1
		for a := 1; a <= power; a++ {
			result *= nb
		}
		return result
	}
}
