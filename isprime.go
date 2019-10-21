package piscine

func IsPrime(nb int) bool {
	if nb < 1000 && nb > 0 {
		if nb&2 == 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
