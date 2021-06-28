package main

import (
	"testing"
)

func TestExtractor(t *testing.T) {
	pageData := extractUrls("https://www.google.com", pageHtmlContent)

	t.Logf("links: %v", pageData.links)
}

const (
	pageHtmlContent = `7L10 10.82 13.82 7 15 8.17l-5 5-5-5z"></path></svg></div><a class="gb_6b gb_Aa gb_Pb" href="https://myaccount.google.com/brandaccounts?authuser=0&amp;continue=https://www.google.com/&amp;service=https://www.google.com/webhp%3Fauthuser%3D%24authuser" aria-hidden="true"><div class="gb_7b"><svg focusable="false" height="20" viewbox="0 0 24 24" width="20" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><path d="M19 3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 2v10.79C16.52 14.37 13.23 14 12 14s-4.52.37-7 1.79V5h14zM5 19v-.77C6.74 16.66 10.32 16 12 16s5.26.66 7 2.23V19H5zm7-6c1.94 0 3.5-1.56 3.5-3.5S13.94 6 12 6 8.5 7.56 8.5 9.5 10.06 13 12 13zm0-5c.83 0 1.5.67 1.5 1.5S12.83 11 12 11s-1.5-.67-1.5-1.5S11.17 8 12 8z" fill="#5F6368"></path><path d="M0 0h24v24H0V0z" fill="none"></path></svg></div><div class="gb_9b gb_ac">All Brand accounts</div><svg class="gb_bc" focusable="false" height="24" viewbox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg"><path d="M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z" fill="#5F6368"></path><path d="M0 0h24v24H0z" fill="none"></path></svg></a></div><div class="gb_Qb" tabindex="-1"><a class="gb_vb gb_Of" href="https://accounts.google.com/AddSession?hl=en&amp;continue=https://www.google.com/&amp;ec=GAlAmgQ" href="/accounts.google.com/AddSession?hl=en&amp;continue=https://www.google.com/&amp;ec=GAlAmgQ">
		`
)
