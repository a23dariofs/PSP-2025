package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mu *sync.Mutex) {
	for {
		time.Sleep(1 * time.Second)

		mu.Lock()
		if *seconds <= 0 {
			mu.Unlock()
			break
		}
		*seconds -= 1
		mu.Unlock()
	}
}

func main() {
	count := 5
	var mu sync.Mutex

	go countdown(&count, &mu)

	for {
		time.Sleep(500 * time.Millisecond)

		mu.Lock()
		if count <= 0 {
			mu.Unlock()
			break
		}
		currentCount := count
		mu.Unlock()

		fmt.Println(currentCount)
	}
}
