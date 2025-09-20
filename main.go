package main

import (
	"fmt"
	"time"
)


func main() {
	bucket := NewTokenBucket(5, 1) // 5 tokens max capacity, refill 1 token per second
	ctx := &LimiterContext{limiter: bucket}

	for i := 0; i < 10; i++ {
		// First five requests should be allowed, all others should be denied
		allowed := ctx.Allow()
		if allowed {
			fmt.Printf("[%d] Request allowed\n", i)
		} else {
			fmt.Printf("[%d] Request denied, fly home buddy, I work alone\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
