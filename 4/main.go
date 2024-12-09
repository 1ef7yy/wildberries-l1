package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	workersNum := 5

	dataChan := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go workerRead(dataChan, &wg)
	}

	go workerWrite(dataChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	select {
	case <-sigChan:
		fmt.Println("\nCtrl+C pressed. Terminating...")
	}

	close(dataChan)
}

func workerRead(dataChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-dataChan:
			fmt.Println(data)
		}
	}
}

func workerWrite(dataChan chan string) {
	for i := 1; ; i++ {
		dataChan <- fmt.Sprintf("Message %d", i)
		time.Sleep(time.Millisecond * 500)
	}
}
