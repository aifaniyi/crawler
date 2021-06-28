package main

import (
	"testing"
)

func TestExtractor(t *testing.T) {
	pageData := extractUrls("https://www.google.com", pageHtmlContent)

	expectedInt := 3
	if len(pageData.links) < expectedInt {
		t.Fatalf("expected %d links found %d", expectedInt, len(pageData.links))
	}
}

const (
	pageHtmlContent = `href="https://myaccount.google.com/brandaccounts?authuser=0&amp;continue=https://www.google.com/&amp;service=https://www.google.com/webhp%3Fauthuser%3D%24authuser"
	 href="https://accounts.google.com/AddSession?hl=en&amp;continue=https://www.google.com/&amp;ec=GAlAmgQ"
	 href="/accounts.google.com/AddSession?hl=en&amp;continue=https://www.google.com/&amp;ec=GAlAmgQ">
		`
)
