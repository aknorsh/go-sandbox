package main

import "fmt"

func main() {

	x := []int{}
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		x = appendInt(x, v)
		fmt.Println(cap(x), x)
	}
	x = appendInt(x, 10, 11, 12, 13, 14, 15)
	fmt.Println(cap(x), x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if cap(x) >= zlen {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		// constでない値で配列を初期化する場合はこう！
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
