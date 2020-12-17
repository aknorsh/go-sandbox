package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("x / X: %d\n", countDiff(c1, c2))
	fmt.Printf("x / x: %d\n", countDiff(c1, c1))
}

func countDiff(c1, c2 [32]uint8) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		c := c1[i] ^ c2[i]
		for j := 0; j < 8; j++ {
			cnt += int((c >> j) & 1)
		}
	}
	return cnt
}
