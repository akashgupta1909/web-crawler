package customPrint

import (
	"fmt"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

func sortPages(pages map[string]int) []Page {
	keys := make([]string, 0, len(pages))
	for key := range pages {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		countI := pages[keys[i]]
		countJ := pages[keys[j]]
		if countI == countJ {
			return keys[i] < keys[j]
		}
		return countI > countJ
	})

	sortedPages := []Page{}
	for _, key := range keys {
		sortedPages = append(sortedPages, Page{
			URL:   key,
			Count: pages[key],
		})
	}
	return sortedPages
}

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Println("==========================")
	fmt.Println("Report for:", baseURL)
	fmt.Println("==========================")
	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}
