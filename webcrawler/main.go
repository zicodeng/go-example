package main

import (
	"fmt"
	"log"
	"os"
)

const usage = `
usage:
	webcrawler <starting-url>
`

type JobResult struct {
	URL   string
	PL    *PageLinks
	Error error
}

func reportResults(result *JobResult, results chan *JobResult) {
	log.Printf("reporting results for %s", result.URL)
	// Hand it to results goroutine to crawl all links
	// found in this result.
	results <- result
}

func startWorking(toFetch chan string, results chan *JobResult) {
	for URL := range toFetch {
		log.Printf("crawling %s", URL)
		links, err := GetPageLinks(URL)
		result := &JobResult{URL, links, err}
		go reportResults(result, results)
	}
}

// numWorkers is the number of worker goroutines
// we will start: begin with just 1 and increase
// to see the benefits of concurrent execution,
// but don't increase beyond the number of concurrent
// socket connections allowed by your OS.
const numWorkers = 100

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	// Use the first argument as our starting URL.
	startingURL := os.Args[1]

	// toFetch is a channel that gets page links
	// for a given URL.
	toFetch := make(chan string)
	// results is a channel that reports page link result.
	results := make(chan *JobResult)
	seen := map[string]bool{}

	// Build a concurrent web crawler
	// with `numWorkers` worker goroutines.
	for i := 0; i < numWorkers; i++ {
		go startWorking(toFetch, results)
	}

	seen[startingURL] = true
	toFetch <- startingURL

	outstandingJobs := 1

	// The JobResult we receive might contain many other page links,
	// we need to crawl those as well.
	for result := range results {
		outstandingJobs--
		if result.Error != nil {
			log.Printf("error crawling %s: %v", result.URL, result.Error)
		} else {
			log.Printf("processing %d links found in %s", len(result.PL.Links), result.URL)
			// Follow all the links found in this page,
			// and add them to toFetch channel, which will be handled
			// by different goroutine.
			for _, URL := range result.PL.Links {
				if !seen[URL] {
					seen[URL] = true
					log.Printf("adding %s to the queue", URL)
					toFetch <- URL
					outstandingJobs++
				}
			}
		}
		if outstandingJobs == 0 {
			break
		}
	}
	log.Println("ALL DONE!")
}
