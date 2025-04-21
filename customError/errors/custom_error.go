package errors

import "fmt"

// InvalidAmountError
type InvalidAmountError struct {
	Amount float64
}

func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("Invalid amount: ₹%.2f. Must be greater than 0", e.Amount)
}

// InsufficientFundsError
type InsufficientFundsError struct {
	Balance float64
	Amount  float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Insufficient funds: Balance ₹%.2f, Tried to withdraw ₹%.2f", e.Balance, e.Amount)
}
