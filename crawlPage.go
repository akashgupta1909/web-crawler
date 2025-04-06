package main

import (
	"fmt"
	"net/url"

	customHTML "github.com/akashgupta1909/web-crawler/internal/customHTML"
	"github.com/akashgupta1909/web-crawler/internal/customURL"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		defer cfg.wg.Done()
	}()

	isMaxPages := cfg.isMaxPagesReached()
	if isMaxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error in parsing current url", err)
		return
	}
	if cfg.baseURL.Host != currentURL.Host {
		return
	}

	normalizedCurrentURL, err := customURL.NormalizeURL(currentURL.String())
	if err != nil {
		fmt.Println("error in normalizing the current url:", err)
		return
	}

	if !cfg.isFirst(normalizedCurrentURL) {
		return
	}

	fmt.Println("starting crawaling on:", normalizedCurrentURL)

	htmlBody, err := customHTML.GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error in fetching html:", err)
	}
	urls, err := customHTML.GetURLsFromHTML(htmlBody, cfg.baseURL)

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
