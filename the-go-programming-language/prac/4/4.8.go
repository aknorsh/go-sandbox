package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	letterCounts := make(map[rune]int)
	digitCounts := make(map[rune]int)
	spaceCounts := make(map[rune]int)
	punctCounts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// add counts
		if unicode.IsLetter(r) {
			letterCounts[r]++
		}
		if unicode.IsDigit(r) {
			digitCounts[r]++
		}
		if unicode.IsSpace(r) {
			spaceCounts[r]++
		}
		if unicode.IsPunct(r) {
			punctCounts[r]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\ndetails\n")
	fmt.Printf("rune\tcount\n")
	details := map[string]map[rune]int{
		"letters": letterCounts,
		"digits":  digitCounts,
		"spaces":  spaceCounts,
		"puncts":  punctCounts,
	}
	for k, v := range details {
		if len(v) > 0 {
			fmt.Printf("%s:\n", k)
			for c, n := range v {
				fmt.Printf("%q\t%d\n", c, n)
			}
		}
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 chars\n", invalid)
	}
}
