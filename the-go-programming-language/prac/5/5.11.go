package main

import (
	"fmt"
)

var prereqsWithoutCycle = map[string][]string{
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

var prereqsWithCycle = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"linear algebra":       {"calculus"},
	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func main() {
	for _, course := range topoSortWithCycle(prereqsWithCycle) {
		if course == "CYCLE DETECTED!" {
			fmt.Printf("Cycle detection has succeeded.\n")
		}
	}
	for i, course := range topoSortWithCycle(prereqsWithoutCycle) {
		if course != "CYCLE DETECTED!" {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}

func topoSortWithCycle(m map[string][]string) []string {
	var order []string
	inpath := make(map[string]bool)
	seen := make(map[string]bool)
	var visitAll func(cur string) bool

	visitAll = func(curItem string) bool {
		if inpath[curItem] {
			return true
		}
		if seen[curItem] {
			return false
		}
		inpath[curItem] = true
		seen[curItem] = true
		for _, item := range m[curItem] {
			if visitAll(item) {
				// when cycle is detected, return immediately
				return true
			}
		}
		order = append(order, curItem)
		inpath[curItem] = false
		return false
	}

	for key := range m {
		cycleDetected := visitAll(key)
		if cycleDetected {
			return []string{"CYCLE DETECTED!"}
		}
	}
	return order
}
