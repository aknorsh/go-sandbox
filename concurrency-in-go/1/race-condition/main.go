package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()

	// time.Sleep(time.Second)
	if data == 0 {
		// time.Sleep(time.Second)
		fmt.Printf("the value is %v\n", data)
	}
}
