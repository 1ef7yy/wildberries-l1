package main

import (
	"fmt"
	"sync"
)

// опять избегаем race condition с помощью мьютекса
type Counter struct {
	mu    *sync.Mutex
	value int64
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int64 {
	return c.value
}

func main() {
	counter := &Counter{
		mu:    &sync.Mutex{},
		value: 0,
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Final value:", counter.GetValue())
}
