package piscine

func Swap(a *int, b *int) {
	var tempA = *a
	var tempB = *b
	*a = tempB
	*b = tempA
}

