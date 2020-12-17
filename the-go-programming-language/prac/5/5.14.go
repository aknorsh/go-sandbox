package main

import (
	"fmt"
	"os"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func main() {
	allPreqs := []string{}
	getPreq := func(sub string) []string {
		allPreqs = append(allPreqs, sub)
		return prereqs[sub]
	}
	bfs(getPreq, []string{strings.Join(os.Args[1:], " ")})
	fmt.Printf("%s\n", strings.Join(allPreqs, " <- "))
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
