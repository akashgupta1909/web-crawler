package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/akashgupta1909/web-crawler/internal/customPrint"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", args[1])

	maxConcurrency := 5
	maxPages := 10
	if args[2] != "" {
		convertedNum, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("error converting the maxcurrency, setting it to 5")
			maxConcurrency = 5
		}
		maxConcurrency = convertedNum
	}
	if args[3] != "" {
		convertedNum, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("error converting the maxPages, setting it to 10")
			maxPages = 10
		}
		maxPages = convertedNum
	}

	cfg, err := configure(maxConcurrency, maxPages, args[1])
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(args[1])
	cfg.wg.Wait()

	customPrint.PrintReport(cfg.pages, cfg.baseURL.String())
}
