package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// опять избегаем race condition с помощью мьютекса
type Counter struct {
	value int32
}

func (c *Counter) Increment(amount int32) {
	atomic.AddInt32(&c.value, amount)
}

func (c *Counter) GetValue() int32 {
	return c.value
}

func main() {
	counter := &Counter{
		value: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment(int32(1))
			}
		}()
	}

	wg.Wait()

	fmt.Println("Final value:", counter.GetValue())
}
