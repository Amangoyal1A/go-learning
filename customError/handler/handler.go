package handler

import (
	"fmt"
	"net/http"
	"strconv"

	cerr "yourapp/errors"

	"errors"

	"github.com/gin-gonic/gin"
)

var currentBalance float64 = 1000.0 // Dummy in-memory balance

func WithdrawHandler(c *gin.Context) {
	amountStr := c.Query("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be a number"})
		return
	}

	err = Withdraw(currentBalance, amount)

	if err != nil {
		var invalidAmountErr *cerr.InvalidAmountError
		var insufficientErr *cerr.InsufficientFundsError

		// ✅ errors.As to detect specific error type
		switch {
		case errors.As(err, &invalidAmountErr):
			fmt.Println(errors.As(err, &invalidAmountErr))
			c.JSON(http.StatusBadRequest, gin.H{"error": invalidAmountErr.Error()})
			return

		case errors.As(err, &insufficientErr):
			c.JSON(http.StatusForbidden, gin.H{
				"error":     insufficientErr.Error(),
				"suggestion": "Try withdrawing ₹" + strconv.FormatFloat(insufficientErr.Balance, 'f', 2, 64) + " or less",
			})
			return

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
	}

	// Success
	c.JSON(http.StatusOK, gin.H{
		"message": "Withdrawal successful",
		"amount":  amount,
	})
}

// Business logic
func Withdraw(balance, amount float64) error {
	if amount <= 0 {
		return &cerr.InvalidAmountError{Amount: amount}
	}
	if amount > balance {
		return &cerr.InsufficientFundsError{
			Balance: balance,
			Amount:  amount,
		}
	}
	currentBalance -= amount
	return nil
}
