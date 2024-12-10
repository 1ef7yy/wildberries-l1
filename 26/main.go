package main

import (
	"fmt"
	"strings"
)

func IsUnique(s string) bool {
	// tolower для регистронезависимости
	lower := strings.ToLower(s)
	chars := make(map[rune]bool)
	// проверяем уникальность
	for _, char := range lower {
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}

func main() {
	fmt.Printf("abcd: %v\n", IsUnique("abcd"))
	fmt.Printf("abCdefAaf: %v\n", IsUnique("abCdefAaf"))
	fmt.Printf("aabcd: %v\n", IsUnique("aabcd"))
}
