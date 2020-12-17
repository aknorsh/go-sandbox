package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	shaLen := flag.Int("len", 256, "256 | 384 | 512")
	flag.Parse()

	var buf bytes.Buffer
	buf.ReadFrom(os.Stdin)
	b := buf.Bytes()
	b = b[:(len(b) - 1)] // remove EOF
	switch *shaLen {
	case 256:
		c := sha256.Sum256(b)
		fmt.Printf("%x\n", c)
	case 384:
		c := sha512.Sum384(b)
		fmt.Printf("%x\n", c)
	case 512:
		c := sha512.Sum512(b)
		fmt.Printf("%x\n", c)
	default:
		fmt.Fprintln(os.Stderr, "Invalid flag: len = 256|384|512")
		os.Exit(1)
	}
}
