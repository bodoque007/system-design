package main

import (
	"sync"
	"time"
)

type FixedWindow struct {
    windowSize time.Duration
    limit int
    count int
    windowStart time.Time
    mu sync.Mutex
}
// duration is in seconds
func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
    return &FixedWindow{
        limit:      limit,
        windowSize: windowSize,
        windowStart: time.Now(),
    }
}

func (fw *FixedWindow) Allow() bool {
    fw.mu.Lock()
    defer fw.mu.Unlock()

    now := time.Now()
	// Restart window in case the current request enters in the next window
    if now.Sub(fw.windowStart) >= fw.windowSize {
        fw.count = 0
        fw.windowStart = now
    }

    if fw.count < fw.limit {
        fw.count++
        return true
    }
    return false
}
