package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("slice before:", s)

	i := 2
	fmt.Println("index to remove:", i)
	s = append(s[:i], s[i+1:]...)
	fmt.Println("slice after:", s)
}
