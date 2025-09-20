package main

import (
	"sync"
	"time"
)


type TokenBucket struct {
	// max number of tokens
	capacity int64
	// current number of tokens
	tokens float64
	// tokens added per second
	refillRate float64
	lastRefill time.Time
	mu sync.Mutex
}

// refillRate is in tokens per second.
func NewTokenBucket(capacity int64, rate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     float64(capacity),
		refillRate: rate,
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

    tb.tokens += elapsed * tb.refillRate
    if tb.tokens > float64(tb.capacity) {
        tb.tokens = float64(tb.capacity)
    }
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 {
		tb.tokens--
		return true
	}
	return false
}