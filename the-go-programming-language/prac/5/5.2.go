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
	tagMap := make(map[string]int)
	tagMapping(tagMap, doc)
	for name, cnt := range tagMap {
		fmt.Printf("%s\t:\t%d\n", name, cnt)
	}
}

func tagMapping(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagMapping(m, c)
	}
}
