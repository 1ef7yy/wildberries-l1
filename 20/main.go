package main

import (
	"fmt"
	"strings"
)

func splitString(s string) string {
	// разбиваем строку на слова
	words := strings.Split(s, " ")
	// опять свапаем in-place
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Printf("Before split: %s\n", s)
	fmt.Printf("After split: %s\n", splitString(s))
}
