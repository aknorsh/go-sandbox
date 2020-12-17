// echo - output command-line arguments as they are
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func echoRanged(args []string) {
	s, separator := "", ""
	for _, arg := range args {
		s += separator + arg
		separator = *sep
	}
	fmt.Print(s)
}

func echoImproved(args []string) {
	fmt.Print(strings.Join(args, *sep))
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args)%2 == 1 {
		echoImproved(args)
	} else {
		echoRanged(args)
	}
	if !*n {
		fmt.Println()
	}
}
