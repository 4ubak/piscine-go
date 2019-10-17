package piscine

func UltimateDivMod(a *int, b *int) {
	var tempA = *a / *b
	var tempB = *a % *b
	*a = tempA
	*b = tempB
}
