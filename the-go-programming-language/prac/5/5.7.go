package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
		forEachNode(doc, startElement, endElement)
		resp.Body.Close()
	}
}

var depth int

func getAttrs(n *html.Node) string {
	s := ""
	for _, attr := range n.Attr {
		s += fmt.Sprintf(" %s=%q", attr.Key, attr.Val)
	}
	return s
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, " ", n.Data, getAttrs(n))
			return
		}
		fmt.Printf("%*s<%s%s>\n", depth*2, " ", n.Data, getAttrs(n))
		depth++
	}
	if n.Type == html.TextNode {
		s := strings.Trim(n.Data, " \t\n")
		if s != "" {
			fmt.Printf("%*s%s\n", depth*2, " ", n.Data)
		}
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, " ", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			return
		}
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
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
