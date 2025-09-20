package main

import (
	"fmt"
	"time"
)

func main() {
	bucket := NewTokenBucket(5, 1) // 5 tokens max capacity, refill 1 token per second
	ctx := &LimiterContext{limiter: bucket}
	for i := 0; i < 13; i++ {
		allowed := ctx.Allow()
		if allowed {
			fmt.Printf("[%d] Request allowed\n", i)
		} else {
			fmt.Printf("[%d] Request denied, fly home buddy, I work alone\n", i)
		}
		time.Sleep(200 * time.Millisecond)
	}

	fixed_window := NewFixedWindow(1, 1*time.Second) // window of 1 second, 1 token per second
	fmt.Print("fixed window!")

	ctx.SetLimiter(fixed_window)
	for i := 0; i < 10; i++ {
		allowed := ctx.Allow()
		if allowed {
			fmt.Printf("[%d] Request allowed\n", i)
		} else {
			fmt.Printf("[%d] Request denied, fly home buddy, I work alone\n", i)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
