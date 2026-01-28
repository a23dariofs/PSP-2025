package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func saludar(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d: Hello!\n", id)
	tiempo := time.Duration(rand.Intn(401)+100) * time.Millisecond
	time.Sleep(tiempo)

	fmt.Printf("Goroutine %d: Goodbye!\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go saludar(i, &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines completed!")
}
