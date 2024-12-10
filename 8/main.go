package main

import (
	"fmt"
)

// если val = 1, то есть мы устанавливаем i-ыму биту значение 1,
// то мы сдвигаем 1 побитово влево i раз (например, если i = 2, то 1 << 2 = 100 (4 в десятичной) ),
// и применяем битовое ИЛИ, чтобы установить 1 в i-ом бите, которое не затронет другие биты, кроме i-го
// Например
// n = 10 (1010)
// i = 2
// val = true
// 1 << 2 = 0100
// 1010 | 0100 = 1110

// если же val = 0, то мы сдвигаем 1 побитово влево i раз
// переворачиваем все биты с помощью ^ (0010 -> 1101)
// и применяем битовое И, чтобы установить 0 в i-ом бите
// Например
// n = 10 (1010)
// i = 0 (первый бит)
// val = false
// 1 << 1 = 0010
// ^0010 = 1101
// 1010 & 1101 = 1000

func setBit(n int64, i int, val bool) int64 {
	if val {
		return n | (1 << uint(i))
	} else {
		return n & ^(1 << uint(i))
	}
}

func main() {
	var n int64 = 10
	i := 0
	val := true

	valOutput := int8(0)
	if val {
		valOutput = 1
	}

	fmt.Printf("Original value: %b\n", n)
	n = setBit(n, i, val)
	fmt.Printf("Value after setting %d-th bit to %d: %b\n", i+1, valOutput, n)
}
