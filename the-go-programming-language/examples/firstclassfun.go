package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune { return r + 1 }

func printAdd1(s string) {
	fmt.Println(s, " -> ", strings.Map(add1, s))
}

func main() {
	printAdd1("golang")
	printAdd1("JavaScript")
	printAdd1("kViewer")
	printAdd1("toyokumo")
	printAdd1("aknorsh")
	printAdd1("google")
}
