package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// IssuesURL is url
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
	CreatedAt  time.Time `json:"created_at"`
}

// Issue is
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User is
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues is
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func splitTerms(results []*Issue) map[string][]*Issue {
	issueMap := map[string][]*Issue{
		"month": []*Issue{},
		"year":  []*Issue{},
		"more":  []*Issue{},
	}

	monthAgo := time.Now().AddDate(0, -1, 0)
	yearAgo := time.Now().AddDate(-1, 0, 0)

	for _, res := range results {
		if res.CreatedAt.Unix() > monthAgo.Unix() {
			issueMap["month"] = append(issueMap["month"], res)
		} else if res.CreatedAt.Unix() > yearAgo.Unix() {
			issueMap["year"] = append(issueMap["year"], res)
		} else {
			issueMap["more"] = append(issueMap["more"], res)
		}
	}
	return issueMap
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	termOrder := []string{"month", "year", "more"}
	resmap := splitTerms(result.Items)
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, term := range termOrder {
		items := resmap[term]
		fmt.Printf("%s: %d issues:\n", term, len(items))
		for _, item := range items {
			fmt.Printf("#%-5d %9.9s %.55s | %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}

	}

}
