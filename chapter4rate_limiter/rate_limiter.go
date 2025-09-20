package main

type RateLimiter interface {
	Allow() bool
}