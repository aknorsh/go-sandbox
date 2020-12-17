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

		// https://golang.org has <div id="nav">
		n := elementByID(doc, "nav")

		if n != nil {
			fmt.Printf("Found: <%s", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf(" %s=%q", attr.Key, attr.Val)
			}
			fmt.Printf(">\n")
		} else {
			fmt.Printf("Not found.\n")
		}
		resp.Body.Close()
	}
}

func elementByID(doc *html.Node, id string) *html.Node {
	res, found := forNodes(doc, id)
	if found {
		return res
	}
	return nil
}

func forNodes(n *html.Node, id string) (res *html.Node, ok bool) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return n, true
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res, ok = forNodes(c, id)
		if ok {
			return
		}
	}
	return
}
