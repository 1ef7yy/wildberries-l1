package main

import (
	"fmt"
)

func main() {
	var i int = 10
	var s string = "hello"
	var b bool = true
	var ch chan int = make(chan int)

	vars := []interface{}{i, s, b, ch}

	for _, v := range vars {
		fmt.Printf("Type of %v is %T\n", v, v)
	}
}
