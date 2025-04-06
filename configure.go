package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	mu                 *sync.RWMutex
	baseURL            *url.URL
	wg                 *sync.WaitGroup
	concurrencyControl chan struct{}
	maxPages           int
}

func configure(maxConcurrency, maxPages int, rawURL string) (*config, error) {
	fmt.Println("configuration settings:")
	fmt.Println("concurrency", maxConcurrency)
	fmt.Println("max pages", maxPages)
	fmt.Println("base url", rawURL)

	baseURL, err := url.Parse(rawURL)
	if err != nil {

		return nil, fmt.Errorf("error in parsing base url: %v", err)
	}
	return &config{
		pages:              make(map[string]int),
		mu:                 &sync.RWMutex{},
		baseURL:            baseURL,
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		maxPages:           maxPages,
	}, nil
}

func (cfg *config) isFirst(normalizedCurrentURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedCurrentURL]
	if ok {
		cfg.pages[normalizedCurrentURL]++
		return false
	}

	cfg.pages[normalizedCurrentURL] = 1
	return true
}

func (cfg *config) isMaxPagesReached() bool {
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()

	if len(cfg.pages) >= cfg.maxPages {
		return true
	}

	return false
}
