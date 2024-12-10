package main

import "fmt"

// https://stackoverflow.com/questions/35707222/swap-two-numbers-golang
func main() {
	a := 1
	b := 2
	fmt.Printf("a: %d, b: %d\n", a, b)

	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
}
