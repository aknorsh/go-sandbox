package main

import (
	"fmt"
	"time"
)

func main() {
	//example()
	exampleWithBuffer()

}

func example() {
	done := make(chan interface{})
	defer close(done)

	take := genTake()
	repeat := genRepeat()
	sleep := genSleep()

	start := time.Now()
	zeros := take(done, repeat(done, 0), 3)
	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long

	for v := range pipeline {
		fmt.Printf("%v\n", v)
		fmt.Printf("Elasped: %v\n", time.Since(start))
	}
}

func exampleWithBuffer() {
	done := make(chan interface{})
	defer close(done)

	take := genTake()
	repeat := genRepeat()
	sleep := genSleep()
	buffer := genBuffer()

	zeros := take(done, repeat(done, 0), 3)
	short := sleep(done, 1*time.Second, zeros)
	buf := buffer(done, short, 2)
	long := sleep(done, 4*time.Second, buf)
	pipeline := long

	start := time.Now()
	for v := range pipeline {
		fmt.Printf("%v\n", v)
		fmt.Printf("Elasped: %v\n", time.Since(start))
	}
}

func genBuffer() func(<-chan interface{}, <-chan interface{}, int) <-chan interface{} {
	return func(
		done <-chan interface{},
		prevStream <-chan interface{},
		num int,
	) <-chan interface{} {
		valStream := make(chan interface{}, num)
		go func() {
			defer close(valStream)
			for i := range prevStream {
				select {
				case <-done:
					return
				case valStream <- i:
				}
			}
		}()
		return valStream
	}
}

func genSleep() func(<-chan interface{}, time.Duration, <-chan interface{}) <-chan interface{} {
	return func(
		done <-chan interface{},
		duration time.Duration,
		prevStream <-chan interface{},
	) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for v := range prevStream {
				time.Sleep(duration)
				//fmt.Println("hi")
				select {
				case <-done:
					return
				case valStream <- v:
				}
			}
		}()
		return valStream
	}
}

func genRepeat() func(<-chan interface{}, ...interface{}) <-chan interface{} {
	return func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}
}

func genTake() func(<-chan interface{}, <-chan interface{}, int) <-chan interface{} {
	return func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
}
