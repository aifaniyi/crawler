package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type IFetcher interface {
	getPage(url string) (string, error)
}

type Fetcher struct {
}

func NewFetcher() *Fetcher {
	return &Fetcher{}
}

func (f *Fetcher) getPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("call to url=%s returned status=%d",
			url, resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("encountered error reading response for url=%s: %v",
			url, err)
	}

	return string(body), nil
}
