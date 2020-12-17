package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	imap := make(map[int64]int64)
	var wg sync.WaitGroup
	var mx sync.Mutex
	increment := func() {

		val := atomic.AddInt64(&count, 1)
		//val := func() int64 { count++; return count }()

		mx.Lock()
		defer mx.Unlock()
		imap[val]++
		wg.Done()
	}

	addCnt := 1000

	wg.Add(addCnt)
	for i := 0; i < addCnt; i++ {
		go increment()

	}
	wg.Wait()
	for i := int64(0); i <= 100; i++ {
		v := imap[i]
		if v == 0 {
			continue
		}
		fmt.Printf("%v: %v\n", i, v)
	}
	fmt.Printf("Unique: %d\n", len(imap))
}
