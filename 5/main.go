package main

import (
	"fmt"
	"time"
)

func sender(ch chan int, timeout time.Duration) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(timeout)
	}
	close(ch)
}

func reader(ch chan int, timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", msg)
		case <-ticker.C:
			fmt.Println("Timeout")
			return
		}
	}
}

func main() {
	// определяет время завершения программы
	reader_timeout := 3 * time.Second
	// период отправки сообщений
	sender_timeout := 500 * time.Millisecond
	fmt.Printf("reader_timeout: %v, sender_timeout: %v\n", reader_timeout, sender_timeout)
	ch := make(chan int)
	go sender(ch, sender_timeout)
	reader(ch, reader_timeout)
}
