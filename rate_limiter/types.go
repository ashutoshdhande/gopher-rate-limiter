package ratelimiter

type Options struct {
	// WindowMs defines the time window in milliseconds.
	WindowMs uint

	// Max defines the maximum number of requests allowed by an IP.
	Max uint

	// Message is an optional error response message in case the rate limit is exceeded.
	Message string
}

type RateLimiter struct {
	options Options
}
