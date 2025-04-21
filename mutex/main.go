package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0           // Shared data
var mutex = &sync.Mutex{} // Lock for synchronizing

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		counter++
		fmt.Printf("[%s] -> Counter: %d\n", name, counter)
		time.Sleep(100 * time.Millisecond) // simulate work
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker("Thread-A", &wg)
	go worker("Thread-B", &wg)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
