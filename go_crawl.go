package main

import (
	"flag"
	"fmt"
)

func main() {
	// Parse command line arguments to get starting URL.
	startUrl := parseArgs()
	
	// Map of all pages.
	var allPages = make(map[string]*WebPage)
	// Create starting page.
	var startPage = newWebPage(startUrl)
	
	// Add starting page to map
	allPages[startUrl] = startPage
	
	// Recurse over links in starting page, and links of linked pages.
	recurseLinks(startPage, allPages)
	
	// Print output.
	for _, page := range allPages {
		fmt.Print(page.toString())
	}
}

// Recurse over links in starting page, and links of linked pages.
func recurseLinks(page *WebPage, allPages map[string]*WebPage) {
	
	page.processPage()
	
	for it := page.links.Front(); it != nil; it = it.Next() {
		url := it.Value.(string)
		
		// Check if we already got this page.
		if allPages[url] == nil {
			// We haven't gotten this page, so get it now.
			newPage := newWebPage(url)
			allPages[url] = newPage
			recurseLinks(newPage, allPages)
		}
	}
}

// Parse command line arguments to get starting URL.
func parseArgs() (string) {
	var startUrl = flag.String("u", "http://localhost:8000", "Starting URL.")
	flag.Parse()
	
	return *startUrl
}