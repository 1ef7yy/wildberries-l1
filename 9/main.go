package main

import (
	"fmt"
	"sync"
)

func producer(nums []int, ch chan int) {
	for _, num := range nums {
		ch <- num
	}
	close(ch)
}

func multiplier(chIn chan int, chOut chan int) {
	for num := range chIn {
		chOut <- num * 2
	}
	close(chOut)
}

func printer(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	producing_ch := make(chan int)
	read_ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		producer(nums, producing_ch)
		wg.Done()
	}()

	go func() {
		multiplier(producing_ch, read_ch)
		wg.Done()
	}()

	go func() {
		printer(read_ch)
		wg.Done()
	}()

	wg.Wait()
}
