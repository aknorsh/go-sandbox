package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "5.2: %v\n", err)
		os.Exit(1)
	}
	showText(doc)
}

func showText(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	} else if n.Type == html.TextNode {
		s := strings.Trim(n.Data, " \n\t")
		if s != "" {
			fmt.Printf("%s\n", s)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		showText(c)
	}
}
