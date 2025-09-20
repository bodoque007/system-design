package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	// max number of tokens
	capacity   int
	// current number of tokens
	tokens     int
	// tokens added per second
	refillRate int
	lastRefill time.Time
	mu sync.Mutex
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Refill won't be called every interval of time. Instead, it is called by the method that verifies if a request can be made
// Thus why we need to count how many seconds have passed since the last refill call (and not have refill as a goroutine that executes every refill_interval seconds)
// This way, tokens are only refilled when needed, less theoretical-book compliant but more performant
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.lastRefill = now

	newTokens := int(elapsed * float64(tb.refillRate))
	if newTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens + newTokens)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}


func main() {
	bucket := NewTokenBucket(5, 1) // 5 tokens max capacity, refill 1 token per second

	for i := 0; i < 10; i++ {
		// First five requests should be allowed, all others should be denied
		allowed := bucket.AllowRequest()
		if allowed {
			fmt.Printf("[%d] Request allowed\n", i)
		} else {
			fmt.Printf("[%d] Request denied, fly home buddy, I work alone\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
