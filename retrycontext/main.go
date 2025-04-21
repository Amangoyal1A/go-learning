package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Step 1: Ye function ek external API ko simulate karta hai.
// 50% chance hai ki fail ho jaye. Delay bhi random hai 1-2 second tak.
func callExternalAPI(ctx context.Context) error {
	// Random delay (1 to 2 seconds)
	select {
	case <-time.After(time.Duration(rand.Intn(3)+1) * time.Second):
		// 50% chance of failure
		if true {
			return errors.New("external API failed ❌")
		}
		fmt.Println("✅ API call success")
		return nil
	case <-ctx.Done(): // Agar context timeout ho gaya
	fmt.Println("timeout ⏰")
		return ctx.Err()
	}
}

// Step 2: Retry logic jahan context ke andar retry kiya ja raha hai
func retryWithContext(ctx context.Context, retries int, delay time.Duration, fn func(context.Context) error) error {
	var err error

	for i := 0; i < retries; i++ {
		fmt.Printf("🔁 Attempt %d...\n", i+1)

		// Step 2.1: External API ko call karo
		err = fn(ctx)

		// Step 2.2: Agar error nahi aayi, toh success!
		if err == nil {
			return nil
		}

		// Step 2.3: Agar context timeout ya cancel ho gaya ho toh break karo
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Step 2.4: Retry se pehle thoda wait karo
		time.Sleep(delay)
	}

	// Sab attempts fail ho gaye
	return fmt.Errorf("all retries failed: %w", err)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Step 3: Overall timeout 5 second ka rakha hai
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Context ko cleanup karna zaroori hai

	// Step 4: Retry with context
	err := retryWithContext(ctx, 3, 1*time.Second, callExternalAPI)
	if err != nil {
		fmt.Printf("❌ Final error: %v\n", err)
	} else {
		fmt.Println("🎉 All done")
	}
}
