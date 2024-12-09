package main

import "fmt"

func reverseString(s []byte) string {
	s = []byte(s)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return string(s)
}

func main() {
	var s string = "hello world"
	fmt.Println(reverseString([]byte(s)))
}
