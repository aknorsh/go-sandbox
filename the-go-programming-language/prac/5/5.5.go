package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			continue
		}
		words, images := countWordsAndImages(doc)
		fmt.Printf("wrd: %d, img: %d\n", words, images)
	}

}

func countWordsAndImages(n *html.Node) (words, images int) {
	recurCntWordsAndImages(&words, &images, n)
	return
}

func recurCntWordsAndImages(w, i *int, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			return
		} else if n.Data == "img" {
			*i++
		}
	}
	if n.Type == html.TextNode {
		*w += len(strings.Split(n.Data, " "))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		recurCntWordsAndImages(w, i, c)
	}
	return
}
