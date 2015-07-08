package main

import (
	"testing"
	"golang.org/x/net/html"
)

const startUrl string = "http://localhost:8000/"

func TestNewWepPageInvalidUrl(t *testing.T) {
	page := newWebPage("%gh&%ij")
	
	if page != nil {
		t.Error("Expected: nil Got: not nil")
	}
}

func TestParseATagAbsoluteSameHost(t *testing.T) {
	node := new(html.Node)
	node.Data = "a"
	attr := html.Attribute{"", "href", startUrl + "1.html"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseATag(node)
	
	expected1 := 1
	val1 := page.links.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.html"
	val2 := page.links.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseATagAbsoluteDiffHost(t *testing.T) {
	node := new(html.Node)
	node.Data = "a"
	attr := html.Attribute{"", "href", "http://www.google.com"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseATag(node)
	
	expected := 0
	val := page.links.Len()
	if val != expected {
		t.Error("Expected:", expected, " Got:", val)
	}
}

func TestParseATagRelative(t *testing.T) {
	node := new(html.Node)
	node.Data = "a"
	attr := html.Attribute{"", "href", "1.html"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseATag(node)
	
	expected1 := 1
	val1 := page.links.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.html"
	val2 := page.links.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseATagInvalidUrl(t *testing.T) {
	node := new(html.Node)
	node.Data = "a"
	attr := html.Attribute{"", "href", "%gh&%ij"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseATag(node)
	
	expected1 := 0
	val1 := page.links.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseATagNoHref(t *testing.T) {
	node := new(html.Node)
	node.Data = "a"
	
	page := newWebPage(startUrl)
	page.parseATag(node)
	
	expected1 := 0
	val1 := page.links.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseImgTagAbsolute(t *testing.T) {
	node := new(html.Node)
	node.Data = "img"
	attr := html.Attribute{"", "src", startUrl + "1.png"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseImgTag(node)
	
	expected1 := 1
	val1 := page.images.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.png"
	val2 := page.images.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseImgTagRelative(t *testing.T) {
	node := new(html.Node)
	node.Data = "img"
	attr := html.Attribute{"", "src", "1.png"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseImgTag(node)
	
	expected1 := 1
	val1 := page.images.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.png"
	val2 := page.images.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseImgTagInvalidUrl(t *testing.T) {
	node := new(html.Node)
	node.Data = "img"
	attr := html.Attribute{"", "src", "%gh&%ij"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseImgTag(node)
	
	expected1 := 0
	val1 := page.images.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseImgTagNoSrc(t *testing.T) {
	node := new(html.Node)
	node.Data = "img"
	
	page := newWebPage(startUrl)
	page.parseImgTag(node)
	
	expected1 := 0
	val1 := page.images.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseLinkTagAbsolute(t *testing.T) {
	node := new(html.Node)
	node.Data = "link"
	attr1 := html.Attribute{"", "href", startUrl + "1.css"}
	attr2 := html.Attribute{"", "rel", "stylesheet"}
	node.Attr = []html.Attribute{attr1, attr2}
	
	page := newWebPage(startUrl)
	page.parseLinkTag(node)
	
	expected1 := 1
	val1 := page.styleSheets.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.css"
	val2 := page.styleSheets.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseLinkTagRelative(t *testing.T) {
	node := new(html.Node)
	node.Data = "link"
	attr1 := html.Attribute{"", "href", "1.css"}
	attr2 := html.Attribute{"", "rel", "stylesheet"}
	node.Attr = []html.Attribute{attr1, attr2}
	
	page := newWebPage(startUrl)
	page.parseLinkTag(node)
	
	expected1 := 1
	val1 := page.styleSheets.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.css"
	val2 := page.styleSheets.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseLinkTagInvalidUrl(t *testing.T) {
	node := new(html.Node)
	node.Data = "link"
	attr1 := html.Attribute{"", "href", "%gh&%ij"}
	attr2 := html.Attribute{"", "rel", "stylesheet"}
	node.Attr = []html.Attribute{attr1, attr2}
	
	page := newWebPage(startUrl)
	page.parseLinkTag(node)
	
	expected1 := 0
	val1 := page.styleSheets.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseLinkTagNoHref(t *testing.T) {
	node := new(html.Node)
	node.Data = "link"
	attr1 := html.Attribute{"", "rel", "stylesheet"}
	node.Attr = []html.Attribute{attr1}
	
	page := newWebPage(startUrl)
	page.parseLinkTag(node)
	
	expected1 := 0
	val1 := page.styleSheets.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseLinkTagNoRel(t *testing.T) {
	node := new(html.Node)
	node.Data = "link"
	attr1 := html.Attribute{"", "href", "1.css"}
	node.Attr = []html.Attribute{attr1}
	
	page := newWebPage(startUrl)
	page.parseLinkTag(node)
	
	expected1 := 0
	val1 := page.styleSheets.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseScriptTagAbsolute(t *testing.T) {
	node := new(html.Node)
	node.Data = "script"
	attr := html.Attribute{"", "src", startUrl + "1.js"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseScriptTag(node)
	
	expected1 := 1
	val1 := page.scriptFiles.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.js"
	val2 := page.scriptFiles.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseScriptTagRelative(t *testing.T) {
	node := new(html.Node)
	node.Data = "script"
	attr := html.Attribute{"", "src", "1.js"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseScriptTag(node)
	
	expected1 := 1
	val1 := page.scriptFiles.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
	expected2 := startUrl + "1.js"
	val2 := page.scriptFiles.Front().Value
	if val2 != expected2 {
		t.Error("Expected:", expected2, " Got:", val2)
	}
}

func TestParseScriptTagInvalidUrl(t *testing.T) {
	node := new(html.Node)
	node.Data = "script"
	attr := html.Attribute{"", "src", "%gh&%ij"}
	node.Attr = []html.Attribute{attr}
	
	page := newWebPage(startUrl)
	page.parseScriptTag(node)
	
	expected1 := 0
	val1 := page.scriptFiles.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestParseScriptTagNoSrc(t *testing.T) {
	node := new(html.Node)
	node.Data = "script"
	
	page := newWebPage(startUrl)
	page.parseScriptTag(node)
	
	expected1 := 0
	val1 := page.scriptFiles.Len()
	if val1 != expected1 {
		t.Error("Expected:", expected1, " Got:", val1)
	}
}

func TestProcessPage(t *testing.T) {
	var startPage = newWebPage(startUrl)
	startPage.processPage()
	output := startPage.toString()
	
	expected := `URL: http://localhost:8000/
    Links(3):
        http://localhost:8000/1.html
        http://localhost:8000/2.html
        http://localhost:8000/404.html
    Images(2):
        http://localhost:8000/1.png
        http://localhost:8000/3.png
    Style Sheets(2):
        http://localhost:8000/1.css
        http://localhost:8000/3.css
    Script Files(2):
        http://localhost:8000/1.js
        http://localhost:8000/3.js
`
		
	if output != expected {
		t.Error("Expected:\n", expected, "\nGot:\n", output)
	}
}

func TestRecurseLinks(t *testing.T) {
	var allPages = make(map[string]*WebPage)
	var startPage = newWebPage(startUrl)
	
	allPages[startUrl] = startPage
	
	recurseLinks(startPage, allPages)
	
	output := ""
	for _, page := range allPages {
		output += page.toString()
	}
	
	expected := `URL: http://localhost:8000/
    Links(3):
        http://localhost:8000/1.html
        http://localhost:8000/2.html
        http://localhost:8000/404.html
    Images(2):
        http://localhost:8000/1.png
        http://localhost:8000/3.png
    Style Sheets(2):
        http://localhost:8000/1.css
        http://localhost:8000/3.css
    Script Files(2):
        http://localhost:8000/1.js
        http://localhost:8000/3.js
URL: http://localhost:8000/1.html
    Links(1):
        http://localhost:8000/2.html
    Images(0):
    Style Sheets(0):
    Script Files(0):
URL: http://localhost:8000/2.html
    Links(1):
        http://localhost:8000/1.html
    Images(0):
    Style Sheets(0):
    Script Files(0):
URL: http://localhost:8000/404.html
    Links(0):
    Images(0):
    Style Sheets(0):
    Script Files(0):
`
		
	if output != expected {
		t.Error("Expected:\n", expected, "\nGot:\n", output)
	}
}