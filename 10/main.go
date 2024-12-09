package main

import "fmt"

func main() {
	ranges := make(map[int][]float32)

	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, temp := range temperatures {
	}

	for tempRange, temps := range ranges {
		fmt.Printf("%d: %v\n", tempRange, temps)
	}

}
