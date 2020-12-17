package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("5.7: %v", err)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Fatalf("5.7: cannot parse: %v", err)
		}
		sE, eE := generateElementFn()
		forEachNode(doc, sE, eE)
		resp.Body.Close()
	}
}

func generateElementFn() (func(*html.Node), func(*html.Node)) {
	var depth int
	stt := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
			depth++
		}
	}
	edd := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
		}
	}
	return stt, edd
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
