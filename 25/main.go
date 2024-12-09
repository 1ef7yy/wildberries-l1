package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	start := time.Now()
	for time.Since(start) < d {
	}
}

func main() {
	duration := 5 * time.Second
	fmt.Println("Time before sleeping:", time.Now())
	fmt.Println("Sleeping for", duration)
	sleep(duration)
	fmt.Println("Time after sleeping: ", time.Now())
}
