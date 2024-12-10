package main

import (
	"fmt"
	"math/big"
)

// библиотека math/big предназначена для работы с большими числами
func main() {
	a, _ := new(big.Int).SetString("4194304", 10) // 2^22
	b, _ := new(big.Int).SetString("2097152", 10) // 2^21

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Div(a, b)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", prod)
	fmt.Println("a / b =", quot)
}
