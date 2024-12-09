package main

import (
	"fmt"
)

func setBit(n int64, i int, val bool) int64 {
	if val {
		return n | (1 << uint(i))
	} else {
		return n & ^(1 << uint(i))
	}
}

func main() {
	var n int64 = 10
	i := 2
	val := true

	valOutput := int8(0)
	if val {
		valOutput = 1
	}

	fmt.Printf("Original value: %b\n", n)
	n = setBit(n, i, val)
	fmt.Printf("Value after setting %d-th bit to %d: %b\n", i, valOutput, n)
}
