package main

import (
	"fmt"
	"strings"
)

func splitString(s string) string {
	words := strings.Split(s, " ")
	var new_word string
	for i := len(words); i > 0; i-- {
		new_word += words[i-1] + " "
	}
	return new_word
}

func main() {
	s := "snow dog sun"
	fmt.Println(splitString(s))
}
