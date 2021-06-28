package main

import (
	"net/url"
	"regexp"
	"strings"
)

type PageData struct {
	pageUrl string
	links   []string
}

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`href=[\'"]?([^\'" >]+)`)
}

func extractUrls(pageUrl, pageHtmlContent string) *PageData {
	parsed, err := url.Parse(pageUrl)
	if err != nil {
		// ignore since we already have content for url
	}
	matches := re.FindAllString(pageHtmlContent, -1)
	urls := []string{}

	for _, match := range matches {
		temp := match
		temp = strings.TrimPrefix(temp, `href="`)
		if strings.HasPrefix(temp, "/") {
			temp = parsed.Scheme + "://" + parsed.Host + temp
		}
		urls = append(urls, temp)
	}

	return &PageData{
		pageUrl: pageUrl,
		links:   urls,
	}
}
