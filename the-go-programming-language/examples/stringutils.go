package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	fmt.Println("hoge/fuga/piyo.go -> ", basename("hoge/fuga/piyo.go"))
	fmt.Println("123456 -> ", comma("123456"))
	fmt.Println("1 2 3 ->", intsToString([]int{1, 2, 3}))

	return
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + comma(s[n-3:])
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
