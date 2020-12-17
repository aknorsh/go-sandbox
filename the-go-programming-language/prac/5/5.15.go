package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%d\n", max(1, 2, 3, 4, 5))
	fmt.Printf("%d\n", max(1, 1, 1, 1, 1))
	fmt.Printf("%d\n", max())

	fmt.Printf("%d\n", min(1, 2, 3, 4, 5))
	fmt.Printf("%d\n", min(5, 5, 5, 5, 5))
	fmt.Printf("%d\n", min())

	_, ok := maxNeedsArg()
	if ok {
		fmt.Println("Error: no arg is allowed")
	} else {
		fmt.Println("maxNeedsArg: ok, no arg is detected.")
	}

	_, ok = minNeedsArg()
	if ok {
		fmt.Println("Error: no arg is allowed")
	} else {
		fmt.Println("minNeedsArg: ok, no arg is detected.")
	}

}

func max(vals ...int) int {
	if len(vals) == 0 {
		return int(math.MaxInt64)
	}
	res := int(math.MinInt64)
	for _, val := range vals {
		if res < val {
			res = val
		}
	}
	return res
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return int(math.MinInt64)
	}
	res := int(math.MaxInt64)
	for _, val := range vals {
		if res > val {
			res = val
		}
	}
	return res
}

func maxNeedsArg(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	return max(vals...), true
}

func minNeedsArg(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	return min(vals...), true
}
