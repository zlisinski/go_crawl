package main

import (
	"container/list"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"golang.org/x/net/html"
)

// A web page and all its links and dependencies.
type WebPage struct {
	// The URL of the page
	url *url.URL
	
	// All the links on the page to the same host.
	links *list.List
	
	// All the images on the page.
	images *list.List
	
	// All the CSS stylesheets used by the page.
	styleSheets *list.List
	
	// All the JS script files used by the page.
	scriptFiles *list.List
}

// A sort of Constructor for type WebPage
func newWebPage(urlString string) (*WebPage) {
	url, err := url.Parse(urlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing URL: %v", err)
		return nil
	}
	
	page := new(WebPage)
	
	page.url = url
	page.links = list.New()
	page.images = list.New()
	page.styleSheets = list.New()
	page.scriptFiles = list.New()
	
	return page
}

// Downloads the page contents and parses all tags.
func (this WebPage) processPage() {
	urlString := this.url.String()
	
	fmt.Printf("Getting %s...", urlString)
	
	// Make the GET request.
	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError getting %s: %v\n", urlString, err)
		return
	}
	
	// Close body when we leave scope.
	defer resp.Body.Close()
	
	// Make sure we got a 200.
	if (resp.StatusCode != 200) {
		fmt.Fprintf(os.Stderr, "\nGET request returned %d for %s\n", resp.StatusCode, urlString)
		return
	}
	
	fmt.Printf("Done\n")
	
	// Parse HTML into tags.
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error tokenizing input")
		return
	}
	
	// Parse all HTML tags for links, images, stylesheets, and script files.
	this.parseTags(doc)
}

// Recursively parses all HTML tags.
func (this WebPage) parseTags(node *html.Node) {
	if node.Type == html.ElementNode {
		//fmt.Println(node.Data)
		switch node.Data {
			case "a":
				this.parseATag(node)
			case "img":
				this.parseImgTag(node)
			case "link":
				this.parseLinkTag(node)
			case "script":
				this.parseScriptTag(node)
		}
	}
	
	// Recurse over children nodes.
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		this.parseTags(child)
	}
}

// Parses an 'a' tag for link URL.
func (this WebPage) parseATag(node *html.Node) {
	for i := 0; i < len(node.Attr); i++ {
		if node.Attr[i].Key == "href" {
			href, err := url.Parse(node.Attr[i].Val)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing href url: %v\n", err)
				return
			}

			if (href.Host == "") {
				// This is a relative URL to the same host, convert it into an absolute URL first.
				resolvedHref := this.url.ResolveReference(href)
				this.links.PushBack(resolvedHref.String())
			} else if href.Host == this.url.Host {
				// The host is same, so add it as is.	
				this.links.PushBack(node.Attr[i].Val)
			}
		}
	}
}

// Parses an 'img' tag for image URL.
func (this WebPage) parseImgTag(node *html.Node) {
	for i := 0; i < len(node.Attr); i++ {
		if node.Attr[i].Key == "src" {
			href, err := url.Parse(node.Attr[i].Val)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing href url: %v\n", err)
				return
			}

			if (href.Host == "") {
				// This is a relative URL to the same host, convert it into an absolute URL first.
				resolvedHref := this.url.ResolveReference(href)
				this.images.PushBack(resolvedHref.String())
			} else {
				this.images.PushBack(node.Attr[i].Val)
			}
		}
	}
}

// Parses a 'link' tag for stylesheet URL.
func (this WebPage) parseLinkTag(node *html.Node) {
	relOk := false
	src := ""
	for i := 0; i < len(node.Attr); i++ {
		// The link tag must have href and rel attributes with the correct values in order to be added.
		if strings.ToLower(node.Attr[i].Key) == "href" {
			src = node.Attr[i].Val
		} else if strings.ToLower(node.Attr[i].Key) == "rel" &&
		          strings.ToLower(node.Attr[i].Val) == "stylesheet" {
			relOk = true
		}

		if (relOk && src != "") {
			href, err := url.Parse(src)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing href url: %v\n", err)
				return
			}

			if (href.Host == "") {
				// This is a relative URL to the same host, convert it into an absolute URL first.
				resolvedHref := this.url.ResolveReference(href)
				this.styleSheets.PushBack(resolvedHref.String())
			} else {
				this.styleSheets.PushBack(src)
			}
		}
	}
}

// Parses a 'script' tag for JS URL.
func (this WebPage) parseScriptTag(node *html.Node) {
	for i := 0; i < len(node.Attr); i++ {
		if node.Attr[i].Key == "src" {
			href, err := url.Parse(node.Attr[i].Val)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing href url: %v\n", err)
				return
			}

			if (href.Host == "") {
				// This is a relative URL to the same host, convert it into an absolute URL first.
				resolvedHref := this.url.ResolveReference(href)
				this.scriptFiles.PushBack(resolvedHref.String())
			} else {
				this.scriptFiles.PushBack(node.Attr[i].Val)
			}
		}
	}
}

// Returns a string representation of the page.
func (this WebPage) toString() (string) {
	var str string
	
	str += fmt.Sprintf("URL: %s\n", this.url.String())
	
	str += fmt.Sprintf("    Links(%d):\n", this.links.Len())
	for e := this.links.Front(); e != nil; e = e.Next() {
		str += fmt.Sprintf("        %s\n", e.Value)
	}
	
	str += fmt.Sprintf("    Images(%d):\n", this.images.Len())
	for e := this.images.Front(); e != nil; e = e.Next() {
		str += fmt.Sprintf("        %s\n", e.Value)
	}
	
	str += fmt.Sprintf("    Style Sheets(%d):\n", this.styleSheets.Len())
	for e := this.styleSheets.Front(); e != nil; e = e.Next() {
		str += fmt.Sprintf("        %s\n", e.Value)
	}
	
	str += fmt.Sprintf("    Script Files(%d):\n", this.scriptFiles.Len())
	for e := this.scriptFiles.Front(); e != nil; e = e.Next() {
		str += fmt.Sprintf("        %s\n", e.Value)
	}
	
	return str
}