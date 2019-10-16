package main

import "fmt"

func main() {
	for smallAlph := 'a'; smallAlph <= 'z' ; smallAlph++ {
		fmt.Printf("%c", smallAlph)
	}
	fmt.Println()
}