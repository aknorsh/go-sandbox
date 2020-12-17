package main

import "fmt"

func main() {

	v := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(v)
	rotate(v[:], 3)
	fmt.Println(v)

}
func rotate(v []int, diffIdx int) []int {
	c := make([]int, len(v))
	copy(c, v)
	gapIdx := len(v) - diffIdx
	for idx, el := range c {
		if idx < diffIdx {
			v[gapIdx+idx] = el
		} else {
			v[idx-diffIdx] = el
		}
	}
	return v
}
