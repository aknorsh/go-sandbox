package main

import (
	"fmt"
	"unicode"
)

func main() {

	v := []byte("      はろー  \n\v      \t  わーるど    \n")
	fmt.Printf("%s\n", v)
	v = pressUTFSpace(v)
	fmt.Printf("%s\n", v)

}

func pressUTFSpace(bs []byte) []byte {
	i := 0
	prev := []byte("X")[0]
	for _, el := range bs {
		if unicode.IsSpace(rune(el)) {
			if unicode.IsSpace(rune(prev)) {
				continue
			} else {
				bs[i] = ' '
				i++
			}
		} else {
			bs[i] = el
			i++
		}
		prev = el
	}
	return bs[:i]
}
