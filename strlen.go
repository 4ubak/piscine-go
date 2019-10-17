package piscine

func StrLen(str string) int {
	c := 0
	for _, s := range str {
		c++
	}
	return c
}
