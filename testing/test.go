package main

import (
	"fmt"
	"sync"
	"time"
)

func task(result *[]int, chans chan int, pauseChan chan bool, resumeChan chan bool, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("🚀 Goroutine started to process tasks")

	for {
		select {
		case <-pauseChan:
			fmt.Println("⏸️  [PAUSE SIGNAL RECEIVED] Execution paused. Waiting for resume signal...")
			<-resumeChan
			fmt.Println("▶️  [RESUME SIGNAL RECEIVED] Execution resumed. Back to work!")
		default:
			select {
			case i, ok := <-chans:
				if !ok {
					fmt.Println("📪 [CHANNEL CLOSED] No more data to process. Exiting goroutine...")
					return
				}
				mu.Lock()
				*result = append(*result, i)
				fmt.Printf("✅ Processed task: %d | Total processed: %d\n", i, len(*result))
				mu.Unlock()
				time.Sleep(1 * time.Millisecond) // simulate work
			default:
				time.Sleep(1 * time.Millisecond) // prevent CPU overuse
			}
		}
	}
}

func main() {
	var chans = make(chan int, 1000)
	var result []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	pauseChan := make(chan bool)
	resumeChan := make(chan bool)

	fmt.Println("📦 Filling channel with 1000 jobs...")
	for i := 0; i < 1000; i++ {
		chans <- i
	}
	close(chans)
	fmt.Println("✅ All tasks enqueued. Channel closed.")

	wg.Add(1)
	go task(&result, chans, pauseChan, resumeChan, &mu, &wg)

	// Watcher that pauses execution when channel size hits 500
	go func() {
		for {
			if len(chans) <= 500 {
				fmt.Println("⚠️  [WATCHER] Channel size dropped to 500 or below. Sending pause signal...")
				pauseChan <- true
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Simulate external resume after delay
	go func() {
		fmt.Println("⌛ [SIMULATOR] Waiting 20 seconds to resume...")
		time.Sleep(20 * time.Second)
		fmt.Println("🎯 [SIMULATOR] Sending resume signal now!")
		resumeChan <- true
	}()

	wg.Wait()
	fmt.Println("🏁 All tasks processed!")
	fmt.Printf("🧮 Final Result Count: %d\n", len(result))
}
