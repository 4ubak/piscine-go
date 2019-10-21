package piscine

func IsPrime(nb int) bool {
	if nb%1 == 0 && nb%nb == 0 && nb%2 == 0 {
		return false
	} else {
		return true
	}
}
