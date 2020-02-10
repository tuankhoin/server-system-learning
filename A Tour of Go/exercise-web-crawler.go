package main

import (
	"fmt"
	"time"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

var s SafeCounter = SafeCounter{v : make(map[string]bool)}
func (s SafeCounter) isVisited(url string) bool{
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.v[url]
	if ok == false {
		s.v[url]=true
		return false
	}
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel (1).
	// TODO: Don't fetch the same URL twice (2).
	
	// Modification for (2): add a boolean to check visit state of url
	if depth <= 0 || s.isVisited(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		// Modification for (1): make them become different goroutines
		go Crawl(u, depth-1, fetcher)
	}
	return
}
