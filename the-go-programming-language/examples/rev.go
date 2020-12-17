package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println("org: ", a)
	reverse(a[:])
	fmt.Println("rev: ", a)

	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("org: ", b)
	reverse(b[:4])
	reverse(b[4:])
	reverse(b[:])
	fmt.Println("prv: ", b)

}

func reverse(v []int) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}
