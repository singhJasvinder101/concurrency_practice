# 1.  crawl_rate_limiter

Overview:
----------
A high-concurrency web crawler with intelligent rate-limiting.

This Go-based crawler explores web pages concurrently while strictly adhering to a rate limit of one page every 0.5 seconds. It demonstrates how to coordinate goroutines using WaitGroups and apply a global rate-limiter using time.Tick, ensuring responsible crawling behavior under concurrency.

Features:
----------
- Recursive concurrent crawling with depth control
- Global rate limiter to throttle requests
- Efficient use of goroutines and channels
- Mocked fetcher for safe testing and demonstration
  
Run It:
--------
```
cd crawl_rate_limiter
go run .
```

# 2. Twitter Producer-Consumer

Overview:
----------
A concurrency-focused simulation of a Twitter-like data stream using Golang. This project demonstrates the classic Producer-Consumer model with goroutines, channels, and error handling.

How It Works:
--------------
1. A mocked stream of tweets is fetched by the Producer.
2. Tweets are passed through a channel to the Consumer.
3. The Consumer filters and prints tweets that mention "Go" or "Gopher".
4. The program exits cleanly after consuming all data.

Run It:
--------
```
cd twitter_producer_consumer
go run .
```
