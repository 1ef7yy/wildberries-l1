package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("Квадрат %d = %d\n", num, num*num)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range nums {
		wg.Add(1)
		// запускаем горутину
		go square(&wg, num)
	}

	// ждем завершения всех горутин
	wg.Wait()
}
