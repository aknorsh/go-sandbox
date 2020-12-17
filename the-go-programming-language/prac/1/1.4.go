// dup - ２回以上出現する行を出現回数とともに返す
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line] = append(counts[line], filename)
		}
	}
	for line, filenames := range counts {
		if len(filenames) > 1 {
			fmt.Printf("%d\t%s\t%s\n", len(filenames), line, strings.Join(remDup(filenames), " "))
		}
	}
}

func remDup(names []string) []string {
	unique := make(map[string]struct{})
	for _, name := range names {
		unique[name] = struct{}{}
	}
	out := []string{}
	for name, _ := range unique {
		out = append(out, name)
	}
	return out
}
