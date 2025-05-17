package main

import "fmt"

func main() {
	var val = []int{1, 2, 3, 4, 5}

	// val = append(val, 3)

	f := val

	f[2] = 100
	fmt.Println(val)

}
