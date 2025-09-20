package main

import (
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