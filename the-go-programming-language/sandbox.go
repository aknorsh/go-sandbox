package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var i, j, k int
	i = (1 + 1) * 1
	i, j = j, i
	x, j := 3, 0
	x, k = k, x
	i, k = k, i
	fmt.Println(i, j, k, x)
	p := f()
	fmt.Println(p)

	// overflows
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var in int8 = 127
	fmt.Println(in, in+1, in*in)

	// * (-1)
	in = ^in
	in++
	fmt.Println(in)

	// bitset
	var xBit uint8 = 1<<2 | 1<<5
	var yBit uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", xBit)
	fmt.Printf("%08b\n", yBit)

	fmt.Printf("%08b\n", xBit&yBit)
	fmt.Printf("%08b\n", xBit|yBit)
	fmt.Printf("%08b\n", xBit^yBit)
	fmt.Printf("%08b\n", xBit&^yBit) // 差集合をとってる

	// test: color
	rCol := 255
	gCol := 0
	bCol := 0
	fmt.Printf("#%02x%02x%02x\n", rCol, gCol, bCol)

	// strings
	s := "Hello, I'm string!  "
	ss := []string{"hoge", "fuga", "piyo"}
	fmt.Println(strings.Contains(s, ","))
	fmt.Println(strings.Count(s, "m"))
	fmt.Println(strings.Fields(s), len(strings.Fields(s)))
	fmt.Println(strings.HasPrefix(s, "Hell"))
	fmt.Println(strings.Index(s, "m"))
	fmt.Println(strings.Join(ss, "m"))

	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

}

func f() *int {
	v := 1
	return &v
}
