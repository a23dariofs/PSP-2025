package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Stats struct {
	success int
	failed  int
	mu      sync.Mutex
}

func worker(id int, urls <-chan string, wg *sync.WaitGroup, stats *Stats) {
	defer wg.Done()

	for url := range urls {
		fmt.Printf("Worker %d started processing: %s\n", id, url)

		sleepTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond
		time.Sleep(sleepTime)

		if rand.Float64() < 0.2 {
			fmt.Printf("Worker %d failed to scrape: %s (error: timeout)\n", id, url)

			stats.mu.Lock()
			stats.failed++
			stats.mu.Unlock()
		} else {
			fmt.Printf("Worker %d successfully scraped: %s\n", id, url)

			stats.mu.Lock()
			stats.success++
			stats.mu.Unlock()
		}
	}
}

func main() {

	startTime := time.Now()

	urlList := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
		"https://example.com/page4",
		"https://example.com/page5",
		"https://example.com/page6",
		"https://example.com/page7",
		"https://example.com/page8",
		"https://example.com/page9",
		"https://example.com/page10",
	}

	numWorkers := 3

	urlChan := make(chan string, len(urlList))
	var wg sync.WaitGroup

	stats := &Stats{}

	// Lanzar workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, urlChan, &wg, stats)
	}

	// Enviar URLs al canal
	for _, url := range urlList {
		urlChan <- url
	}
	close(urlChan)

	// Esperar a que terminen los workers
	wg.Wait()

	totalTime := time.Since(startTime)

	fmt.Println("\nAll workers completed!")
	fmt.Println("---------------------------------")
	fmt.Println("Statistics:")
	fmt.Printf("Total URLs:       %d\n", len(urlList))
	fmt.Printf("Successful:       %d\n", stats.success)
	fmt.Printf("Failed:           %d\n", stats.failed)
	fmt.Printf("Total Time:       %.2fs\n", totalTime.Seconds())
	fmt.Println("---------------------------------")
}
