package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// 非効率バージョン: 15.633 micro sec for 30 args
	stt := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(stt))

	// 効率バージョン: 2.515 micro sec for 30 args
	stt = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(stt))
}
