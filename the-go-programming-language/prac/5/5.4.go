package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.2: %v\n", err)
		os.Exit(1)
	}

	for _, s := range visitImproved(nil, doc) {
		fmt.Println(s)
	}
}

func visitImproved(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		} else if n.Data == "img" || n.Data == "script" || n.Data == "style" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitImproved(links, c)
	}
	return links
}
