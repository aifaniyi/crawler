package main

import "log"

// consumes urls from the job channel, fetch pages and then
// and writes links in page to queue to be further processed
func worker(id int, jobs <-chan string, fetcher IFetcher, safeQ *SafeQueue) {
	for url := range jobs {
		content, err := fetcher.getPage(url)
		if err != nil {
			log.Printf("error fetching url %s", url)
		} else {
			pageData := extractUrls(url, content)
			log.Printf("found %d links on %s", len(pageData.links), pageData.pageUrl)
			for _, link := range pageData.links {
				safeQ.enqueue(link)
			}
		}
	}

}
