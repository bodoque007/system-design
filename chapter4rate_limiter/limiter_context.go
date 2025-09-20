package main

type LimiterContext struct {
    limiter RateLimiter
}

func (c *LimiterContext) SetLimiter(l RateLimiter) {
    c.limiter = l
}

func (c *LimiterContext) Allow() bool {
    return c.limiter.Allow()
}
