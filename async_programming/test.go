package main

import (
	"fmt"
	"sync"
	"time"
)

func task(i int, result *[]int, mu *sync.Mutex) {
	// time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Println("Task started:", i)
	mu.Lock()
	*result = append(*result, i)
	mu.Unlock()
}

func main() {
	// Case 1: 10 tasks with 10 goroutines (one per task)
	var result1 []int
	var wg1 sync.WaitGroup
	var mu1 sync.Mutex

	start1 := time.Now()
	for i := 0; i < 10000; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			task(i, &result1, &mu1)
		}(i)
	}
	wg1.Wait()
	elapsed1 := time.Since(start1)
	// fmt.Println("Goroutine Result:", result1)
	fmt.Println("Goroutine Time Taken:", elapsed1)

	// Case 2: Corrected worker pool implementation
	var result2 []int
	var wg2 sync.WaitGroup
	var mu2 sync.Mutex

	tasks := make(chan int, 10000)
	workers := 50

	// Add all tasks to the channel first
	for i := 0; i < 10000; i++ {
		tasks <- i
	}
	close(tasks)

	start2 := time.Now()

	// Create worker goroutines
	for w := 0; w < workers; w++ {
		wg2.Add(1)
		go func(workerID int) {
			defer wg2.Done()
			for i := range tasks {
				task(i, &result2, &mu2)
			}
			
		}(w)
	}

	wg2.Wait()
	elapsed2 := time.Since(start2)
	// fmt.Println("Worker Pool Result:", result2)
	fmt.Println("Worker Pool Time Taken:", elapsed2)
}


