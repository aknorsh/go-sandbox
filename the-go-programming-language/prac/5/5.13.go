package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var baseHosts []string

func main() {
	for _, arg := range os.Args[1:] {
		url, err := url.Parse(arg)
		if err != nil {
			continue
		}
		baseHosts = append(baseHosts, url.Host)
	}
	bfs(crawl, os.Args[1:])
}

func bfs(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}

	}
}

// url => all hrefs in html
func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func saveHTML(r *http.Response, buf io.Reader) {
	url := r.Request.URL
	dname := "prac/crawled/" + url.Host + url.Path
	fname := strings.TrimSuffix(dname, "/") + "/index.html"

	f, err := os.Create(fname)
	defer f.Close()
	if err != nil {
		log.Print("crawl: cannot mkdir %s: %v", dname, err)
	}
	a, err := ioutil.ReadAll(buf)
	if err != nil {
		log.Print("crawl: cannot mkdir %s: %v", dname, err)
		return
	}
	f.Write(a)
}

func hostIsBase(host string) bool {
	for _, base := range baseHosts {
		if base == host {
			return true
		}
	}
	return false
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	bff := new(bytes.Buffer)

	host := resp.Request.Host
	if hostIsBase(host) {
		bf := io.TeeReader(resp.Body, bff)
		saveHTML(resp, bf)
	} else {
		bff.ReadFrom(resp.Body)
	}

	doc, err := html.Parse(bff)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as html: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // invalid url
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
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
