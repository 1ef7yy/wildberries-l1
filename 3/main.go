package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, num int, sum *int) {
	defer wg.Done()
	*sum += num * num
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	sum := 0

	var wg sync.WaitGroup
	for _, num := range nums {
		wg.Add(1)
		go square(&wg, num, &sum)
	}

	wg.Wait()

	fmt.Println("Сумма квадратов:", sum)
}
