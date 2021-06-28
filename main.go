package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	workers = 1
)

// crawl provided url
func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf(`
		Usage: crawler <url>
		example crawler https://mywebsite.com
		`)
	}

	Init(args[1])
}

func Init(url string) {
	var pageFetcher IFetcher

	jobs := make(chan string, workers)
	pageFetcher = NewFetcher()
	safeQ := NewSafeQueue()
	safeQ.enqueue(url)
	history := map[string]struct{}{}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// start workers
	for i := 0; i < workers; i++ {
		go worker(i+1, jobs, pageFetcher, safeQ)
	}

	running := true
	for running {
		select {
		case <-exit: // catch exit signal
			running = false
			log.Printf("crawling stopped")

		default:
			if !safeQ.isEmpty() {
				next := safeQ.dequeue()
				url := next.(string)
				if _, visited := history[url]; !visited {
					jobs <- url
					history[url] = struct{}{} // mark as visited to prevent loops
				}
			}
		}
	}
}
