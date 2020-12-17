package main

import "fmt"

func main() {
	v := [5]int{1, 2, 3, 4, 5}
	fmt.Println("bef: ", v)
	reverse(&v)
	fmt.Println("aft: ", v)
}

func reverse(v *[5]int) {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
}
