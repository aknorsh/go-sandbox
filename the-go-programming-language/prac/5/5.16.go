package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := []string{"hoge", "fuga"}
	var original string
	var ours string

	original, ours = strings.Join(s, "/"), join("/", s...)
	if original == ours {
		fmt.Printf("%s = %s\t:ok\n", original, ours)
	} else {
		fmt.Printf("%s != %s\t:ng\n", original, ours)
		os.Exit(1)
	}

	s = []string{"a"}

	original, ours = strings.Join(s, "/"), join("/", s...)
	if original == ours {
		fmt.Printf("%s = %s\t:ok\n", original, ours)
	} else {
		fmt.Printf("%s != %s\t:ng\n", original, ours)
		os.Exit(1)
	}

	s = []string{}
	original, ours = strings.Join(s, "/"), join("/", s...)
	if original == ours {
		fmt.Printf("%s = %s\t:ok\n", original, ours)
	} else {
		fmt.Printf("%s != %s\t:ng\n", original, ours)
		os.Exit(1)
	}

}

func join(sep string, targets ...string) string {
	s := ""
	for _, target := range targets {
		if s != "" {
			s += sep
		}
		s += target
	}
	return s
}
