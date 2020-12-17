package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("hoge $fuga", strings.ToUpper))
	fmt.Println(expand("HOGE $FUGA", strings.ToLower))
}

func expand(s string, f func(string) string) string {
	ss := strings.Split(s, " ")
	res := ""
	for _, st := range ss {
		if st[0:1] == "$" {
			res += f(st[1:])
		} else {
			res += st
		}

	}
	return res
}
