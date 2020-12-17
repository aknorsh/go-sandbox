package main

import "fmt"

func main() {

	s := []string{"hoge", "fuga", "dup", "dup", "piyo", "mo", "mo", "mo"}
	fmt.Println(s)
	s = removeDuplicate(s)
	fmt.Println(s)

}

func removeDuplicate(ss []string) []string {
	idx := 0
	prevS := ""
	for i, s := range ss {
		if i == 0 || i > 0 && s != prevS {
			ss[idx] = s
			idx++
		}
		prevS = s
	}
	return ss[:idx]
}
