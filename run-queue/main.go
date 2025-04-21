package main

import (
	"fmt"
)

func blockCPU(id int) {
	fmt.Printf("Goroutine-%d running\n", id)
	for i := 0; i < 1e9; i++ {
		_ = i * i // burn CPU
	}
	fmt.Printf("Goroutine-%d done\n", id)
}

func main() {
	for i := 0; i < 100; i++ {
		go blockCPU(i)
	}

	select {} // block forever to observe
}
