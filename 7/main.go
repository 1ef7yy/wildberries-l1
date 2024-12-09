package main

import (
	"fmt"
	"strconv"
	"sync"
)

type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[string]string
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]string),
	}
}

func (cm *ConcurrentMap) Put(key, value string) {
	cm.mu.Lock()
	cm.m[key] = value
	cm.mu.Unlock()
}

func (cm *ConcurrentMap) Get(key string) (string, bool) {
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
			cm.Put(strconv.Itoa(i), strconv.Itoa(i+1))
			fmt.Printf("put %d into map\n", i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		value, ok := cm.Get(strconv.Itoa(i))
		if !ok {
			fmt.Printf("key %d not found\n", i)
		} else {
			fmt.Printf("key %d: %s\n", i, value)
		}
	}
}
