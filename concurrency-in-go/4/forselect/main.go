package main

import (
	"fmt"
)

func main() {
	loop()
	//infinite()
}

func loop() {
	stringStream := make(chan string)
	done := make(chan string)

	go func(ch <-chan string) {
		for {
			s, ok := <-ch
			fmt.Printf("(%v)%s\n", ok, s)
			if s == "b" {
				close(done)
			}
		}
	}(stringStream)

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			fmt.Println("End") // sometimes it's triggered
			return
		case stringStream <- s:
		}
	}
}

func infinite() {
	done := make(chan interface{})
	x := 0
	for {
		select {
		case <-done:
			fmt.Printf("\ninfinite loop finished.\n")
			return
		default:
		}
		x++
		fmt.Printf("%d ", x)
		if x > 20 {
			close(done)
		}
	}

}
