package main

import (
	"fmt"
	"sync"
	"time"
)


type result struct {
	body string
	urls []string
}

type mockFetcher map[string]*result

var fetcher = mockFetcher{
	"https://golang.org/": &result{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &result{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
}


func (f mockFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}


type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, limiter <- chan time.Time) {
	defer wg.Done()
	
	if depth <= 0 {
		return
	}

	// visit 1 webpage/second with every 0.5 sec wait for next Tick..
	<- limiter

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// wg.Add(len(urls))

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg, limiter)
	}
}

func main() {
	var wg sync.WaitGroup

	limiter := time.Tick(500 * time.Millisecond)

	wg.Add(1)
	go Crawl("https://golang.org/", 3, fetcher, &wg, limiter)

	wg.Wait()
}

