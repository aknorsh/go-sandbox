package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatalf("5.17: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("5.17: %v", err)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("parsing golang.org as html: %v", err)
	}
	ress := elementByTagName(doc, os.Args[1:]...)
	for _, res := range ress {
		fmt.Printf("<%s", res.Data)
		for _, attr := range res.Attr {
			fmt.Printf(" %s=%s", attr.Key, attr.Val)
		}
		fmt.Printf(">\n")
	}
}

func elementByTagName(doc *html.Node, name ...string) []*html.Node {
	// init
	res := []*html.Node{}
	bfsQueue := []*html.Node{doc}
	elSet := make(map[string]bool)
	for _, elName := range name {
		elSet[elName] = true
	}

	for len(bfsQueue) > 0 {
		// exec
		n := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if n.Type == html.ElementNode && elSet[n.Data] {
			res = append(res, n)
		}

		// add nxt node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			bfsQueue = append(bfsQueue, c)
		}
	}
	return res
}
