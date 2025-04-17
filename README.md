Project-1  (crawl_rate_limiter)

A high-concurrency web crawler with intelligent rate-limiting.

This Go-based crawler explores web pages concurrently while strictly adhering to a rate limit of one page every 0.5 seconds. It demonstrates how to coordinate goroutines using WaitGroups and apply a global rate-limiter using time.Tick, ensuring responsible crawling behavior under concurrency.

🚀 Features:
- Recursive concurrent crawling with depth control
- Global rate limiter to throttle requests
- Efficient use of goroutines and channels
- Mocked fetcher for safe testing and demonstration

Perfect for practicing Go concurrency, parallel task execution, and safe resource limiting.

Run: `go run main.go`

