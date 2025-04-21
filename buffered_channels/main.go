package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	var arr []int        // arr == nil
	arr = append(arr, 10) // ✅ works!
	fmt.Println(arr)      // Output: [10]
	


	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Got:", msg)
	case msg := <-ch2:
		fmt.Println("Got:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
