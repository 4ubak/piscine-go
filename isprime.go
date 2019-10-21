package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 2 {
			return true
		}
		if (nb%1 == 0 || nb%2 == 0) && nb%nb == 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
