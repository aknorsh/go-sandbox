package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	b := []byte("あいうえお")
	fmt.Printf("%s\n", b)
	reverseBytes(b)
	fmt.Printf("%s\n", b)

	b = []byte("abcde")
	fmt.Printf("%s\n", b)
	reverseBytes(b)
	fmt.Printf("%s\n", b)

}

func reverseBytes(bs []byte) {
	x := bs
	for len(x) > 0 {
		_, sz := utf8.DecodeRune(x)

		// Reverse each UTF-8 character
		for i, j := 0, sz-1; i < j; i, j = i+1, j-1 {
			x[i], x[j] = x[j], x[i]
		}

		x = x[sz:]
	}

	// Reverse all of bytes
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
}
