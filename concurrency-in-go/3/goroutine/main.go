package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//tryGoroutine()
	memoryConsumeCheck()
}

func tryGoroutine() {
	var wg sync.WaitGroup

	for _, s := range []string{"hello", "welcome", "hi", "hoge"} {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(s)
	}
	wg.Wait()

}

func memoryConsumeCheck() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() { wg.Done(); <-c } // never end till process finishes

	const numGoroutines = 1e6
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
