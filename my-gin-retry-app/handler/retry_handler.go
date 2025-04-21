package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ExternalRetryHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	respBody, err := callWithRetry(ctx, "https://httpstat.us/503", 3, 1*time.Second)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": respBody})
}

func callWithRetry(ctx context.Context, url string, retries int, delay time.Duration) (string, error) {
	var lastErr error

	for i := 0; i < retries; i++ {
		fmt.Printf("🔁 Attempt %d...\n", i+1)

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		resp, err := http.DefaultClient.Do(req)

		fmt.Println("line 41",resp,err);

		if err != nil {
			fmt.Println("❌ Network error:", err)
			lastErr = err
		} else {
			defer resp.Body.Close()

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				body, _ := io.ReadAll(resp.Body)
				return string(body), nil
			}

			if resp.StatusCode >= 500 && resp.StatusCode < 600 {
				fmt.Printf("⚠️ Server error (%d), retrying...\n", resp.StatusCode)
				lastErr = fmt.Errorf("status: %d", resp.StatusCode)
			} else {
				return "", fmt.Errorf("non-retriable error: %d", resp.StatusCode)
			}
		}

		time.Sleep(delay)
	}

	return "", fmt.Errorf("all retries failed: %w", lastErr)
}
