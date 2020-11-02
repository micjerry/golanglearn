package main

import (
	"fmt"
	"os"
	"net/http"
	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		findlinks(url)
	}
}

func findlinks(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	
	forEachNode(doc, startElement, endElement);
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
}
