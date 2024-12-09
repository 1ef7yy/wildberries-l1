package main

import (
	"context"
	"fmt"
	"sync"
)

func workerChan(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("stop workerChan")
			return
		default:
			fmt.Println("do work")
		}
	}
}

func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop workerContext")
			return
		default:
			fmt.Println("do work")
		}
	}
}

var stopMutex bool
var mu sync.Mutex

func workerMutex() {
	for {
		mu.Lock()
		if stopMutex {
			fmt.Println("stop workerMutex")
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}

func workerWG(wg *sync.WaitGroup) {
	fmt.Println("workerWG work")
	defer wg.Done()
}

func main() {
	stop := make(chan bool)
	go workerChan(stop)
	stop <- true

	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx)
	cancel()

	go workerMutex()
	mu.Lock()
	stopMutex = true
	mu.Unlock()

	var wg sync.WaitGroup
	wg.Add(1)
	go workerWG(&wg)
	wg.Wait()

}
