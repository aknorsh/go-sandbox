package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// fanIn :       /- val
	// fanIn : ch <-- - val
	// fanIn :       \- val

	// fanOut : goroutine <-\
	// fanOut : goroutine <--- pipeline
	// fanOut : goroutine <-/

	findPrime()
}

func findPrime() {
	rand := func() interface{} { return rand.Intn(500000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	toInt := genToInt()
	repeatFn := genRepeatFn()
	take := genTake()

	randIntStream := toInt(done, repeatFn(done, rand))

	primeFinder := func(
		done <-chan interface{},
		intStream <-chan int,
	) <-chan interface{} {
		primeStream := make(chan interface{})
		go func() {
			defer close(primeStream)
			for {
				select {
				case <-done:
					return
				case val := <-intStream:
					for i := 2; i < val; i++ {
						if i+1 == val {
							primeStream <- val
							break
						}
						if val%i != 0 {
							continue
						}
						break
					}
				}
			}
		}()
		return primeStream
	}

	fanIn := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()
		return multiplexedStream
	}

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("PRIMES:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	/*
		for prime := range take(done, primeFinder(done, randIntStream), 10) {
			fmt.Printf("\t%d\n", prime)
		}
	*/
	fmt.Printf("Search took: %v\n", time.Since(start))
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

func genRepeatFn() func(<-chan interface{}, func() interface{}) <-chan interface{} {
	return func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}

		}()
		return valueStream
	}
}

func genToString() func(<-chan interface{}, <-chan interface{}) <-chan string {
	return func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}
}

func genToInt() func(<-chan interface{}, <-chan interface{}) <-chan int {
	return func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
}
