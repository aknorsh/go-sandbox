package main

import (
	"fmt"
	"time"
)

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	done := make(chan interface{})

	myChan := make(chan interface{})
	go func(done, c chan<- interface{}) {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			c <- i
		}
		close(done)
	}(done, myChan)

	for val := range orDone(done, myChan) {
		fmt.Println(val)
	}

}
