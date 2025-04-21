package main

import (
	"fmt"
	"time"
)


func Worker(ch chan string) {
    msg := <-ch
    fmt.Println("Received:", msg)
}

func sender(ch chan string) {
	fmt.Println("Sender: Sending data to channel...")
	ch <- "Hello from Sender!" // send
	fmt.Println("Sender: Data sent.")
}

func receiver(ch chan string) {
	fmt.Println("Receiver: Waiting to receive data...")
	msg := <-ch // receive
	fmt.Println("Receiver: Got data:", msg)
}

func mains() {
	ch := make(chan string)

	// Launch both go routines
	go sender(ch)
	go receiver(ch)

	// Give them time to complete (or use sync.WaitGroup)
	time.Sleep(2 * time.Second)
	fmt.Println("Main function exiting.")
}


// func main() {
//     ch := make(chan string)
//     go Worker(ch)
// 	time.Sleep(3 * time.Second)
//     ch <- "Hello Go" // send message
// }
