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

	// инициализируем канал
	dataChan := make(chan string)

	// не забываем закрыть канал по завершению
	defer close(dataChan)

	var wg sync.WaitGroup

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go workerRead(dataChan, &wg)
	}

	go workerWrite(dataChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	// ждем ctrl+c
	for {
		sig := <-sigChan
		if sig == os.Interrupt {
			fmt.Println("\nCtrl+C pressed. Terminating...")
			break
		}
	}

}

func workerRead(dataChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// читаем из канала
		msg := <-dataChan
		fmt.Println(msg)
	}
}

func workerWrite(dataChan chan string) {
	for i := 1; ; i++ {
		// пишем в канал инкрементированное значение
		dataChan <- fmt.Sprintf("Message %d", i)
		time.Sleep(time.Millisecond * 500)
	}
}
