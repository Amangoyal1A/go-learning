package main

import "fmt"

//////////////////////////////////////////////////////
// Step 1: Define a custom interface
//////////////////////////////////////////////////////

type Payable interface {
	Pay() 
}

//////////////////////////////////////////////////////
// Step 2: Define structs which implement the interface
//////////////////////////////////////////////////////

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay() {
	fmt.Println("Processing payment using Credit Card")
}

type UPI struct {
	UPIID string
}

func (u UPI) Pay() {
	fmt.Println("Processing payment using UPI")
}

//////////////////////////////////////////////////////
// Step 3: Function that accepts the interface
//////////////////////////////////////////////////////

func ProcessPayment(p Payable) {
	// p can be any type that implements Payable
	p.Pay()
}

//////////////////////////////////////////////////////
// Step 4: Main function to call all
//////////////////////////////////////////////////////

func main() {
	cc := CreditCard{CardNumber: "1234-5678-9876"}
	upi := UPI{UPIID: "aman@upi"}

	ProcessPayment(cc)  // Pass struct to interface-based function
	ProcessPayment(upi) // Same function reused for different types
}
