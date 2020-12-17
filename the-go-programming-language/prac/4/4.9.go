package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run it file1 [, file2...]")
		return
	}

	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("cannot open %s: %v\n", path, err)
			continue
		}
		sum := 0
		counts := make(map[string]int)
		s := bufio.NewScanner(f)
		s.Split(bufio.ScanWords)
		for s.Scan() {
			counts[s.Text()]++
			sum++
		}
		fmt.Printf("%s: %d words\n", path, sum)
		for c, n := range counts {
			fmt.Printf("%q %d (%.4v)\n", c, n, float64(n)/float64(sum))
		}
	}
}
