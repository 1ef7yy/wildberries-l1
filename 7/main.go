package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	// map в golang не поддерживает конкурентный доступ из коробки
	// https://stackoverflow.com/questions/36167200/how-safe-are-golang-maps-for-concurrent-read-write-operations
	// поэтому используем RWMutex
	// https://stackoverflow.com/questions/76939286/locking-map-of-maps-for-concurrent-access
	mu sync.RWMutex
	m  map[int]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[int]int),
	}
}

func (cm *ConcurrentMap) Put(key, value int) {
	// избегаем race condition
	cm.mu.Lock()
	cm.m[key] = value
	cm.mu.Unlock()
}

func (cm *ConcurrentMap) Get(key int) (int, bool) {
	cm.mu.RLock()
	value, ok := cm.m[key]
	cm.mu.RUnlock()
	return value, ok
}

func main() {
	cm := NewConcurrentMap()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cm.Put(i, i+1)
			fmt.Printf("put %d into map\n", i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		value, ok := cm.Get(i)
		if !ok {
			fmt.Printf("key %d not found\n", i)
		} else {
			fmt.Printf("key %d: %d\n", i, value)
		}
	}
}
