package web_crawler

import (
	"fmt"
	"testing"
)

func TestCrawler(t *testing.T) {
	urlMap := SafeUrlMap{v: make(map[string]string)}

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, urlMap)
	wg.Wait()

	for url := range urlMap.v {
		body, _ := urlMap.Value(url)
		fmt.Printf("found: %s %q\n", url, body)
	}
}
