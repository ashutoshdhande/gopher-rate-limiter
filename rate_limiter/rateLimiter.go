package ratelimiter

func NewRateLimiter(options Options) *RateLimiter {
	return &RateLimiter{options: options}
}
