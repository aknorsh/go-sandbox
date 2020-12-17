package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	fmt.Println("hoge/fuga/piyo.go -> ", basename("hoge/fuga/piyo.go"))
	fmt.Println("123456 -> ", comma("123456"))
	fmt.Println("12345678910 -> ", comma("12345678910"))
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
	var buf bytes.Buffer
	for i, el := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(el)
	}
	return buf.String()
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
